package main

import (
	"github.com/zrcoder/amisgo-examples/dev-toys/pages"

	"github.com/zrcoder/amisgo/comp"
)

var nav = getNav()

func formatPage() any {
	return page(getFormatters())
}

func convPage() any {
	return page(getConverters())
}

func genPage() any {
	return page(getGenerators())
}

func chartPage() any {
	return page(getCharts())
}

func encDecPage() any {
	return page(getEncoders())
}

func page(content any) any {
	return comp.Page().
		Title("Dev Toys").
		Toolbar(
			comp.ThemeSelect().Name("theme").Label("Theme").Mode("inline"),
		).
		Aside(nav).
		AsideClassName("w-56").
		AsideResizor(true).
		Body(content)
}

func getNav() any {
	return comp.Nav().Stacked(true).Links(
		navLink("", "fa fa-home", "/"),
		comp.NavLink().Mode("divider"),
		navLink("Formatters", "fa fa-laptop-code", fmtPath),
		navLink("Converters", "fa fa-right-left", convPath),
		navLink("Generators", "fa fa-seedling", genPath),
		navLink("Charts", "fa fa-bar-chart", chartPath),
		navLink("Encoders/Decoders", "fa fa-code", encDecPath),
		comp.NavLink().Mode("divider"),
		navExtraLink("amisgo", "fa fa-github", "https://github.com/zrcoder/amisgo"),
		navExtraLink("Ndor", "fa fa-image", "https://ndor.netlify.app"),
	)
}

func navLink(label, icon, path string) any {
	return comp.NavLink().Label(label).Icon(icon).To(path)
}

func navExtraLink(label, icon, path string) any {
	return comp.NavLink().Label(label).Icon(icon).To(path).Target("_blank")
}

func getFormatters() any {
	return genTabs(
		genTab("Json", pages.JsonFormatter),
		genTab("Yaml", pages.YamlFormatter),
		genTab("Toml", pages.TomlFormatter),
		genTab("Html", pages.HtmlFormatter),
	)
}

func getConverters() any {
	return genTabs(
		genTab("Json-Yaml", pages.JsonYamlCvt),
		genTab("Yaml-Toml", pages.YamlTomlCvt),
		genTab("Json-Toml", pages.JsonTomlCvt),
	)
}

func getGenerators() any {
	return genTabs(
		genTab("Json Graph", pages.JsonGraph),
		genTab("Qrcoder", pages.Qrcode),
		genTab("Json to Struct", pages.Json2struct),
		genTab("Hash", pages.Hash),
	)
}

func getCharts() any {
	return genTabs(
		genTab("Line", pages.LineChart),
		genTab("Bar", pages.BarChart),
		genTab("Scatter", pages.ScatterChart),
		genTab("Polar", pages.PolarChart),
		genTab("Pie", pages.PieChart),
		genTab("Radar", pages.RadarChart),
		genTab("DIY", pages.DiyChart),
	)
}

func getEncoders() any {
	return genTabs(
		genTab("Base64", pages.Base64ED),
		genTab("Url", pages.UrlED),
		genTab("Html", pages.HtmlED),
		genTab("Qrcode Decoder", pages.Decqr),
	)
}

func genTabs(tabs ...any) any {
	return comp.Tabs().
		TabsMode("simple").
		Swipeable(true).
		Tabs(tabs...)
}

func genTab(title string, page any) any {
	return comp.Tab().Title(title).Tab(page)
}
