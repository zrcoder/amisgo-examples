package chart

import (
	_ "embed"
	"encoding/json"
	"strings"
	"sync"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/dev-toys/comp"

	am "github.com/zrcoder/amisgo/model"
)

const (
	keyLine    = "line"
	keyBar     = "bar"
	keyPolar   = "polar"
	keyPie     = "pie"
	keyScatter = "scatter"
	keyRadar   = "radar"
	keyCommon  = "common"
)

func GenLine(app *amisgo.App, xAxis, values string) any {
	storeCfg(keyLine, genCfg(app, xAxis, values, keyLine))
	return gen(app, xAxis, values, keyLine)
}

func GenBar(app *amisgo.App, xAxis, values string) any {
	storeCfg(keyBar, genCfg(app, xAxis, values, keyBar))
	return gen(app, xAxis, values, keyBar)
}

func gen(app *amisgo.App, xAxis, values, cType string) any {
	return app.Wrapper().Body(
		app.Chart().Name(cType).GetData(func() (any, error) {
			return loadCfg(cType), nil
		}),
		app.Form().Mode("horizontal").Horizontal(am.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			app.Flex().ClassName("pb-4").Items(
				app.Button().Label("▲").Reload(cType).ActionType("submit"),
			),
			app.InputText().Label("XAxis").Name("xAxisData").Value(xAxis),
			app.InputText().Label("Values").Name("values").Value(values),
		).Submit(func(d am.Schema) error {
			cfg := genCfg(app, d.Get("xAxisData").(string), d.Get("values").(string), cType)
			storeCfg(cType, cfg)
			return nil
		}),
	)
}

func genCfg(app *amisgo.App, xAxisData, values string, cType string) am.ChartCfg {
	return app.ChartConfig().
		Tooltip(am.Schema{"trigger": "axis"}).
		XAxis(am.ChartAxis{"type": "category", "data": strings.Split(xAxisData, ",")}).
		YAxis(am.ChartAxis{"type": "value"}).
		Series(
			app.ChartSeries().
				Type(cType).
				Data(strings.Split(values, ",")))
}

func GenPolar(app *amisgo.App, input1, input2 string) any {
	storeCfg(keyPolar, genPolarCfg(app, input1, input2))
	return app.Wrapper().Body(
		app.Chart().Name("polar-out").GetData(func() (any, error) {
			return loadCfg(keyPolar), nil
		}),
		app.Form().Mode("horizontal").Horizontal(am.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			app.Flex().ClassName("pb-4").Items(
				app.Button().Label("▲").Reload("polar-out").ActionType("submit"),
			),
			app.InputText().Label("values1").Name("xAxisData").Value(input1),
			app.InputText().Label("Values2").Name("values").Value(input2),
		).Submit(func(d am.Schema) error {
			cfg := genPolarCfg(app, d.Get("xAxisData").(string), d.Get("values").(string))
			storeCfg(keyPolar, cfg)
			return nil
		}),
	)
}

func genPolarCfg(app *amisgo.App, input1, input2 string) am.ChartCfg {
	values1, values2 := strings.Split(input1, ","), strings.Split(input2, ",")
	data := make([][2]string, min(len(values1), len(values2)))
	for i := range data {
		data[i] = [2]string{values1[i], values2[i]}
	}
	return app.ChartConfig().
		Polar(am.Schema{"center": []string{"50%", "54%"}}).
		Tooltip(am.Schema{"trigger": "axis", "axisPointer": am.Schema{"type": "cross"}}).
		AngleAxis(am.ChartAxis{"type": "value"}).
		RadiusAxis(am.ChartAxis{"min": 0}).
		AngleAxis(am.ChartAxis{"type": "value", "startAngle": 0}).
		Series(
			app.ChartSeries().
				CoordinateSystem(keyPolar).
				Name(keyLine).
				Type(keyLine).
				ShowSymbol(false).
				Data(data),
		)
}

func GenScatter(app *amisgo.App, input1, input2 string) any {
	storeCfg(keyScatter, genScatterCfg(app, input1, input2))
	return app.Wrapper().Body(
		app.Chart().Name("scatter-out").GetData(func() (any, error) {
			return loadCfg(keyScatter), nil
		}),
		app.Form().Mode("horizontal").Horizontal(am.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			app.Flex().ClassName("pb-4").Items(
				app.Button().Label("▲").Reload("scatter-out").ActionType("submit"),
			),
			app.InputText().Label("X").Name("x").Value(input1),
			app.InputText().Label("Y").Name("y").Value(input2),
		).Submit(func(d am.Schema) error {
			cfg := genScatterCfg(app, d.Get("x").(string), d.Get("y").(string))
			storeCfg(keyScatter, cfg)
			return nil
		}),
	)
}

