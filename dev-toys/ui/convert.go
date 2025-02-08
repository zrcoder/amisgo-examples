package ui

import (
	"bytes"

	"github.com/zrcoder/amisgo-examples/dev-toys/util"
	"github.com/zrcoder/amisgo/comp"
)

func (u *UI) JsonYamlCvt() comp.Form {
	return u.DualEditor(jsonCfg, yamlCfg, "Json-Yaml",
		func(input any) (output any, err error) {
			return convert(input, util.Json2Yaml)
		}, func(input any) (output any, err error) {
			return convert(input, util.Yaml2Json)
		})
}

func (u *UI) JsonTomlCvt() comp.Form {
	return u.DualEditor(jsonCfg, tomlCfg, "Json-Toml",
		func(input any) (output any, err error) {
			return convert(input, util.Json2Toml)
		}, func(input any) (output any, err error) {
			return convert(input, util.Toml2Json)
		})
}

func (u *UI) YamlTomlCvt() comp.Form {
	return u.DualEditor(yamlCfg, tomlCfg, "Yaml-Toml",
		func(input any) (output any, err error) {
			return convert(input, util.Yaml2Toml)
		}, func(input any) (output any, err error) {
			return convert(input, util.Toml2Yaml)
		})
}

func convert(input any, cvt func([]byte) (*bytes.Buffer, error)) (string, error) {
	buf, err := cvt([]byte(input.(string)))
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
