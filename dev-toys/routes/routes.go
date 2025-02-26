package routes

import (
	"net/http"
	"os"

	sdk "gitee.com/rdor/amis-sdk/v6"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/dev-toys/assets"
	"github.com/zrcoder/amisgo-examples/dev-toys/routes/path"
	"github.com/zrcoder/amisgo-examples/dev-toys/ui"
	"github.com/zrcoder/amisgo/conf"
)

func Setup() *amisgo.App {
	options := []conf.Option{
		conf.WithTitle("Dev Toys"),
		conf.WithThemes(
			conf.Theme{Value: conf.ThemeDark},
			conf.Theme{Value: conf.ThemeCxd},
			conf.Theme{Value: conf.ThemeAntd},
		),
		conf.WithIcon("/static/favicon.ico"),
	}
	if os.Getenv("DEV") != "" {
		options = append(options, conf.WithLocalSdk(http.FS(sdk.FS)))
	}
	app := amisgo.New(options...)
	app.StaticFS("/static", http.FS(assets.FS))
	app.Redirect("/", path.Fmt, http.StatusPermanentRedirect)
	app.HandleFunc(path.HealthCheck, healthz)

	ui := ui.New(app)
	app.Mount(path.Fmt, ui.FormatPage())
	app.Mount(path.Conv, ui.ConvPage())
	app.Mount(path.Gen, ui.GenPage())
	app.Mount(path.Chart, ui.ChartPage())
	app.Mount(path.EncDec, ui.EncDecPage())
	return app
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
