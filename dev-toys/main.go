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
	var (
		formatters = genTabs(
			ac.Tab().Title("Json").Tab(jsonFormatter),
			ac.Tab().Title("Yaml").Tab(yamlFormatter),
			ac.Tab().Title("Toml").Tab(tomlFormatter),
			ac.Tab().Title("Html").Tab(htmlFormatter),
		)
		converters = genTabs(
			ac.Tab().Title("Json-Yaml").Tab(jsonYamlCvt),
			ac.Tab().Title("Yaml-Toml").Tab(yamlTomlCvt),
			ac.Tab().Title("Json-Toml").Tab(jsonTomlCvt),
		)
		generaters = genTabs(
			ac.Tab().Title("Json Graph").Tab(jsonGraph),
			ac.Tab().Title("Qrcoder").Tab(qrcode),
			ac.Tab().Title("Json to Struct").Tab(json2struct),
			ac.Tab().Title("Hash").Tab(hash),
			ac.Tab().Title("Ndor").Tab(ndor),
		)
		charts = genTabs(
			ac.Tab().Title("Common").Tab(chart.Common),
			ac.Tab().Title("Line").Tab(chart.Line),
			ac.Tab().Title("Bar").Tab(chart.Bar),
			ac.Tab().Title("Scatter").Tab(chart.Scatter),
			ac.Tab().Title("Polar").Tab(chart.Polar),
			ac.Tab().Title("Pie").Tab(chart.Pie),
			ac.Tab().Title("Radar").Tab(chart.Radar),
		)
	)

	app := ac.App().
		Logo("/assets/gopher.svg").
		BrandName("Dev Toys").
		Header(
			ac.Flex().Justify("end").Style(ac.Schema{"width": "100%", "padding-right": "50px"}).Items(
				ac.Action().ActionType("url").Icon("fa fa-github").Link("https://github.com/zrcoder/amisgo").Label("amisgo"),
			),
		).
		Pages(
			ac.PageItem().Url("/").Redirect("/fmt").Children(
				ac.PageItem().Label("Fommaters").Icon("fa fa-laptop-code").Url("/fmt").Schema(formatters),
				ac.PageItem().Label("Converters").Icon("fa fa-right-left").Url("/conv").Schema(converters),
				ac.PageItem().Label("Generators").Icon("fa fa-seedling").Url("/gen").Schema(generaters),
				ac.PageItem().Label("Charts").Icon("fa fa-bar-chart").Url("/chart").Schema(charts),
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

func genTabs(tabs ...any) any {
	return wrap(ac.Tabs().TabsMode("radio").Swipeable(true).Tabs(tabs...))
}

func wrap(c any) any {
	return ac.Wrapper().Style(ac.Schema{"padding": "25px"}).Body(c)
}
