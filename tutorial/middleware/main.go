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
		// Pre-processing actions, such as logging access, authentication, etc.
		fmt.Println("Method:", r.Method, "Path:", r.URL.Path)
		if r.URL.Path != loginUrl && !checkAuth(r) {
			amisgo.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		w.Header().Set("test", "test heander value")

		next.ServeHTTP(w, r)

		// Post-processing actions, such as logging debug information
		fmt.Println("response heander for [test]:", w.Header().Get("test"))
	})
}

func echo(w http.ResponseWriter, r *http.Request) {
	resp := comp.SuccessResponse("", comp.Data{"body": "Hello, Amisgo!"})
	w.Write(resp.Json())
}

func checkAuth(r *http.Request) bool {
	// Parse the token from the request and process authentication.
	// This is just a demonstration; it randomly returns the authentication result.
	return rand.Intn(2) == 0
}
