package main

import (
	"github.com/zrcoder/amisgo-examples/dev-toys/pages"

	"github.com/zrcoder/amisgo/comp"
)

var nav = initNav()

var (
	formatPage = page(initFormatters())
	convPage   = page(initConverters())
	genPage    = page(initGenerators())
	chartPage  = page(initCharts())
	encDecPage = page(initEncoders())
)

func initNav() any {
	return comp.Nav().Stacked(true).Links(
		navLink("Dev Toys", "fa fa-home", "/"),
		comp.NavLink().Mode("divider"),
		navLink("Formatters", "fa fa-laptop-code", fmtPath),
		navLink("Converters", "fa fa-right-left", convPath),
		navLink("Generators", "fa fa-seedling", genPath),
		navLink("Charts", "fa fa-bar-chart", chartPath),
		navLink("Encoders/Decoders", "fa fa-code", encDecPath),
		comp.NavLink().Mode("divider"),
		navLink("Amisgo", "fa fa-github", "https://github.com/zrcoder/amisgo"),
		navLink("Ndor", "fa fa-image", "https://ndor.netlify.app"),
	)
}

func initFormatters() any {
	return genTabs(
		genTab("Json", pages.JsonFormatter),
		genTab("Yaml", pages.YamlFormatter),
		genTab("Toml", pages.TomlFormatter),
		genTab("Html", pages.HtmlFormatter),
	)
}

func initConverters() any {
	return genTabs(
		genTab("Json-Yaml", pages.JsonYamlCvt),
		genTab("Yaml-Toml", pages.YamlTomlCvt),
		genTab("Json-Toml", pages.JsonTomlCvt),
	)
}

func initGenerators() any {
	return genTabs(
		genTab("Json Graph", pages.JsonGraph),
		genTab("Qrcoder", pages.Qrcode),
		genTab("Json to Struct", pages.Json2struct),
		genTab("Hash", pages.Hash),
	)
}

func initCharts() any {
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

func initEncoders() any {
	return genTabs(
		genTab("Base64", pages.Base64ED),
		genTab("Url", pages.UrlED),
		genTab("Html", pages.HtmlED),
		genTab("Qrcode Decoder", pages.Decqr),
	)
}

func page(content any) any {
	return comp.Page().
		Aside(nav).
		AsideClassName("w-56").
		AsideResizor(true).
		Body(content)
}

func navLink(label, icon, path string) any {
	return comp.NavLink().Label(label).Icon(icon).To(path)
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
