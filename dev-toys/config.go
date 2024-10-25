package main

import (
	"embed"

	"dtoy/comp"

	"github.com/zrcoder/amisgo"
)

//go:embed assets/samples/diy_chart.json
var sampleChartCfg string

//go:embed assets/samples/sample.json
var sampleJson string

//go:embed assets/samples/sample.yaml
var sampleYaml string

//go:embed assets/samples/sample.toml
var sampleToml string

//go:embed assets/*
var assetsFS embed.FS

var appConfig *amisgo.Config

func init() {
	appConfig = amisgo.GetDefaultConfig()
	appConfig.Theme = amisgo.ThemeDark
	appConfig.Lang = amisgo.LangEn
	appConfig.StaticDir = "assets"
	appConfig.StaticFS = assetsFS
	appConfig.Icon = "/assets/favicon.ico"
}

var (
	jsonCfg = comp.EditorCfg{Lang: "json", Value: sampleJson}
	yamlCfg = comp.EditorCfg{Lang: "yaml", Value: sampleYaml}
	tomlCfg = comp.EditorCfg{Lang: "toml", Value: sampleToml}
	htmlCfg = comp.EditorCfg{Lang: "html"}
)
