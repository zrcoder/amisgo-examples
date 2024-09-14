package main

import (
	_ "embed"

	"amisgo-examples/dev-toys/comp/chart"
)

const (
	lineXAxis      = `Jan, Feb, Mar, Apr, May`
	lineValues     = `5, -1, 7, 8, -3`
	barXAxis       = `A, B, C, D, E`
	barValues      = `40, 20, 30, 50, 10`
	polarValues1   = "0, 0.01, 0.02, 0.06, 0.1, 0.2"
	polarValues2   = "0, 1, 2, 3, 4, 5"
	scatterXValues = "-30, 50, 100, 150"
	scatterYValues = "50, 70, -20, 80"
)

var pieRadarData = map[string]any{
	"A": 100.0,
	"B": 88.0,
	"C": 96.0,
	"D": 72.0,
}

//go:embed data/diy_chart.json
var commCfg string

var (
	lineChart    = chart.GenLine(lineXAxis, lineValues)
	barChart     = chart.GenBar(barXAxis, barValues)
	polarChart   = chart.GenPolar(polarValues1, polarValues2)
	pieChart     = chart.GenPie(pieRadarData)
	scatterChart = chart.GenScatter(scatterXValues, scatterYValues)
	radarChart   = chart.GenRadar(pieRadarData)
	diyChart     = chart.GenCommon(commCfg)
)
