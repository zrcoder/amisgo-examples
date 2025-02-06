package ui

import (
	_ "embed"

	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
)

//go:embed samples/diy_chart.json
var sampleChartCfg string

//go:embed samples/sample.json
var sampleJson string

//go:embed samples/sample.yaml
var sampleYaml string

//go:embed samples/sample.toml
var sampleToml string

var (
	jsonCfg = comp.EditorCfg{Lang: "json", Value: sampleJson}
	yamlCfg = comp.EditorCfg{Lang: "yaml", Value: sampleYaml}
	tomlCfg = comp.EditorCfg{Lang: "toml", Value: sampleToml}
	htmlCfg = comp.EditorCfg{Lang: "html"}
)
