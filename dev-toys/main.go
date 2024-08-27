package main

import (
	"embed"

	"github.com/zrcoder/amisgo"
	ac "github.com/zrcoder/amisgo/comp"
)

//go:embed asserts/*
var assertsFS embed.FS

func main() {
	app := ac.App().
		BrandName("⎈ DEV TOYS ⎈").
		Header(
			ac.Action().ActionType("url").Icon("fa fa-github").Link("https://github.com/zrcoder/amisgo").Label("amisgo"),
		).
		Pages(
			ac.PageItem().Url("/").Redirect("/data/fmt").Children(
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
					ac.PageItem().Label("Json viewer").Url("view").Schema(wrap(jsonViewer)),
					ac.PageItem().Label("Qrcode").Url("qrcode").Schema(wrap(qrcode)),
				),
			),
		)

	cfg := amisgo.GetDefaultConfig()
	cfg.Theme = amisgo.ThemeDark
	cfg.Lang = amisgo.LangEn
	cfg.Icon = "/favicon.png"
	cfg.AssertsPath = "/asserts"
	cfg.AssertsFS = assertsFS

	panic(amisgo.ListenAndServe(app, cfg))
}

func wrap(c any) any {
	return ac.Wrapper().Body(c)
}
