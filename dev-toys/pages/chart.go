package pages

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/dev-toys/comp/chart"
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

func LineChart(app *amisgo.App) any  { return chart.GenLine(app, lineXAxis, lineValues) }
func BarChart(app *amisgo.App) any   { return chart.GenBar(app, barXAxis, barValues) }
func PolarChart(app *amisgo.App) any { return chart.GenPolar(app, polarValues1, polarValues2) }
func PieChart(app *amisgo.App) any   { return chart.GenPie(app, pieRadarData) }
func ScatterChart(app *amisgo.App) any {
	return chart.GenScatter(app, scatterXValues, scatterYValues)
}
func RadarChart(app *amisgo.App) any { return chart.GenRadar(app, pieRadarData) }
func DiyChart(app *amisgo.App) any   { return chart.GenCommon(app, sampleChartCfg) }
