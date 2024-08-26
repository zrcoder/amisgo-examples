package main

import (
	"bytes"

	"amisgo-examples/dev-toys/comp"

	"github.com/zrcoder/ttoy/pkg/formatter"
)

var (
	fmts = map[string]func([]byte) (*bytes.Buffer, error){
		"json": formatter.Json,
		"yaml": formatter.Yaml,
		"toml": formatter.Toml,
		"html": formatter.Html,
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
