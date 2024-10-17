package main

import (
	_ "amisgo-examples/todo-app/db"
	"amisgo-examples/todo-app/page"
	"log"

	"github.com/zrcoder/amisgo"
)

func main() {
	amisgo.Serve("/", page.List)

	log.Println("Listening on http://localhost:80")
	if err := amisgo.ListenAndServe(":80"); err != nil {
		log.Fatal(err)
	}
}
