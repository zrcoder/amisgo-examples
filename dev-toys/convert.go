package main

import (
	"bytes"

	"amisgo-examples/dev-toys/comp"

	"github.com/zrcoder/ttoy/pkg/converter"
)

var jsonYamlCvt = comp.DualEditor(jsonCfg, yamlCfg, "Json-Yaml",
	func(input any) (output any, err error) {
		return convert(input, converter.Json2Yaml)
	}, func(input any) (output any, err error) {
		return convert(input, converter.Yaml2Json)
	})

var jsonTomlCvt = comp.DualEditor(jsonCfg, tomlCfg, "Json-Toml",
	func(input any) (output any, err error) {
		return convert(input, converter.Json2Toml)
	}, func(input any) (output any, err error) {
		return convert(input, converter.Toml2Json)
	})

var yamlTomlCvt = comp.DualEditor(yamlCfg, tomlCfg, "Yaml-Toml",
	func(input any) (output any, err error) {
		return convert(input, converter.Yaml2Toml)
	}, func(input any) (output any, err error) {
		return convert(input, converter.Toml2Yaml)
	})

func convert(input any, cvt func([]byte) (*bytes.Buffer, error)) (string, error) {
	buf, err := cvt([]byte(input.(string)))
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}