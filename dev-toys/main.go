package main

import (
	"embed"

	"github.com/zrcoder/amisgo"
	ac "github.com/zrcoder/amisgo/comp"
)

//go:embed asserts/*
var assertsFS embed.FS

func main() {
	app := ac.App().
		BrandName("DEV TOYS ⌘◉⎈").
		Header(
			ac.Action().ActionType("url").Icon("fa fa-github").Link("https://github.com/zrcoder/amisgo").Label("amisgo"),
		).
		Pages(
			ac.PageItem().Url("/").Redirect("/data/fmt").Children(
				ac.PageItem().Label("DATA").Icon("fa fa-code").Url("/data").Children(
					ac.PageItem().Label("Format").Url("fmt").Schema(fmtEditor).IsDefaultPage(true),
					ac.PageItem().Label("Convert").Url("cvt").Schema(ac.Page().Title("Page B")),
					ac.PageItem().Label("View").Url("view").Schema(jsonViewer),
				),
				ac.PageItem().Label("TEXT").Url("/text").Children(
					ac.PageItem().Label("Difference").Url("diff").Schema(diff),
				),
			),
		).
		Footer("© 2024 zrcoder.")

	cfg := amisgo.GetDefaultConfig()
	cfg.Theme = amisgo.ThemeDark
	cfg.Lang = amisgo.LangEn
	cfg.Icon = "/favicon.png"
	cfg.AssertsPath = "/asserts"
	cfg.AssertsFS = assertsFS

	panic(amisgo.ListenAndServe(app, cfg))
}
