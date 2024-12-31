package main

import (
	"log"
	"net/http"

	"github.com/zrcoder/amisgo-examples/dev-toys/assets"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

const (
	healthCheckPath = "/healthz"
)

const (
	fmtPath    = "/fmt"
	convPath   = "/conv"
	genPath    = "/gen"
	chartPath  = "/chart"
	encDecPath = "/enc"
)

func main() {
	ag := amisgo.New(
		conf.WithTitle("Dev Toys"),
		conf.WithTheme(conf.ThemeDark),
		conf.WithIcon("/static/favicon.ico"),
	).
		StaticFS("/static", http.FS(assets.FS)).
		Redirect("/", fmtPath, http.StatusPermanentRedirect).
		Mount(fmtPath, formatPage).
		Mount(convPath, convPage).
		Mount(genPath, genPage).
		Mount(chartPath, chartPage).
		Mount(encDecPath, encDecPage).
		HandleFunc(healthCheckPath, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

	port := ":8888"
	log.Printf("Starting server on http://localhost%s\n", port)

	if err := ag.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
