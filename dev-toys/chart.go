package main

import (
	"encoding/json"
	"sync"

	"amisgo-examples/dev-toys/comp"

	ac "github.com/zrcoder/amisgo/comp"
)

var (
	chartLineCfg = ac.ChartConfig().
			Title(ac.Schema{"text": "Line Chart"}).
			Tooltip(ac.Schema{"trigger": "axis"}).
			XAxis(ac.ChartAxis{"type": "category", "data": []string{"Jan", "Feb", "Mar", "Apr", "May"}}).
			YAxis(ac.ChartAxis{"type": "value"}).
			Series(
			ac.ChartSeries().
				Type("line").
				Data([]int{5, -1, 7, 8, -3}).
				Smooth(true).
				LineStyle(ac.Schema{"color": "#42A5F5"}))
	chartBarCfg = ac.ChartConfig().
			XAxis(ac.ChartAxis{"type": "category", "data": []string{"A", "B", "C", "D", "E"}}).
			YAxis(ac.ChartAxis{"type": "value"}).
			Series(
			ac.ChartSeries().
				Name("Category Values").
				Type("bar").
				Data([]int{40, 20, 30, 50, 10}))
	chartPolarCfg = ac.ChartConfig().
			Polar(ac.Schema{"center": []string{"50%", "54%"}}).
			Tooltip(ac.Schema{"trigger": "axis", "axisPointer": ac.Schema{"type": "cross"}}).
			AngleAxis(ac.ChartAxis{"type": "value"}).
			RadiusAxis(ac.ChartAxis{"min": 0}).
			AngleAxis(ac.ChartAxis{"type": "value", "startAngle": 0}).
			Series(
			ac.ChartSeries().
				CoordinateSystem("polar").
				Name("line").
				Type("line").
				ShowSymbol(false).
				Data([][2]float64{
					{0, 0},
					{0.03487823687206265, 1},
					{0.06958655048003272, 2},
					{0.10395584540887964, 3},
					{0.13781867790849958, 4},
				}),
		).AnimationDuration(2000)
	chartPieCfg = ac.ChartConfig().
			Series(
			ac.ChartSeries().
				Name("Pie Chart").
				Type("pie").
				Radius("50%").
				Data([]ac.Schema{
					{"value": 300, "name": "Red"},
					{"value": 50, "name": "Blue"},
					{"value": 100, "name": "Yellow"},
				}))
	chartScatterCfg = ac.ChartConfig().
			XAxis(ac.ChartAxis{"type": "value"}).
			YAxis(ac.ChartAxis{"type": "value"}).
			Series(
			ac.ChartSeries().
				Name("Scatter Dataset").
				Type("scatter").
				Data([][2]int{
					{-30, 50},
					{50, 70},
					{100, -20},
					{150, 80},
				}).
				ItemStyle(ac.Schema{"color": "#4CAF50"}))
	chartRadarCfg = ac.ChartConfig().
			Radar(
			ac.Schema{
				"indicator": []ac.Schema{
					{"name": "Speed", "max": 100},
					{"name": "Strength", "max": 100},
					{"name": "Endurance", "max": 100},
					{"name": "Agility", "max": 100},
				},
			}).
		Series(
			ac.ChartSeries().
				Name("Athlete Performance").
				Type("radar").
				Data([]ac.Schema{
					{
						"name":  "Athlete",
						"value": []int{60, 70, 80, 90},
					},
				}).
				AreaStyle(ac.Schema{"color": "rgba(255, 99, 132, 0.2)"}).
				LineStyle(ac.Schema{"color": "#FF6384"}))

	chartCfgMutex sync.Mutex
	chartCfg      = chartLineCfg
)

func getChartCfg() any {
	chartCfgMutex.Lock()
	defer chartCfgMutex.Unlock()
	return chartCfg
}

func setChartCfg(value ac.ChartCfg) {
	chartCfgMutex.Lock()
	defer chartCfgMutex.Unlock()
	chartCfg = value
}

var chart = ac.Form().WrapWithPanel(false).ColumnCount(3).Body(
	ac.Wrapper().Style(ac.Schema{"width": "50%"}).Body(
		ac.Select().Name("sample").Label("Examples").Mode("inline").Options(
			ac.Option().Label("line").Value(chartLineCfg.JsonStr()),
			ac.Option().Label("bar").Value(chartBarCfg.JsonStr()),
			ac.Option().Label("polar").Value(chartPolarCfg.JsonStr()),
			ac.Option().Label("pie").Value(chartPieCfg.JsonStr()),
			ac.Option().Label("scatter").Value(chartScatterCfg.JsonStr()),
			ac.Option().Label("radar").Value(chartRadarCfg.JsonStr()),
		).Value(chartLineCfg.JsonStr()),
		comp.Editor(comp.EditorCfg{Lang: "json", Name: "in", Value: "${sample}"}),
	),
	ac.ButtonGroup().Vertical(true).Buttons(
		ac.Button().Icon("fa fa-arrow-right").Reload("out").ActionType("submit"),
	),
	ac.Flex().Style(ac.Schema{"width": "40%"}).AlignItems("center").Items(
		ac.Chart().Name("out").GetData(getChartCfg),
	),
).Go(func(m map[string]any) {
	data := []byte(m["in"].(string))
	var cfg ac.ChartCfg
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	setChartCfg(cfg)
})
