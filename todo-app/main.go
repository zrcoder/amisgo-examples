package main

import (
	"log"
	"log/slog"

	"amisgo-examples/todo-app/api"
	"amisgo-examples/todo-app/db"
	"amisgo-examples/todo-app/page"

	"github.com/zrcoder/amisgo"
)

func main() {
	initDb()
	api.Init()

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
