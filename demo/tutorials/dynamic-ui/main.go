package main

import "github.com/zrcoder/amisgo"

var app *amisgo.App

func main() {
	app = amisgo.New()
	index := app.Page().Body(
		app.Service().Name("ui").GetData(func() (any, error) {
			return map[string]any{
				"ui": getDynamicUI(),
			}, nil
		}).Body(
			app.Amis().Name("ui"),
		),
	)
	app.Mount("/", index)
	app.Run(":8080")
}

func getDynamicUI() any {
	return app.Tpl().Tpl("Hello, world!")
}
