package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/schema"
)

func main() {
	app := amisgo.New()
	app.Page().Body(
		app.Form().WrapWithPanel(false).Body(
			app.Flex().Justify("space-between").Items(
				app.Group().Mode("inline").Body(
					app.Image().Alt("Go+").Src("/static/gop.svg").Height("20px").InnerClassName("border-none"),
					app.InputGroup().Body(
						app.Button().Primary(true).Label("Run").Transform(func(input any) (any, error) {
							// TODO
							return nil, nil
						}, "body", "result"),
						app.Button().Primary(true).Label("Format").Transform(func(input any) (any, error) {
							// TODO
							return nil, nil
						}, "body", "body"),
					),
					app.Select().Name("examples").Value("TODO").Options("TODO"),
				),
				app.Button().Label("Github").ActionType("url").Icon("fa fa-github").Url("https://github.com/goplus/gop"),
			),
			app.Editor().Language("c").Name("body").Size("xxl").Value("${examples}").
				AllowFullscreen(false).Options(schema.Schema{"fontSize": 15}),
			app.Code().Name("result").Language("plaintext"),
		),
	)
	app.Run()
}
