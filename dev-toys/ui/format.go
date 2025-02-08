package ui

import (
	"github.com/zrcoder/amisgo-examples/dev-toys/util"
	"github.com/zrcoder/amisgo/comp"
)

func (u *UI) JsonFormatter() comp.Form {
	return u.DualEditor(jsonCfg, jsonCfg, "Json", func(input any) (any, error) {
		return util.Json((input.(string)))
	}, nil)
}

func (u *UI) YamlFormatter() comp.Form {
	return u.DualEditor(yamlCfg, yamlCfg, "Yaml", func(input any) (any, error) {
		return util.Yaml((input.(string)))
	}, nil)
}

func (u *UI) TomlFormatter() comp.Form {
	return u.DualEditor(tomlCfg, tomlCfg, "Toml", func(input any) (any, error) {
		return util.Toml((input.(string)))
	}, nil)
}

func (u *UI) HtmlFormatter() comp.Form {
	return u.DualEditor(htmlCfg, htmlCfg, "Html", func(input any) (any, error) {
		return util.Html((input.(string)))
	}, nil)
}
