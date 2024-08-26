package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	app := comp.App().BrandName("DEV TOYS ⌘◉⎈").
		Header(
			comp.Action().ActionType("url").Icon("fa fa-github").Link("https://github.com/zrcoder/amisgo").Label("amisgo"),
		).
		Pages(
			comp.PageItem().Children(
				comp.PageItem().Label("DATA").Icon("fa fa-code").Url("/data").Children(
					comp.PageItem().Label("Format").Url("fmt").Schema(comp.Page().Title("Page A")),
					comp.PageItem().Label("Convert").Url("cvt").Schema(comp.Page().Title("Page B")),
				),
			),
			comp.PageItem().Label("Others").Children(),
		)

	cfg := amisgo.GetDefaultConfig()
	cfg.Theme = amisgo.ThemeDark

	panic(amisgo.ListenAndServe(app, cfg))
}
