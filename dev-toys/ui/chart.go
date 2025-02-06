package ui

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

type Chart struct {
	*chart.Chart
}

func NewChart(app *amisgo.App) *Chart { return &Chart{chart.New(app)} }
func (c *Chart) Line() any            { return c.GenLine(lineXAxis, lineValues) }
func (c *Chart) Bar() any             { return c.GenBar(barXAxis, barValues) }
func (c *Chart) Polar() any           { return c.GenPolar(polarValues1, polarValues2) }
func (c *Chart) Pie() any             { return c.GenPie(pieRadarData) }
func (c *Chart) Scatter() any         { return c.GenScatter(scatterXValues, scatterYValues) }
func (c *Chart) Radar() any           { return c.GenRadar(pieRadarData) }
func (c *Chart) Diy() any             { return c.GenCommon(sampleChartCfg) }
