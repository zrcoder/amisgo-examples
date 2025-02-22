package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zrcoder/amisgo-examples/gop-playground/example"
	"github.com/zrcoder/amisgo-examples/gop-playground/static"

	sdk "gitee.com/rdor/amis-sdk/v6"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/conf"
	"github.com/zrcoder/amisgo/schema"
)

var app *amisgo.App

func main() {
	options := []conf.Option{
		conf.WithTitle("Goplus Playground"),
		conf.WithIcon("/static/gop.svg"),
	}
	if os.Getenv("DEV") != "" {
		options = append(options, conf.WithLocalSdk(http.FS(sdk.FS)))
	}

	app = amisgo.New(options...)
	app.Mount("/", index())
	app.StaticFS("/static", http.FS(static.FS))
	fmt.Println("serving on http://localhost:8080")
	err := app.Run(":8080")
	check(err)
}

func index() comp.Page {
	examples, defaultExample, err := example.Get()
	check(err)
	return app.Page().Body(
		app.Form().WrapWithPanel(false).Body(
			app.Flex().Justify("space-between").Items(
				app.Group().Mode("inline").Body(
					app.Image().Alt("Go+").Src("/static/gop.svg").Height("20px").InnerClassName("border-none"),
					app.InputGroup().Body(
						app.Button().Primary(true).Label("Run").TransformMultiple(func(s schema.Schema) (schema.Schema, error) {
							res, err := compile(s.Get("body").(string))
							if err != nil {
								return schema.Schema{"result": "❌ " + err.Error()}, nil
							}
							return schema.Schema{"result": res}, nil
						}, "body"),
						app.Button().Primary(true).Label("Format").Transform(func(input any) (any, error) {
							return format(input.(string))
						}, "body", "body"),
					),
					app.Select().Name("examples").Value(defaultExample).Options(
						examples...,
					),
				),
				app.Button().Label("Github").ActionType("url").Icon("fa fa-github").Url("https://github.com/goplus/gop"),
			),
			app.Editor().Language("c").Name("body").Size("xxl").Value("${examples}").
				AllowFullscreen(false).Options(schema.Schema{"fontSize": 15}),
			app.Code().Name("result").Language("plaintext"),
		),
	)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
