package main

import (
	"net/http"

	"github.com/zrcoder/amisgo"
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

func setupRoutes() {
	amisgo.Redirect("/", fmtPath)

	routes := map[string]any{
		fmtPath:    formatPage,
		convPath:   convPage,
		genPath:    genPage,
		chartPath:  chartPage,
		encDecPath: encDecPage,
	}

	for path, handler := range routes {
		amisgo.Serve(path, handler)
	}

	http.HandleFunc(healthCheckPath, healthCheck)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
