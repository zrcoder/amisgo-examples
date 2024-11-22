package main

import (
	"log"
	"net/http"

	"dtoy/assets"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/config"
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
		config.WithTheme(config.ThemeDark),
		config.WithStaticFS("/static/", http.FS(assets.FS)),
		config.WithIcon("/static/favicon.ico"),
	).
		Redirect("/", fmtPath).
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
