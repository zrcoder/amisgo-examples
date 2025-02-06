package ui

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
	"github.com/zrcoder/amisgo-examples/dev-toys/util"
)

type Formatters struct {
	*amisgo.App
}

func NewFormatters(app *amisgo.App) *Formatters { return &Formatters{App: app} }

func (f *Formatters) JsonFormatter() any {
	return comp.DualEditor(f.App, jsonCfg, jsonCfg, "Json", func(input any) (any, error) {
		return util.Json((input.(string)))
	}, nil)
}

func (f *Formatters) YamlFormatter() any {
	return comp.DualEditor(f.App, yamlCfg, yamlCfg, "Yaml", func(input any) (any, error) {
		return util.Yaml((input.(string)))
	}, nil)
}

func (f *Formatters) TomlFormatter() any {
	return comp.DualEditor(f.App, tomlCfg, tomlCfg, "Toml", func(input any) (any, error) {
		return util.Toml((input.(string)))
	}, nil)
}

func (f *Formatters) HtmlFormatter() any {
	return comp.DualEditor(f.App, htmlCfg, htmlCfg, "Html", func(input any) (any, error) {
		return util.Html((input.(string)))
	}, nil)
}
