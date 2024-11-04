package main

import (
	"embed"
	"log"
	"net/http"

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

//go:embed assets/*
var assetsFS embed.FS

func main() {
	http.Handle("/assets/", http.FileServer(http.FS(assetsFS)))

	http.HandleFunc(healthCheckPath, healthCheck)

	ag := amisgo.New(
		config.WithTheme(config.ThemeDark),
		config.WithLang(config.LangEn),
		config.WithIcon("/assets/favicon.ico"),
	).
		Redirect("/", fmtPath).
		Register(fmtPath, formatPage).
		Register(convPath, convPage).
		Register(genPath, genPage).
		Register(chartPath, chartPage).
		Register(encDecPath, encDecPage)

	port := ":8888"
	log.Printf("Starting server on http://localhost%s\n", port)

	if err := ag.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
