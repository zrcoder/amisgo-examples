package main

import (
	"github.com/zrcoder/amisgo/comp"
)

var (
	header = comp.Flex().Justify("end").Style(comp.Schema{"width": "100%", "padding-right": "22px"}).Items(
		comp.ButtonToolbar().Buttons(
			comp.Action().ActionType("url").Icon("fa fa-github").Link("https://github.com/zrcoder/amisgo").Label("amisgo"),
			comp.Action().ActionType("url").Icon("fa fa-image").Link("https://ndor.netlify.app").Label("ndor"),
		),
	)
	pages = comp.PageItem().Url("/").Redirect("/fmt").Children(
		comp.PageItem().Label("Fommaters").Icon("fa fa-laptop-code").Url("/fmt").Schema(formatters),
		comp.PageItem().Label("Converters").Icon("fa fa-right-left").Url("/conv").Schema(converters),
		comp.PageItem().Label("Generators").Icon("fa fa-seedling").Url("/gen").Schema(generaters),
		comp.PageItem().Label("Charts").Icon("fa fa-bar-chart").Url("/chart").Schema(charts),
		comp.PageItem().Label("Encoders/Decoders").Icon("fa fa-code").Url("/enc").Schema(encoders),
	)
	formatters = genTabs(
		genTab("Json", jsonFormatter),
		genTab("Yaml", yamlFormatter),
		genTab("Toml", tomlFormatter),
		genTab("Html", htmlFormatter),
	)
	converters = genTabs(
		genTab("Json-Yaml", jsonYamlCvt),
		genTab("Yaml-Toml", yamlTomlCvt),
		genTab("Json-Toml", jsonTomlCvt),
	)
	generaters = genTabs(
		genTab("Json Graph", jsonGraph),
		genTab("Qrcoder", qrcode),
		genTab("Json to Struct", json2struct),
		genTab("Hash", hash),
	)
	charts = genTabs(
		genTab("Line", lineChart),
		genTab("Bar", barChart),
		genTab("Scatter", scatterChart),
		genTab("Polar", polarChart),
		genTab("Pie", pieChart),
		genTab("Radar", radarChart),
		genTab("DIY", diyChart),
	)
	encoders = genTabs(
		genTab("Base64", base64ED),
		genTab("Url", urlED),
		genTab("Html", htmlED),
		genTab("Qrcode Decoder", decqr),
	)
)

func wrap(c any) any {
	return comp.Wrapper().Style(comp.Schema{"padding": "25px"}).Body(c)
}

func genTabs(tabs ...any) any {
	return wrap(comp.Tabs().TabsMode("radio").Swipeable(true).Tabs(tabs...))
}

func genTab(title string, page any) any {
	return comp.Tab().Title(styledTabTitle(title)).Tab(page)
}

func styledTabTitle(title string) any {
	return comp.Static().Style(comp.Schema{"margin": "10px"}).Text(title)
}
