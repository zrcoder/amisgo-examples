package main

import (
	"log"
	"net/http"
	"os"

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
	options := []conf.Option{
		conf.WithTitle("Dev Toys"),
		conf.WithThemes(conf.ThemeCxd, conf.ThemeDark, conf.ThemeAntd, conf.ThemeAng),
		conf.WithIcon("/static/favicon.ico"),
	}
	if os.Getenv("DEV") != "" {
		options = append(options, conf.WithLocalSdk(http.FS(sdk.FS)))
	}
	app := amisgo.New(options...)
	app.StaticFS("/static", http.FS(assets.FS))
	app.Redirect("/", fmtPath, http.StatusPermanentRedirect)
	app.HandleFunc(healthCheckPath, healthz)
	app.Mount(fmtPath, formatPage(app))
	app.Mount(convPath, convPage(app))
	app.Mount(genPath, genPage(app))
	app.Mount(chartPath, chartPage(app))
	app.Mount(encDecPath, encDecPage(app))

	port := ":8888"
	log.Printf("Starting server on http://localhost%s\n", port)

	if err := app.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
