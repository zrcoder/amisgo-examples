package main

import (
	"net/http"

	"github.com/zrcoder/amisgo-examples/gop-playground/ex"
	"github.com/zrcoder/amisgo-examples/gop-playground/static"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	examples, defaultExample, err := ex.Get()
	if err != nil {
		panic(err)
	}

	index := comp.Page().Body(
		comp.Form().WrapWithPanel(false).Actions().
			Body(
				comp.Group().Mode("inline").Body(
					comp.Image().Alt("Go+").Src("/static/gop.svg").Height("20px").InnerClassName("border-none"),
					comp.Button().Primary(true).Label("Run").Transform("body", "result", "Done", func(input any) (any, error) {
						return compile(input.(string))
					}),
					comp.Select().Name("examples").Value(defaultExample).Options(
						examples...,
					),
				),
				comp.Editor().Language("c").Name("body").Size("xl").Value("${examples}").AllowFullscreen(false),
				comp.Code().Name("result").Language("plaintext"),
			),
	)

	ag := amisgo.New().Mount("/", index).StaticFS("/static/", http.FS(static.FS))

	panic(ag.Run())
}
