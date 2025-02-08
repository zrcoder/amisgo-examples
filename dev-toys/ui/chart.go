package ui

import "github.com/zrcoder/amisgo/comp"

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

func (u *UI) Line() comp.Wrapper    { return u.chart.GenLine(lineXAxis, lineValues) }
func (u *UI) Bar() comp.Wrapper     { return u.chart.GenBar(barXAxis, barValues) }
func (u *UI) Polar() comp.Wrapper   { return u.chart.GenPolar(polarValues1, polarValues2) }
func (u *UI) Pie() comp.Wrapper     { return u.chart.GenPie(pieRadarData) }
func (u *UI) Scatter() comp.Wrapper { return u.chart.GenScatter(scatterXValues, scatterYValues) }
func (u *UI) Radar() comp.Wrapper   { return u.chart.GenRadar(pieRadarData) }
func (u *UI) Diy() comp.Form        { return u.chart.GenCommon(sampleChartCfg) }
