package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/zrcoder/amisgo-examples/todo-app/db"
	"github.com/zrcoder/amisgo-examples/todo-app/model"

	"github.com/gin-gonic/gin"
	"github.com/zrcoder/amisgo/comp"
)

const (
	Prefix = "/api/"
	todos  = "todos"
	todo   = "todo"
)

var (
	Todos = Prefix + todos
	Todo  = Prefix + todo
)

func GetApiHandler() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	handler := gin.Default()
	api := handler.Group(Prefix)
	{
		api.GET(todos, listTodos)
		api.GET(todo, getTodo)
		api.DELETE(todo, deleteTodo)
		api.PATCH(todo, updateTodo)
	}
	return handler
}

func listTodos(c *gin.Context) {
	page, perPage := parsePage(c)
	slog.Info("list todos", "page", page, "perPage", perPage)

	todos, total, err := db.ListTodos(perPage, (page-1)*perPage)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, comp.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, comp.SuccessResponse("", comp.Data{"items": todos, "total": total}))
}

func getTodo(c *gin.Context) {
	id, errMsg := parseID(c)
	if errMsg != "" {
		slog.Error(errMsg)
		c.JSON(http.StatusBadRequest, comp.ErrorResponse(errMsg))
		return
	}

	todo, err := db.GetTodo(id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, comp.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
	id, errMsg := parseID(c)
	if errMsg != "" {
		slog.Error(errMsg)
		c.JSON(http.StatusBadRequest, comp.ErrorResponse(errMsg))
		return
	}

	err := db.DeleteTodo(id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, comp.ErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}

func updateTodo(c *gin.Context) {
	id, errMsg := parseID(c)
	if errMsg != "" {
		slog.Error(errMsg)
		c.JSON(http.StatusBadRequest, comp.ErrorResponse(errMsg))
		return
	}

	todo := &model.Todo{}
	if err := c.ShouldBindJSON(todo); err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, comp.ErrorResponse(err.Error()))
		return
	}
	todo.ID = id

	slog.Info("update todo", "id", id, "title", todo.Title, "detail", todo.Detail)
	if err := db.UpdateTodo(todo); err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, comp.ErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}

func parseID(c *gin.Context) (int64, string) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		return 0, err.Error()
	}
	return id, ""
}

func parsePage(c *gin.Context) (int, int) {
	page, perPage := 1, 10
	if p, err := strconv.Atoi(c.Query("page")); err == nil {
		page = p
	}
	if p, err := strconv.Atoi(c.Query("perPage")); err == nil {
		perPage = p
	}
	return page, perPage
}
