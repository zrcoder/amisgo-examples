package ui

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/dev-toys/routes"
)

type UI struct {
	*amisgo.App
}

func New(app *amisgo.App) *UI {
	return &UI{App: app}
}

func (u *UI) FormatPage() any {
	return u.page(u.getFormatters())
}

func (u *UI) ConvPage() any {
	return u.page(u.getConverters())
}

func (u *UI) GenPage() any {
	return u.page(u.getGenerators())
}

func (u *UI) ChartPage() any {
	return u.page(u.getCharts())
}

func (u *UI) EncDecPage() any {
	return u.page(u.getEncoders())
}

func (u *UI) page(content any) any {
	return u.Page().
		Aside(u.getNav(), u.ThemeSelect().Label("Theme").ClassName("px-5")).
		AsideClassName("w-56").
		AsideResizor(true).
		Body(content)
}

func (u *UI) getNav() any {
	return u.Nav().Stacked(true).Links(
		u.navLink("Dev Toys", "fa fa-home", "/"),
		u.NavLink().Mode("divider"),
		u.navLink("Formatters", "fa fa-laptop-code", routes.Fmt),
		u.navLink("Converters", "fa fa-right-left", routes.Conv),
		u.navLink("Generators", "fa fa-seedling", routes.Gen),
		u.navLink("Charts", "fa fa-bar-chart", routes.Chart),
		u.navLink("Encoders/Decoders", "fa fa-code", routes.EncDec),
		u.NavLink().Mode("divider"),
		u.navExtraLink("amisgo", "fa fa-github", "https://github.com/zrcoder/amisgo"),
		u.navExtraLink("Ndor", "fa fa-image", "https://ndor.netlify.app"),
		u.NavLink().Mode("divider"),
	)
}

func (u *UI) navLink(label, icon, path string) any {
	return u.NavLink().Label(label).Icon(icon).To(path)
}

func (u *UI) navExtraLink(label, icon, path string) any {
	return u.NavLink().Label(label).Icon(icon).To(path).Target("_blank")
}

func (u *UI) getFormatters() any {
	fmts := NewFormatters(u.App)
	return u.genTabs(
		u.genTab("Json", fmts.JsonFormatter()),
		u.genTab("Yaml", fmts.YamlFormatter()),
		u.genTab("Toml", fmts.TomlFormatter()),
		u.genTab("Html", fmts.HtmlFormatter()),
	)
}

func (u *UI) getConverters() any {
	cvt := NewConverters(u.App)
	return u.genTabs(
		u.genTab("Json-Yaml", cvt.JsonYamlCvt()),
		u.genTab("Yaml-Toml", cvt.YamlTomlCvt()),
		u.genTab("Json-Toml", cvt.JsonTomlCvt()),
	)
}

func (u *UI) getGenerators() any {
	gens := NewGenerators(u.App)
	return u.genTabs(
		u.genTab("Json Viewer", gens.JsonViewer()),
		u.genTab("Json Graph", gens.JsonGraph()),
		u.genTab("Json to Struct", gens.Json2struct()),
		u.genTab("Hash", gens.Hash()),
		u.genTab("Qrcoder", gens.Qrcode()),
	)
}

func (u *UI) getCharts() any {
	chart := NewChart(u.App)
	return u.genTabs(
		u.genTab("Line", chart.Line()),
		u.genTab("Bar", chart.Bar()),
		u.genTab("Scatter", chart.Scatter()),
		u.genTab("Polar", chart.Polar()),
		u.genTab("Pie", chart.Pie()),
		u.genTab("Radar", chart.Radar()),
		u.genTab("DIY", chart.Diy()),
	)
}

func (u *UI) getEncoders() any {
	enc := NewEncoders(u.App)
	return u.genTabs(
		u.genTab("Base64", enc.Base64ED()),
		u.genTab("Url", enc.UrlED()),
		u.genTab("Html", enc.HtmlED()),
		u.genTab("Qrcode Decoder", enc.Decqr()),
	)
}

func (u *UI) genTabs(tabs ...any) any {
	return u.Tabs().TabsMode("simple").Swipeable(true).Tabs(tabs...)
}

func (u *UI) genTab(title string, page any) any {
	return u.Tab().Title(title).Tab(page)
}
