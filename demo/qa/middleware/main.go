package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/util"
)

const loginPath = "/login"

func main() {
	app := amisgo.New()
	index := app.Page().Body("Hello, Amisgo!")
	login := app.Page().Body(
		app.Form().Body(
			app.InputEmail().Name("user"),
			app.InputPassword().Name("password"),
		),
	)
	app.Mount("/", index, checkAuthMiddleware, testMiddleware)
	app.Mount(loginPath, login)

	panic(app.Run(":8080"))
}

// 鉴权检查，失败则重定向到登录页。
func checkAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("check auth middleware")
		if r.URL.Path != loginPath && !checkAuth(r) {
			util.Redirect(w, r, loginPath, http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// 测试中间件，设置响应头并记录调试信息。
func testMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test middleware")
		w.Header().Set("test", "test heander value")
		next.ServeHTTP(w, r)
		fmt.Println("response heander for [test]:", w.Header().Get("test"))
	})
}

func checkAuth(r *http.Request) bool {
	return rand.Intn(2) == 0 // 仅示例，这里随机返回鉴权结果
}
