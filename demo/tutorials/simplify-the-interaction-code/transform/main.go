package main

import (
	"github.com/zrcoder/amisgo"
)

func main() {
	app := amisgo.New()
	index := app.Page().Body(
		app.Form().WrapWithPanel(false).Body(
			app.InputText().Name("input"),
			app.InputText().Name("output").ReadOnly(true),
			app.Action().Label("Greet").
				Level("primary").
				Transform(func(input any) (any, error) {
					return "hello " + input.(string), nil
				}, "input", "output"),
		),
	)
	app.Mount("/", index)

	app.Run(":8888")
}
