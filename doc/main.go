package main

import (
	"log"
	"net/http"

	"github.com/zrcoder/amisgo-examples/doc/static"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/config"
)

func main() {
	menu, err := Menu()
	if err != nil {
		log.Panic(err)
	}
	index := Page(menu, Doc())
	ag := amisgo.New(
		config.WithTheme(config.ThemeAntd),
	).
		StaticFS("/static/", http.FS(static.FS)).
		Mount("/", index).
		HandleFunc(docsApi, getDoc)

	log.Println("Serving on http://localhost:6789")
	log.Fatal(ag.Run(":6789"))
}
