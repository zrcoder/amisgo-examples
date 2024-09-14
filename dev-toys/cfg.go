package main

import (
	"embed"

	"amisgo-examples/dev-toys/comp"

	"github.com/zrcoder/amisgo"
)

//go:embed assets/*
var assetsFS embed.FS

var appConfig = amisgo.GetDefaultConfig()

func init() {
	appConfig.Theme = amisgo.ThemeDark
	appConfig.Lang = amisgo.LangEn
	appConfig.StaticDir = "assets"
	appConfig.StaticFS = assetsFS
	appConfig.Icon = "/assets/favicon.ico"
}

const (
	sampleJson = `{
  "name": "Tom",
  "age": 27,
  "address": {
    "country": "US",
    "code": "7000000"
  }
}`
	sampleYaml = `address:
    code: "7000000"
    country: US
age: 27
name: Tom`
	sampleToml = `age = 27.0
name = "Tom"

[address]
  code = "7000000"
  country = "US"`
)

var (
	jsonCfg = comp.EditorCfg{Lang: "json", Value: sampleJson}
	yamlCfg = comp.EditorCfg{Lang: "yaml", Value: sampleYaml}
	tomlCfg = comp.EditorCfg{Lang: "toml", Value: sampleToml}
	htmlCfg = comp.EditorCfg{Lang: "html"}
)
