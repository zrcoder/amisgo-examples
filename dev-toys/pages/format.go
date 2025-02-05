package pages

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
	"github.com/zrcoder/amisgo-examples/dev-toys/util"
)

func JsonFormatter(app *amisgo.App) any {
	return comp.DualEditor(app, jsonCfg, jsonCfg, "Json", func(input any) (any, error) {
		return util.Json((input.(string)))
	}, nil)
}

func YamlFormatter(app *amisgo.App) any {
	return comp.DualEditor(app, yamlCfg, yamlCfg, "Yaml", func(input any) (any, error) {
		return util.Yaml((input.(string)))
	}, nil)
}

func TomlFormatter(app *amisgo.App) any {
	return comp.DualEditor(app, tomlCfg, tomlCfg, "Toml", func(input any) (any, error) {
		return util.Toml((input.(string)))
	}, nil)
}

func HtmlFormatter(app *amisgo.App) any {
	return comp.DualEditor(app, htmlCfg, htmlCfg, "Html", func(input any) (any, error) {
		return util.Html((input.(string)))
	}, nil)
}
