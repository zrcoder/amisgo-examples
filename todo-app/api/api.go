package api

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/zrcoder/amisgo-examples/todo-app/db"
	"github.com/zrcoder/amisgo-examples/todo-app/model"
	"github.com/zrcoder/amisgo-examples/todo-app/util"

	"github.com/gin-gonic/gin"
	"github.com/zrcoder/amisgo/comp"
)

const (
	Prefix = "/api/"
	todos  = "todos"
	todo   = "todo"

	ReadonlyMsg = "the demo is readonly"
)

var (
	Todos = Prefix + todos
	Todo  = Prefix + todo
)

func GetApiHandler() http.Handler {
	gin.SetMode(gin.ReleaseMode)
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
	req := &model.ListRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, comp.ErrorResponse(err.Error()))
		return
	}
	req.Regular()

	todos, total, err := db.ListTodos(req)
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
	if util.ReadOnly() {
		c.String(http.StatusForbidden, ReadonlyMsg)
		return
	}
	ids, err := parseIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, comp.ErrorResponse(err.Error()))
		return
	}
	err = db.DeleteTodos(ids)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, comp.ErrorResponse(err.Error()))
		return
	}
	c.Status(http.StatusNoContent)
}

func parseIDs(c *gin.Context) ([]int64, error) {
	ids := strings.Split(c.Query("ids"), ",")
	fmt.Println("IDS:", ids)

	if len(ids) == 0 {
		return nil, errors.New("no ids found")
	}

	var err error
	res := make([]int64, len(ids))
	for i, v := range ids {
		res[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			slog.Error("pars id failed", slog.String("id", v), slog.String("error", err.Error()))
			return nil, errors.New("invalid id")
		}
	}
	return res, nil
}

func updateTodo(c *gin.Context) {
	if util.ReadOnly() {
		c.String(http.StatusForbidden, ReadonlyMsg)
		return
	}
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