func genScatterCfg(app *amisgo.App, input1, input2 string) am.ChartCfg {
	arr1, arr2 := strings.Split(input1, ","), strings.Split(input2, ",")
	data := make([][2]string, min(len(arr1), len(arr2)))
	for i := range data {
		data[i] = [2]string{strings.TrimSpace(arr1[i]), strings.TrimSpace(arr2[i])}
	}
	return app.ChartConfig().
		XAxis(am.ChartAxis{"type": "value"}).
		YAxis(am.ChartAxis{"type": "value"}).
		Tooltip(am.Schema{"trigger": "item"}).
		Series(
			app.ChartSeries().
				Type(keyScatter).
				Data(data).
				ItemStyle(am.Schema{"color": "#4CAF50"}))
}

func GenPie(app *amisgo.App, data map[string]any) any {
	storeCfg(keyPie, genPieCfg(app, data))
	return app.Wrapper().Body(
		app.Chart().Name("pie-out").GetData(func() (any, error) {
			return loadCfg(keyPie), nil
		}),
		app.Form().Mode("horizontal").Horizontal(am.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			app.Flex().ClassName("pb-4").Items(
				app.Button().Label("▲").Reload("pie-out").ActionType("submit"),
			),
			app.InputKV().Name("pd").ValueType("input-number").Value(data),
		).Submit(func(d am.Schema) error {
			kvs := d.Get("pd").(map[string]any)
			storeCfg(keyPie, genPieCfg(app, kvs))
			return nil
		}),
	)
}

func genPieCfg(app *amisgo.App, kvs map[string]any) am.ChartCfg {
	data := make([]am.Schema, 0, len(kvs))
	for k, v := range kvs {
		data = append(data, am.Schema{"name": k, "value": v})
	}
	return app.ChartConfig().
		Tooltip(am.Schema{"trigger": "item"}).
		Series(app.ChartSeries().Type(keyPie).
			Data(data).
			Label(am.Schema{"formatter": "{b}:  {d}%", "backgroundColor": "#5971C0", "borderRadius": 10, "padding": 5}))
}

func GenRadar(app *amisgo.App, data map[string]any) any {
	storeCfg(keyRadar, genRadarCfg(app, data))
	return app.Wrapper().Body(
		app.Chart().Name("radar-out").GetData(func() (any, error) {
			return loadCfg(keyRadar), nil
		}),
		app.Form().Mode("horizontal").Horizontal(am.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			app.Flex().ClassName("pb-4").Items(
				app.Button().Label("▲").Reload("radar-out").ActionType("submit"),
			),
			app.InputKV().Name("rd").ValueType("input-number").Value(data),
		).Submit(func(d am.Schema) error {
			kvs := d.Get("rd").(map[string]any)
			storeCfg(keyRadar, genRadarCfg(app, kvs))
			return nil
		}),
	)
}

func genRadarCfg(app *amisgo.App, kvs map[string]any) am.ChartCfg {
	ind := make([]am.Schema, 0, len(kvs))
	values := make([]float64, 0, len(kvs))
	for k, v := range kvs {
		ind = append(ind, am.Schema{"name": k, "max": 100})
		values = append(values, v.(float64))
	}
	return app.ChartConfig().
		Tooltip(am.Schema{"trigger": "item"}).
		Radar(am.Schema{"indicator": ind}).
		Series(
			app.ChartSeries().
				Name("Athlete Performance").
				Type(keyRadar).
				Data([]am.Schema{
					{
						"name":  "Athlete",
						"value": values,
					},
				}).
				AreaStyle(am.Schema{"color": "rgba(255, 99, 132, 0.2)"}).
				LineStyle(am.Schema{"color": "#FF6384"}))
}

func GenCommon(app *amisgo.App, commCfg string) any {
	storeCfg(keyCommon, genCommonCfg(commCfg))
	return app.Form().WrapWithPanel(false).ColumnCount(3).AutoFocus(true).Body(
		app.Wrapper().Style(am.Schema{"width": "50%"}).Body(
			comp.Editor(app, comp.EditorCfg{Lang: "json", Name: "in", Value: commCfg}),
		),
		app.ButtonGroup().Vertical(true).Buttons(
			app.Button().Label("▶︎").Reload("diy-out").ActionType("submit"),
		),
		app.Flex().Style(am.Schema{"width": "40%"}).AlignItems("center").Items(
			app.Chart().Name("diy-out").GetData(func() (any, error) {
				return loadCfg(keyCommon), nil
			}),
		),
	).Submit(func(d am.Schema) error {
		data := []byte(d.Get("in").(string))
		var cfg am.ChartCfg
		err := json.Unmarshal(data, &cfg)
		if err != nil {
			return err
		}
		storeCfg(keyCommon, cfg)
		return nil
	})
}

func genCommonCfg(input string) am.ChartCfg {
	var cfg am.ChartCfg
	json.Unmarshal([]byte(input), &cfg)
	return cfg
}

var ConfigCache sync.Map

func loadCfg(key string) am.ChartCfg {
	res, _ := ConfigCache.Load(key)
	return res.(am.ChartCfg)
}

func storeCfg(key string, val am.ChartCfg) {
	ConfigCache.Store(key, val)
}
