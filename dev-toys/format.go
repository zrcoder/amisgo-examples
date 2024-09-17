package main

import (
	"bytes"

	"amisgo-examples/dev-toys/comp"
	"amisgo-examples/dev-toys/util"
)

var (
	fmts = map[string]func([]byte) (*bytes.Buffer, error){
		"json": util.Json,
		"yaml": util.Yaml,
		"toml": util.Toml,
		"html": util.Html,
	}
	jsonFormatter = comp.DualEditor(jsonCfg, jsonCfg, "Json", func(input any) (any, error) {
		return format("json", input)
	}, nil)
	yamlFormatter = comp.DualEditor(yamlCfg, yamlCfg, "Yaml", func(input any) (any, error) {
		return format("yaml", input)
	}, nil)
	tomlFormatter = comp.DualEditor(tomlCfg, tomlCfg, "Toml", func(input any) (any, error) {
		return format("toml", input)
	}, nil)
	htmlFormatter = comp.DualEditor(htmlCfg, htmlCfg, "Html", func(input any) (any, error) {
		return format("html", input)
	}, nil)
)

func format(lang string, input any) (any, error) {
	fn := fmts[lang]
	buf, err := fn([]byte(input.(string)))
	if err != nil {
		return nil, err
	}
	return buf.String(), nil
}
