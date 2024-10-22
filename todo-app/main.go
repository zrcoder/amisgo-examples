package main

import (
	_ "amisgo-examples/todo-app/db"
	"amisgo-examples/todo-app/page"

	"log"

	"github.com/zrcoder/amisgo"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	amisgo.Serve("/", page.List())

	cfg := amisgo.GetDefaultConfig()
	cfg.Lang = amisgo.LangEn

	log.Println("Listening on http://localhost")
	if err := amisgo.ListenAndServe(":80", cfg); err != nil {
		log.Fatal(err)
	}
}
