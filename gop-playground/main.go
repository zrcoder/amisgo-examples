package main

import (
	"net/http"

	"github.com/zrcoder/amisgo-examples/gop-playground/example"
	"github.com/zrcoder/amisgo-examples/gop-playground/static"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	examples, defaultExample, err := example.Get()
	check(err)

	index := comp.Page().Body(
		comp.Form().WrapWithPanel(false).Body(
			comp.Group().Mode("inline").Body(
				comp.Image().Alt("Go+").Src("/static/gop.svg").Height("20px").InnerClassName("border-none"),
				comp.InputGroup().Body(
					comp.Button().Primary(true).Label("Run").Transform("body", "result", "Done", func(input any) (any, error) {
						return compile(input.(string))
					}),
					comp.Button().Primary(true).Label("Format").Transform("body", "body", "Done", func(input any) (any, error) {
						return format(input.(string))
					}),
				),
				comp.Select().Name("examples").Value(defaultExample).Options(
					examples...,
				),
			),
			comp.Editor().Language("c").Name("body").Size("xl").Value("${examples}").
				AllowFullscreen(false).Options(comp.Schema{"fontSize": 15}),
			comp.Code().Name("result").Language("plaintext"),
		),
	)

	ag := amisgo.New().Mount("/", index).StaticFS("/static/", http.FS(static.FS))

	err = ag.Run()
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}