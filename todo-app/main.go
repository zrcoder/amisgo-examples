package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"amisgo-examples/todo-app/api"
	"amisgo-examples/todo-app/db"
	"amisgo-examples/todo-app/page"

	"github.com/zrcoder/amisgo"
)

func main() {
	initDb()
	api.Init()
	go waitForGracefulExit()

	amisgo.Redirect("/", "/todos")
	amisgo.Serve("/todos", page.List())
	cfg := amisgo.GetDefaultConfig()
	cfg.Lang = amisgo.LangEn
	slog.Info("Listening on http://localhost")
	log.Fatal(amisgo.ListenAndServe(":80", cfg))
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
