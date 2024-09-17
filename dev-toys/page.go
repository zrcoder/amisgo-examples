package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func init() {
	amisgo.Serve(fmtPath, page(formatters))
	amisgo.Serve(convPath, page(converters))
	amisgo.Serve(genPath, page(generaters))
	amisgo.Serve(chartPath, page(charts))
	amisgo.Serve(encDecPath, page(encoders))
}

func page(content any) any {
	return comp.Page().
		Aside(nav).
		AsideClassName("w-16").
		AsideMinWidth(80).
		AsideMaxWidth(100).
		AsideResizor(true).
		Body(content)
}

var (
	nav = comp.Nav().Stacked(true).Collapsed(true).Links(
		navLink("Home", "fa fa-home", "/"),
		comp.NavLink().Mode("group").ClassName("pt-20"),
		navLink("Formatters", "fa fa-laptop-code", fmtPath),
		navLink("Converters", "fa fa-right-left", convPath),
		navLink("Generators", "fa fa-seedling", genPath),
		navLink("Charts", "fa fa-bar-chart", chartPath),
		navLink("Encoders/Decoders", "fa fa-code", encDecPath),
		comp.NavLink().Mode("group").ClassName("pb-20"),
		navLink("Amisgo", "fa fa-github", "https://github.com/zrcoder/amisgo"),
		navLink("Ndor", "fa fa-image", "https://ndor.netlify.app"),
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

func navLink(label, icon, path string) any {
	return comp.NavLink().Label(label).Icon(icon).To(path)
}

func genTabs(tabs ...any) any {
	return (comp.Tabs().TabsMode("simple").Swipeable(true).Tabs(tabs...))
}

func genTab(title string, page any) any {
	return comp.Tab().Title(title).Tab(page)
}
