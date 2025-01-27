package main

import (
	"net/http"
	"os"

	"github.com/zrcoder/amisgo-examples/gop-playground/example"
	"github.com/zrcoder/amisgo-examples/gop-playground/static"

	"gitee.com/rdor/amis-sdk/sdk"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/conf"
	"github.com/zrcoder/amisgo/model"
)

func main() {
	options := []conf.Option{
		conf.WithTitle("Goplus Playground"),
		conf.WithIcon("/static/gop.svg"),
	}
	if os.Getenv("DEV") == "1" {
		options = append(options, conf.WithLocalSdk(http.FS(sdk.FS)))
	}

	app := amisgo.New(options...).
		Mount("/", index()).
		StaticFS("/static", http.FS(static.FS))

	err := app.Run()
	check(err)
}

func index() any {
	examples, defaultExample, err := example.Get()
	check(err)
	return comp.Page().Body(
		comp.Form().WrapWithPanel(false).Body(
			comp.Flex().Justify("space-between").Items(
				comp.Group().Mode("inline").Body(
					comp.Image().Alt("Go+").Src("/static/gop.svg").Height("20px").InnerClassName("border-none"),
					comp.InputGroup().Body(
						comp.Button().Primary(true).Label("Run").Transform(func(input any) (any, error) {
							return compile(input.(string))
						}, "body", "result"),
						comp.Button().Primary(true).Label("Format").Transform(func(input any) (any, error) {
							return format(input.(string))
						}, "body", "body"),
					),
					comp.Select().Name("examples").Value(defaultExample).Options(
						examples...,
					),
				),
				comp.Button().Label("Github").ActionType("url").Icon("fa fa-github").Url("https://github.com/goplus/gop"),
			),
			comp.Editor().Language("c").Name("body").Size("xxl").Value("${examples}").
				AllowFullscreen(false).Options(model.Schema{"fontSize": 15}),
			comp.Code().Name("result").Language("plaintext"),
		),
	)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
