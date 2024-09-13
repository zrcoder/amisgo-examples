package main

import (
	"embed"
	"fmt"

	"amisgo-examples/dev-toys/comp/chart"

	"github.com/zrcoder/amisgo"
	ac "github.com/zrcoder/amisgo/comp"
)

//go:embed assets/*
var assetsFS embed.FS

func main() {
	app := ac.App().
		Logo("/assets/gopher.svg").
		BrandName("Dev Toys").
		Header(
			ac.Flex().Justify("end").Style(ac.Schema{"width": "100%", "padding-right": "50px"}).Items(
				ac.Action().ActionType("url").Icon("fa fa-github").Link("https://github.com/zrcoder/amisgo").Label("amisgo"),
			),
		).
		Pages(
			ac.PageItem().Url("/").Redirect("/fmt/json").Children(
				ac.PageItem().Label("Fommaters").Icon("fa fa-laptop-code").Url("/fmt").Children(
					ac.PageItem().Label("Json").Url("json").Schema(wrap(jsonFormatter)).IsDefaultPage(true),
					ac.PageItem().Label("Yaml").Url("yaml").Schema(wrap(yamlFormatter)),
					ac.PageItem().Label("Toml").Url("toml").Schema(wrap(tomlFormatter)),
					ac.PageItem().Label("Html").Url("html").Schema(wrap(htmlFormatter)),
				),
				ac.PageItem().Label("Converters").Icon("fa fa-right-left").Url("/conv").Children(
					ac.PageItem().Label("Json-Yaml").Url("json-yaml").Schema(wrap(jsonYamlCvt)),
					ac.PageItem().Label("Json-Toml").Url("json-toml").Schema(wrap(jsonTomlCvt)),
					ac.PageItem().Label("Yaml-Toml").Url("yaml-toml").Schema(wrap(yamlTomlCvt)),
				),
				ac.PageItem().Label("Generators").Icon("fa fa-seedling").Url("/gen").Children(
					ac.PageItem().Label("Json Graph").Url("js-graph").Schema(wrap(jsonGraph)),
					ac.PageItem().Label("Qrcode").Url("qrcode").Schema(wrap(qrcode)),
					ac.PageItem().Label("Json to Struct").Url("js-struct").Schema(wrap(json2struct)),
					ac.PageItem().Label("Hash").Url("hash").Schema(wrap(hash)),
					ac.PageItem().Label("Ndor").Url("ndor").Schema(wrap(ndor)),
				),
				ac.PageItem().Label("Chart").Icon("fa fa-bar-chart").Url("/chart").Schema(wrap(chart.Common)).Children(
					ac.PageItem().Label("Line").Url("line").Schema(wrap(chart.Line)),
					ac.PageItem().Label("Bar").Url("bar").Schema(wrap(chart.Bar)),
					ac.PageItem().Label("Pie").Url("pie").Schema(wrap(chart.Pie)),
					ac.PageItem().Label("Scatter").Url("scatter").Schema(wrap(chart.Scatter)),
					ac.PageItem().Label("Polar").Url("polar").Schema(wrap(chart.Polar)),
					ac.PageItem().Label("Radar").Url("radar").Schema(wrap(chart.Radar)),
				),
			))

	cfg := amisgo.GetDefaultConfig()
	cfg.Theme = amisgo.ThemeDark
	cfg.Lang = amisgo.LangEn
	cfg.StaticDir = "assets"
	cfg.StaticFS = assetsFS
	cfg.Icon = "/assets/favicon.ico"

	fmt.Println("Serve on http://localhost")
	panic(amisgo.ListenAndServe(app, cfg))
}

func wrap(c any) any {
	return ac.Wrapper().Style(ac.Schema{"padding": "50px"}).Body(c)
}
