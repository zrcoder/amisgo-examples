package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

const (
	loginUrl = "/login"
)

func main() {
	index := comp.Page().InitApi("/echo").Body("${body}")
	login := comp.Page().Body(
		comp.Form().Body(
			comp.InputEmail().Name("user"),
			comp.InputPassword().Name("password"),
		),
	)
	ag := amisgo.New().
		Use(testMiddleware).
		Mount("/", index).
		Mount(loginUrl, login).
		HandleFunc("/echo", echo)

	panic(ag.Run())
}

func testMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// before actions, such as access log, auth, and so on
		fmt.Println("Method:", r.Method, "Path:", r.URL.Path)
		if r.URL.Path != loginUrl && !checkAuth(r) {
			amisgo.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		w.Header().Set("test", "test heander value")

		next.ServeHTTP(w, r)

		// after actions, such as debug logs
		fmt.Println("response heander for [test]:", w.Header().Get("test"))
	})
}

func echo(w http.ResponseWriter, r *http.Request) {
	resp := comp.SuccessResponse("", comp.Data{"body": "Hello, Amisgo!"})
	w.Write(resp.Json())
}

func checkAuth(r *http.Request) bool {
	// parse token from r and process auth
	// here just demo, randomly return the auth result
	return rand.Intn(2) == 0
}
