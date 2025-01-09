package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zrcoder/amisgo-examples/todo-app/auth"
)

const (
	Prefix = "/api/"

	loginPath      = "login"
	registerPath   = "register"
	logoutPath     = "logout"
	unregisterPath = "unregister"
	userPath       = "user"
	todosPath      = "todos"
	todoPath       = "todo"

	Login      = Prefix + loginPath
	Register   = Prefix + registerPath
	Logout     = Prefix + logoutPath
	Unregister = Prefix + unregisterPath
	User       = Prefix + userPath
	Todos      = Prefix + todosPath
	Todo       = Prefix + todoPath

	ReadonlyMsg = "the demo is readonly"
)

func New() http.Handler {
	gin.SetMode(gin.ReleaseMode)

	handler := gin.Default()
	api := handler.Group(Prefix)
	{

		api.POST(registerPath, register)
		api.POST(loginPath, login)
		api.POST(logoutPath, logout)
		api.DELETE(unregisterPath, unregister)

		api.Use(auth.Api)
		api.GET(userPath, getUser)
		api.GET(todosPath, listTodos)
		api.POST(todoPath, addTodo)
		api.GET(todoPath, getTodo)
		api.DELETE(todoPath, deleteTodo)
		api.PATCH(todoPath, updateTodo)
	}
	return handler
}
