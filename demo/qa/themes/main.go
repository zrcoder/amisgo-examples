package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

func main() {
	app := amisgo.New(
		conf.WithThemes(
			conf.Theme{Value: conf.ThemeCxd, Label: "Light"},
			conf.Theme{Value: conf.ThemeDark, Label: "Dark"},
		),
	)
	app.Mount("/", app.Page().Body(
		app.ThemeButtonGroupSelect(),
		"Hello, World!",
	))
	app.Run(":8888")
}
