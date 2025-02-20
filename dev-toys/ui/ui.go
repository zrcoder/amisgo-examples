package ui

import (
	"github.com/zrcoder/amisgo"
	ac "github.com/zrcoder/amisgo/comp"

	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
	"github.com/zrcoder/amisgo-examples/dev-toys/comp/chart"
	"github.com/zrcoder/amisgo-examples/dev-toys/routes/path"
)

type UI struct {
	*comp.Comp
	chart  *chart.Chart
	qrData []byte
}

func New(app *amisgo.App) *UI {
	c := comp.New(app)
	return &UI{Comp: c, chart: chart.New(c)}
}

func (u *UI) FormatPage() ac.Page {
	return u.page(u.getFormatters())
}

func (u *UI) ConvPage() ac.Page {
	return u.page(u.getConverters())
}

func (u *UI) GenPage() ac.Page {
	return u.page(u.getGenerators())
}

func (u *UI) ChartPage() ac.Page {
	return u.page(u.getCharts())
}

func (u *UI) EncDecPage() ac.Page {
	return u.page(u.getEncoders())
}

func (u *UI) page(content any) ac.Page {
	return u.Page().
		Aside(u.getNav(), u.ThemeSelect().Label("Theme").ClassName("px-5")).
		AsideClassName("w-56").
		AsideResizor(true).
		Body(content)
}

func (u *UI) getNav() ac.Nav {
	return u.Nav().Stacked(true).Links(
		u.navLink("Dev Toys", "fa fa-home", "/"),
		u.NavLink().Mode("divider"),
		u.navLink("Formatters", "fa fa-laptop-code", path.Fmt),
		u.navLink("Converters", "fa fa-right-left", path.Conv),
		u.navLink("Generators", "fa fa-seedling", path.Gen),
		u.navLink("Charts", "fa fa-bar-chart", path.Chart),
		u.navLink("Encoders/Decoders", "fa fa-code", path.EncDec),
		u.NavLink().Mode("divider"),
		u.navExtraLink("amisgo", "fa fa-github", "https://github.com/zrcoder/amisgo"),
		u.navExtraLink("Ndor", "fa fa-image", "https://ndor.netlify.app"),
		u.NavLink().Mode("divider"),
	)
}

func (u *UI) navLink(label, icon, path string) ac.NavLink {
	return u.NavLink().Label(label).Icon(icon).To(path)
}

func (u *UI) navExtraLink(label, icon, path string) ac.NavLink {
	return u.NavLink().Label(label).Icon(icon).To(path).Target("_blank")
}

func (u *UI) getFormatters() ac.Tabs {
	return u.genTabs(
		u.genTab("Json", u.JsonFormatter()),
		u.genTab("Yaml", u.YamlFormatter()),
		u.genTab("Toml", u.TomlFormatter()),
		u.genTab("Html", u.HtmlFormatter()),
	)
}

func (u *UI) getConverters() ac.Tabs {
	return u.genTabs(
		u.genTab("Json-Yaml", u.JsonYamlCvt()),
		u.genTab("Yaml-Toml", u.YamlTomlCvt()),
		u.genTab("Json-Toml", u.JsonTomlCvt()),
	)
}

func (u *UI) getGenerators() ac.Tabs {
	return u.genTabs(
		u.genTab("Json Viewer", u.JsonViewer()),
		u.genTab("Json Graph", u.JsonGraph()),
		u.genTab("Json to Struct", u.Json2struct()),
		u.genTab("Hash", u.Hash()),
		u.genTab("Qrcoder", u.Qrcode()),
	)
}

func (u *UI) getCharts() ac.Tabs {
	return u.genTabs(
		u.genTab("Line", u.Line()),
		u.genTab("Bar", u.Bar()),
		u.genTab("Scatter", u.Scatter()),
		u.genTab("Polar", u.Polar()),
		u.genTab("Pie", u.Pie()),
		u.genTab("Radar", u.Radar()),
		u.genTab("DIY", u.Diy()),
	)
}

func (u *UI) getEncoders() ac.Tabs {
	return u.genTabs(
		u.genTab("Base64", u.Base64ED()),
		u.genTab("Url", u.UrlED()),
		u.genTab("Html", u.HtmlED()),
		u.genTab("Qrcode Decoder", u.Decqr()),
	)
}

func (u *UI) genTabs(tabs ...any) ac.Tabs {
	return u.Tabs().TabsMode("simple").Swipeable(true).Tabs(tabs...)
}

func (u *UI) genTab(title string, page any) ac.Tab {
	return u.Tab().Title(title).Tab(page)
}
