package pages

import (
	"dtoy/comp"
	"dtoy/util"
)

var (
	JsonFormatter = comp.DualEditor(jsonCfg, jsonCfg, "Json", func(input any) (any, error) {
		return util.Json((input.(string)))
	}, nil)
	YamlFormatter = comp.DualEditor(yamlCfg, yamlCfg, "Yaml", func(input any) (any, error) {
		return util.Yaml((input.(string)))
	}, nil)
	TomlFormatter = comp.DualEditor(tomlCfg, tomlCfg, "Toml", func(input any) (any, error) {
		return util.Toml((input.(string)))
	}, nil)
	HtmlFormatter = comp.DualEditor(htmlCfg, htmlCfg, "Html", func(input any) (any, error) {
		return util.Html((input.(string)))
	}, nil)
)
