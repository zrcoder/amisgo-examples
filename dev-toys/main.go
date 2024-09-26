package main

import (
	"fmt"
	"net/http"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

const (
	fmtPath        = "/fmt"
	convPath       = "/conv"
	genPath        = "/gen"
	chartPath      = "/chart"
	encDecPath     = "/enc"
	helthCheckPath = "/healthz"
)

var (
	nav = comp.Nav().Stacked(true).Links(
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

func main() {
	amisgo.Redirect("/", fmtPath)
	amisgo.Serve(fmtPath, page(formatters))
	amisgo.Serve(convPath, page(converters))
	amisgo.Serve(genPath, page(generaters))
	amisgo.Serve(chartPath, page(charts))
	amisgo.Serve(encDecPath, page(encoders))

	// helth check endpoint, just for render now.
	http.HandleFunc(helthCheckPath, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Serving on http://localhost")
	panic(amisgo.ListenAndServe(":80", appConfig))
}

func page(content any) any {
	return comp.Page().Aside(nav).AsideClassName("w-56").AsideResizor(true).Body(content)
}

func navLink(label, icon, path string) any {
	return comp.NavLink().Label(label).Icon(icon).To(path)
}

func genTabs(tabs ...any) any {
	return (comp.Tabs().TabsMode("simple").Swipeable(true).Tabs(tabs...))
}

func genTab(title string, page any) any {
	return comp.Tab().Title(title).Tab(page)
}
