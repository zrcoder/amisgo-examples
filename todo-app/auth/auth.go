package auth

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/zrcoder/amisgo/model"
	"github.com/zrcoder/amisgo/util"
)

const (
	SessionKey = "todo_app_session_id"
	UserKey    = "user_id"

	sessionLife = 30 * time.Minute
)

var sessions = cache.New(sessionLife, 5*time.Minute)

func UI(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(SessionKey)
		if err != nil {
			slog.Error("check auth", "error", err)
			util.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		slog.Debug("ui auth", slog.String("session id", c.Value))
		if _, ok := sessions.Get(c.Value); !ok {
			slog.Error("ui auth", slog.String("error", "no user found for the session"))
			util.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Api(c *gin.Context) {
	id, err := c.Cookie(SessionKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse("auth failed"))
		c.Abort()
		return
	}
	slog.Debug("api auth", slog.String("session id", id))
	userID, ok := sessions.Get(id)
	if !ok {
		c.JSON(http.StatusBadRequest, model.ErrorResponse("auth failed"))
		c.Abort()
		return
	}
	c.Set(UserKey, userID.(int64))
	c.Next()
}

func Add(seesionID string, userID int64) error {
	return sessions.Add(seesionID, userID, sessionLife)
}
func Delete(sessionID string) {
	sessions.Delete(sessionID)
}

// Get gets the user id for the session, returns -1 if not found.
func Get(sessionID string) int64 {
	res, found := sessions.Get(sessionID)
	if found {
		return res.(int64)
	}
	return -1
}
