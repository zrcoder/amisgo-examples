package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zrcoder/amisgo-examples/goplay/example"
	"github.com/zrcoder/amisgo-examples/goplay/static"

	sdk "gitee.com/rdor/amis-sdk/v6"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/conf"
	"github.com/zrcoder/amisgo/schema"
)

var app *amisgo.App

func main() {
	options := []conf.Option{
		conf.WithTitle("Go Playground"),
		conf.WithIcon("/static/favicon-gopher.svg"),
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

func index() comp.Wrapper {
	examples, defaultExample, err := example.Get()
	check(err)
	return app.Wrapper().Body(app.Form().WrapWithPanel(false).Body(
		app.Flex().Justify("space-between").Items(
			app.Flex().AlignItems("center").ClassName("h-16").Items(
				app.Image().Alt("Go").Src("/static/go-logo-blue.svg").Height("30px").InnerClassName("border-none"),
				app.Tpl().Tpl("The Go Playground").ClassName("ml-2 text-2xl font-bold"),
			),
			app.Group().Mode("inline").Body(
				app.Wrapper().Body(
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
					app.Select().ClassName("w-48").Overlay(app.Overlay().Width(200)).Name("examples").Value(defaultExample).Options(
						examples...,
					),
					app.Button().Label("Code").ActionType("url").Icon("fa fa-github").Url("https://github.com/zrcoder/amisgo-example"),
				),
			),
		),
		app.Editor().Language("c").Name("body").Size("xxl").Value("${examples}").
			AllowFullscreen(false).Options(schema.Schema{"fontSize": 15}),
		app.Code().Name("result").Language("plaintext"),
	))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
