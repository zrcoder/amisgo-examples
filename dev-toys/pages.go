package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/dev-toys/pages"
)

func formatPage(app *amisgo.App) any {
	return page(app, getFormatters(app))
}

func convPage(app *amisgo.App) any {
	return page(app, getConverters(app))
}

func genPage(app *amisgo.App) any {
	return page(app, getGenerators(app))
}

func chartPage(app *amisgo.App) any {
	return page(app, getCharts(app))
}

func encDecPage(app *amisgo.App) any {
	return page(app, getEncoders(app))
}

func page(app *amisgo.App, content any) any {
	return app.Page().
		Aside(getNav(app), app.ThemeSelect().Label("Theme").ClassName("px-5")).
		AsideClassName("w-56").
		AsideResizor(true).
		Body(content)
}

func getNav(app *amisgo.App) any {
	return app.Nav().Stacked(true).Links(
		navLink(app, "Dev Toys", "fa fa-home", "/"),
		app.NavLink().Mode("divider"),
		navLink(app, "Formatters", "fa fa-laptop-code", fmtPath),
		navLink(app, "Converters", "fa fa-right-left", convPath),
		navLink(app, "Generators", "fa fa-seedling", genPath),
		navLink(app, "Charts", "fa fa-bar-chart", chartPath),
		navLink(app, "Encoders/Decoders", "fa fa-code", encDecPath),
		app.NavLink().Mode("divider"),
		navExtraLink(app, "amisgo", "fa fa-github", "https://github.com/zrcoder/amisgo"),
		navExtraLink(app, "Ndor", "fa fa-image", "https://ndor.netlify.app"),
		app.NavLink().Mode("divider"),
	)
}

func navLink(app *amisgo.App, label, icon, path string) any {
	return app.NavLink().Label(label).Icon(icon).To(path)
}

func navExtraLink(app *amisgo.App, label, icon, path string) any {
	return app.NavLink().Label(label).Icon(icon).To(path).Target("_blank")
}

func getFormatters(app *amisgo.App) any {
	return genTabs(
		app,
		genTab(app, "Json", pages.JsonFormatter(app)),
		genTab(app, "Yaml", pages.YamlFormatter(app)),
		genTab(app, "Toml", pages.TomlFormatter(app)),
		genTab(app, "Html", pages.HtmlFormatter(app)),
	)
}

func getConverters(app *amisgo.App) any {
	return genTabs(
		app,
		genTab(app, "Json-Yaml", pages.JsonYamlCvt(app)),
		genTab(app, "Yaml-Toml", pages.YamlTomlCvt(app)),
		genTab(app, "Json-Toml", pages.JsonTomlCvt(app)),
	)
}

func getGenerators(app *amisgo.App) any {
	return genTabs(
		app,
		genTab(app, "Json Graph", pages.JsonGraph(app)),
		genTab(app, "Qrcoder", pages.Qrcode(app)),
		genTab(app, "Json to Struct", pages.Json2struct(app)),
		genTab(app, "Hash", pages.Hash(app)),
	)
}

func getCharts(app *amisgo.App) any {
	return genTabs(
		app,
		genTab(app, "Line", pages.LineChart(app)),
		genTab(app, "Bar", pages.BarChart(app)),
		genTab(app, "Scatter", pages.ScatterChart(app)),
		genTab(app, "Polar", pages.PolarChart(app)),
		genTab(app, "Pie", pages.PieChart(app)),
		genTab(app, "Radar", pages.RadarChart(app)),
		genTab(app, "DIY", pages.DiyChart(app)),
	)
}

func getEncoders(app *amisgo.App) any {
	return genTabs(
		app,
		genTab(app, "Base64", pages.Base64ED(app)),
		genTab(app, "Url", pages.UrlED(app)),
		genTab(app, "Html", pages.HtmlED(app)),
		genTab(app, "Qrcode Decoder", pages.Decqr(app)),
	)
}

func genTabs(app *amisgo.App, tabs ...any) any {
	return app.Tabs().
		TabsMode("simple").
		Swipeable(true).
		Tabs(tabs...)
}

func genTab(app *amisgo.App, title string, page any) any {
	return app.Tab().Title(title).Tab(page)
}
