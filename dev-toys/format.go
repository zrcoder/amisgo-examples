package main

import (
	"dtoy/comp"
	"dtoy/util"
)

var (
	jsonFormatter = comp.DualEditor(jsonCfg, jsonCfg, "Json", func(input any) (any, error) {
		return util.Json((input.(string)))
	}, nil)
	yamlFormatter = comp.DualEditor(yamlCfg, yamlCfg, "Yaml", func(input any) (any, error) {
		return util.Yaml((input.(string)))
	}, nil)
	tomlFormatter = comp.DualEditor(tomlCfg, tomlCfg, "Toml", func(input any) (any, error) {
		return util.Toml((input.(string)))
	}, nil)
	htmlFormatter = comp.DualEditor(htmlCfg, htmlCfg, "Html", func(input any) (any, error) {
		return util.Html((input.(string)))
	}, nil)
)
