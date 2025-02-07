package main

import (
	"log"
	"net/http"
	"os"

	"github.com/zrcoder/amisgo-examples/dev-toys/assets"
	"github.com/zrcoder/amisgo-examples/dev-toys/routes"
	"github.com/zrcoder/amisgo-examples/dev-toys/ui"

	"gitee.com/rdor/amis-sdk/sdk"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
	"github.com/zrcoder/amisgo/theme"
)

func main() {
	options := []conf.Option{
		conf.WithTitle("Dev Toys"),
		conf.WithThemes(
			theme.Theme{Value: theme.Dark},
			theme.Theme{Value: theme.Cxd},
			theme.Theme{Value: theme.Antd},
		),
		conf.WithIcon("/static/favicon.ico"),
	}
	if os.Getenv("DEV") != "" {
		options = append(options, conf.WithLocalSdk(http.FS(sdk.FS)))
	}
	app := amisgo.New(options...)
	app.StaticFS("/static", http.FS(assets.FS))
	app.Redirect("/", routes.Fmt, http.StatusPermanentRedirect)
	app.HandleFunc(routes.HealthCheck, healthz)

	ui := ui.New(app)
	app.Mount(routes.Fmt, ui.FormatPage())
	app.Mount(routes.Conv, ui.ConvPage())
	app.Mount(routes.Gen, ui.GenPage())
	app.Mount(routes.Chart, ui.ChartPage())
	app.Mount(routes.EncDec, ui.EncDecPage())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	log.Printf("Starting server on http://localhost:%s\n", port)

	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
