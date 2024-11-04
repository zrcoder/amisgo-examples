package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"todo/api"
	"todo/db"
	"todo/page"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/config"
)

func main() {
	initDb()
	api.Init()
	go waitForGracefulExit()

	ag := amisgo.New(
		config.WithLang(config.LangEn),
	).
		Redirect("/", "/todos").
		Register("/todos", page.List())

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
