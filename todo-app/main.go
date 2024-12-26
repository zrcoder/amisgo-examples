package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/zrcoder/amisgo-examples/todo-app/api"
	"github.com/zrcoder/amisgo-examples/todo-app/db"
	"github.com/zrcoder/amisgo-examples/todo-app/page"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

func main() {
	initDb()

	ag := amisgo.New(
		conf.WithIcon("https://raw.githubusercontent.com/zrcoder/amisgo-assets/refs/heads/main/logo.svg"),
	).
		Handle(api.Prefix, api.GetApiHandler()).
		Redirect("/", "/todos", http.StatusPermanentRedirect).
		Mount("/todos", page.List())

	go waitForGracefulExit()

	slog.Info("Listening on http://localhost:8888")
	log.Fatal(ag.Run(":8888"))
}

func initDb() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
}

func waitForGracefulExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	slog.Info("terminating")

	// Ensure DB connection closes on application exit
	db.Close()

	os.Exit(0)
}
