package api

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/zrcoder/amisgo-examples/todo-app/auth"
	"github.com/zrcoder/amisgo-examples/todo-app/db"
	"github.com/zrcoder/amisgo-examples/todo-app/model"
	"github.com/zrcoder/amisgo-examples/todo-app/util"

	"github.com/gin-gonic/gin"
	am "github.com/zrcoder/amisgo/model"
	"golang.org/x/crypto/bcrypt"
)

func register(c *gin.Context) {
	if util.ReadOnly() {
		c.String(http.StatusForbidden, ReadonlyMsg)
		return
	}

	input := &model.UserRequest{}
	err := c.ShouldBind(input)
	if err != nil {
		invalidInputResp(c)
		return
	}
	usr, err := db.GetUserByName(input.Name)
	if err != sql.ErrNoRows {
		if err != nil {
			slog.Error("db query user", "error", err)
			internalErrorResp(c)
			return
		}
		slog.Info("user already exist", "user", usr.Name)
		c.JSON(http.StatusBadRequest, am.ErrorResponse("user already exist"))
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("bcrypt generate", "error", err)
		internalErrorResp(c)
		return
	}
	user := &model.User{
		Name:         input.Name,
		PasswordHash: string(hashedPassword),
	}
	err = db.AddUser(user)
	if err != nil {
		slog.Error("db add user", "error", err)
		internalErrorResp(c)
		return
	}
	c.JSON(http.StatusCreated, am.SuccessResponse("succeed", nil))
}

func login(c *gin.Context) {
	input := &model.UserRequest{}
	err := c.ShouldBind(input)
	if err != nil {
		invalidInputResp(c)
		return
	}
	user, err := db.GetUserByName(input.Name)
	if err != nil {
		if err != sql.ErrNoRows {
			slog.Error("db get user", "error", err)
			internalErrorResp(c)
		} else {
			slog.Info("login", slog.String("error", "no user found"))
			c.JSON(http.StatusBadRequest, am.ErrorResponse("invalid user or password"))
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		slog.Info("compare hashed password", "error", err)
		c.JSON(http.StatusBadRequest, am.ErrorResponse("invalid user or password"))
		return
	}
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	err = auth.Add(id, user.ID)
	if err != nil {
		slog.Error("add session cache", slog.String("error", err.Error()))
		internalErrorResp(c)
		return
	}
	c.SetCookie(auth.SessionKey, id, 1800, "/", "", true, true)
	c.JSON(http.StatusOK, am.SuccessResponse("succeed", nil))
}

func logout(c *gin.Context) {
	sessionID, err := c.Cookie(auth.SessionKey)
	if err != nil {
		slog.Error("logout", slog.String("error", "no cookie"))
		c.JSON(http.StatusBadRequest, am.ErrorResponse("invalid session"))
		return
	}
	auth.Delete(sessionID)
	c.JSON(http.StatusOK, am.SuccessResponse("succeed", nil))
}

func unregister(c *gin.Context) {
	if util.ReadOnly() {
		c.String(http.StatusForbidden, ReadonlyMsg)
		return
	}

	defer c.Redirect(http.StatusPermanentRedirect, "/register")
	sessionID, err := c.Cookie(auth.SessionKey)
	if err != nil {
		slog.Error("unregister", slog.String("error", "no cookie"))
		return
	}
	userID := auth.Get(sessionID)
	if userID == -1 {
		return
	}
	auth.Delete(sessionID)
	err = db.DeleteUser(userID)
	if err != nil {
		slog.Error("db delete user", "error", err)
		return
	}
}

func getUser(c *gin.Context) {
	userID := c.GetInt64(auth.UserKey)
	user, err := db.GetUser(userID)
	if err != nil {
		slog.Error("get user", slog.String("error", err.Error()), slog.Int64("user id", userID))
		c.JSON(http.StatusOK, am.ErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, am.SuccessResponse("success", am.Data{"name": user.Name}))
}

func invalidInputResp(c *gin.Context) {
	c.JSON(http.StatusBadRequest, am.ErrorResponse("invalid input"))
}
func internalErrorResp(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, am.ErrorResponse("internal error"))
}
