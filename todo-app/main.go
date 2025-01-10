package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zrcoder/amisgo-examples/todo-app/api"
	"github.com/zrcoder/amisgo-examples/todo-app/auth"
	"github.com/zrcoder/amisgo-examples/todo-app/db"
	"github.com/zrcoder/amisgo-examples/todo-app/page"
	"github.com/zrcoder/amisgo-examples/todo-app/util"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

const (
	icon  = "https://raw.githubusercontent.com/zrcoder/amisgo-assets/refs/heads/main/logo.svg"
	title = "Todos"
	theme = conf.ThemeAng
)

func main() {
	app := setup()
	done := make(chan bool, 1)

	go waitForGracefulExit(app, done)
	go run(app)

	<-done
	slog.Info("Graceful shutdown complete.")
}

func setup() *amisgo.Engine {
	if !util.ReadOnly() {
		logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug})
		slog.SetDefault(slog.New(logHandler))
	}

	return amisgo.New(
		conf.WithIcon(icon),
		conf.WithTitle(title),
		conf.WithTheme(theme),
	).
		Handle(api.Prefix, api.New()).
		Redirect("/", "/todos", http.StatusPermanentRedirect).
		Mount("/todos", page.Index(), auth.UI).
		Mount("/login", page.Login()).
		Mount("/register", page.Register())
}

func run(app *amisgo.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	slog.Info("Listening on http://localhost:" + port)
	err := app.Run(":" + port)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("http server error: %s", err)
	}
}

func waitForGracefulExit(app *amisgo.Engine, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	slog.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := app.Shutdown(ctx)
	if err != nil {
		slog.Info("Server forced to shutdown", "error", err)
	}

	err = db.Close()
	if err != nil {
		slog.Info("Database forced to shutdown", "error", err)
	}

	slog.Info("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}
