package pages

import (
	"bytes"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
	"github.com/zrcoder/amisgo-examples/dev-toys/util"
)

func JsonYamlCvt(app *amisgo.App) any {
	return comp.DualEditor(app, jsonCfg, yamlCfg, "Json-Yaml",
		func(input any) (output any, err error) {
			return convert(input, util.Json2Yaml)
		}, func(input any) (output any, err error) {
			return convert(input, util.Yaml2Json)
		})
}

func JsonTomlCvt(app *amisgo.App) any {
	return comp.DualEditor(app, jsonCfg, tomlCfg, "Json-Toml",
		func(input any) (output any, err error) {
			return convert(input, util.Json2Toml)
		}, func(input any) (output any, err error) {
			return convert(input, util.Toml2Json)
		})
}

func YamlTomlCvt(app *amisgo.App) any {
	return comp.DualEditor(app, yamlCfg, tomlCfg, "Yaml-Toml",
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
