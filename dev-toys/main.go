package main

import (
	"log"
	"net/http"

	"github.com/zrcoder/amisgo-examples/dev-toys/assets"

	"gitee.com/rdor/amis-sdk/sdk"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

const (
	fmtPath    = "/fmt"
	convPath   = "/conv"
	genPath    = "/gen"
	chartPath  = "/chart"
	encDecPath = "/enc"

	healthCheckPath = "/healthz"
)

func main() {
	app := amisgo.New(
		conf.WithTitle("Dev Toys"),
		conf.WithThemes(conf.ThemeCxd, conf.ThemeDark, conf.ThemeAntd, conf.ThemeAng),
		conf.WithLocalSdk(http.FS(sdk.FS)),
		conf.WithIcon("/static/favicon.ico"),
	).
		StaticFS("/static", http.FS(assets.FS)).
		Redirect("/", fmtPath, http.StatusPermanentRedirect).
		Mount(fmtPath, formatPage()).
		Mount(convPath, convPage()).
		Mount(genPath, genPage()).
		Mount(chartPath, chartPage()).
		Mount(encDecPath, encDecPage()).
		HandleFunc(healthCheckPath, healthz)

	port := ":8888"
	log.Printf("Starting server on http://localhost%s\n", port)

	if err := app.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
