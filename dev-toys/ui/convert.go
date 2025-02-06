package ui

import (
	"bytes"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
	"github.com/zrcoder/amisgo-examples/dev-toys/util"
)

type Converters struct {
	*amisgo.App
}

func NewConverters(app *amisgo.App) *Converters { return &Converters{App: app} }

func (c *Converters) JsonYamlCvt() any {
	return comp.DualEditor(c.App, jsonCfg, yamlCfg, "Json-Yaml",
		func(input any) (output any, err error) {
			return c.convert(input, util.Json2Yaml)
		}, func(input any) (output any, err error) {
			return c.convert(input, util.Yaml2Json)
		})
}

func (c *Converters) JsonTomlCvt() any {
	return comp.DualEditor(c.App, jsonCfg, tomlCfg, "Json-Toml",
		func(input any) (output any, err error) {
			return c.convert(input, util.Json2Toml)
		}, func(input any) (output any, err error) {
			return c.convert(input, util.Toml2Json)
		})
}

func (c *Converters) YamlTomlCvt() any {
	return comp.DualEditor(c.App, yamlCfg, tomlCfg, "Yaml-Toml",
		func(input any) (output any, err error) {
			return c.convert(input, util.Yaml2Toml)
		}, func(input any) (output any, err error) {
			return c.convert(input, util.Toml2Yaml)
		})
}

func (c *Converters) convert(input any, cvt func([]byte) (*bytes.Buffer, error)) (string, error) {
	buf, err := cvt([]byte(input.(string)))
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
