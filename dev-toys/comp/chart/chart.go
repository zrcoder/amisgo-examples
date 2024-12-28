package chart

import (
	_ "embed"
	"encoding/json"
	"strings"
	"sync"

	"github.com/zrcoder/amisgo-examples/dev-toys/comp"

	ac "github.com/zrcoder/amisgo/comp"
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

func GenLine(xAxis, values string) any {
	storeCfg(keyLine, genCfg(xAxis, values, keyLine))
	return gen(xAxis, values, keyLine)
}

func GenBar(xAxis, values string) any {
	storeCfg(keyBar, genCfg(xAxis, values, keyBar))
	return gen(xAxis, values, keyBar)
}

func gen(xAxis, values, cType string) any {
	return ac.Wrapper().Body(
		ac.Chart().Name(cType).GetData(func() (any, error) {
			return loadCfg(cType), nil
		}),
		ac.Form().Mode("horizontal").Horizontal(ac.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			ac.Flex().ClassName("pb-4").Items(
				ac.Button().Icon("fa fa-arrow-up").Reload(cType).ActionType("submit"),
			),
			ac.InputText().Label("XAxis").Name("xAxisData").Value(xAxis),
			ac.InputText().Label("Values").Name("values").Value(values),
		).Submit(func(d ac.Data) error {
			cfg := genCfg(d.Get("xAxisData").(string), d.Get("values").(string), cType)
			storeCfg(cType, cfg)
			return nil
		}),
	)
}

func genCfg(xAxisData, values string, cType string) ac.ChartCfg {
	return ac.ChartConfig().
		Tooltip(ac.Schema{"trigger": "axis"}).
		XAxis(ac.ChartAxis{"type": "category", "data": strings.Split(xAxisData, ",")}).
		YAxis(ac.ChartAxis{"type": "value"}).
		Series(
			ac.ChartSeries().
				Type(cType).
				Data(strings.Split(values, ",")))
}

func GenPolar(input1, input2 string) any {
	storeCfg(keyPolar, genPolarCfg(input1, input2))
	return ac.Wrapper().Body(
		ac.Chart().Name("polar-out").GetData(func() (any, error) {
			return loadCfg(keyPolar), nil
		}),
		ac.Form().Mode("horizontal").Horizontal(ac.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			ac.Flex().ClassName("pb-4").Items(
				ac.Button().Icon("fa fa-arrow-up").Reload("polar-out").ActionType("submit"),
			),
			ac.InputText().Label("values1").Name("xAxisData").Value(input1),
			ac.InputText().Label("Values2").Name("values").Value(input2),
		).Submit(func(d ac.Data) error {
			cfg := genPolarCfg(d.Get("xAxisData").(string), d.Get("values").(string))
			storeCfg(keyPolar, cfg)
			return nil
		}),
	)
}

func genPolarCfg(input1, input2 string) ac.ChartCfg {
	values1, values2 := strings.Split(input1, ","), strings.Split(input2, ",")
	data := make([][2]string, min(len(values1), len(values2)))
	for i := range data {
		data[i] = [2]string{values1[i], values2[i]}
	}
	return ac.ChartConfig().
		Polar(ac.Schema{"center": []string{"50%", "54%"}}).
		Tooltip(ac.Schema{"trigger": "axis", "axisPointer": ac.Schema{"type": "cross"}}).
		AngleAxis(ac.ChartAxis{"type": "value"}).
		RadiusAxis(ac.ChartAxis{"min": 0}).
		AngleAxis(ac.ChartAxis{"type": "value", "startAngle": 0}).
		Series(
			ac.ChartSeries().
				CoordinateSystem(keyPolar).
				Name(keyLine).
				Type(keyLine).
				ShowSymbol(false).
				Data(data),
		)
}

func GenScatter(input1, input2 string) any {
	storeCfg(keyScatter, genScatterCfg(input1, input2))
	return ac.Wrapper().Body(
		ac.Chart().Name("scatter-out").GetData(func() (any, error) {
			return loadCfg(keyScatter), nil
		}),
		ac.Form().Mode("horizontal").Horizontal(ac.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			ac.Flex().ClassName("pb-4").Items(
				ac.Button().Icon("fa fa-arrow-up").Reload("scatter-out").ActionType("submit"),
			),
			ac.InputText().Label("X").Name("x").Value(input1),
			ac.InputText().Label("Y").Name("y").Value(input2),
		).Submit(func(d ac.Data) error {
			cfg := genScatterCfg(d.Get("x").(string), d.Get("y").(string))
			storeCfg(keyScatter, cfg)
			return nil
		}),
	)
}

func genScatterCfg(input1, input2 string) ac.ChartCfg {
	arr1, arr2 := strings.Split(input1, ","), strings.Split(input2, ",")
	data := make([][2]string, min(len(arr1), len(arr2)))
	for i := range data {
		data[i] = [2]string{strings.TrimSpace(arr1[i]), strings.TrimSpace(arr2[i])}
	}
	return ac.ChartConfig().
		XAxis(ac.ChartAxis{"type": "value"}).
		YAxis(ac.ChartAxis{"type": "value"}).
		Tooltip(ac.Schema{"trigger": "item"}).
		Series(
			ac.ChartSeries().
				Type(keyScatter).
				Data(data).
				ItemStyle(ac.Schema{"color": "#4CAF50"}))
}

func GenPie(data map[string]any) any {
	storeCfg(keyPie, genPieCfg(data))
	return ac.Wrapper().Body(
		ac.Chart().Name("pie-out").GetData(func() (any, error) {
			return loadCfg(keyPie), nil
		}),
		ac.Form().Mode("horizontal").Horizontal(ac.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			ac.Flex().ClassName("pb-4").Items(
				ac.Button().Icon("fa fa-arrow-up").Reload("pie-out").ActionType("submit"),
			),
			ac.InputKV().Name("pd").ValueType("input-number").Value(data),
		).Submit(func(d ac.Data) error {
			kvs := d.Get("pd").(map[string]any)
			storeCfg(keyPie, genPieCfg(kvs))
			return nil
		}),
	)
}

func genPieCfg(kvs map[string]any) ac.ChartCfg {
	data := make([]ac.Data, 0, len(kvs))
	for k, v := range kvs {
		data = append(data, ac.Data{"name": k, "value": v})
	}
	return ac.ChartConfig().
		Tooltip(ac.Schema{"trigger": "item"}).
		Series(ac.ChartSeries().Type(keyPie).
			Data(data).
			Label(ac.Schema{"formatter": "{b}:  {d}%", "backgroundColor": "#5971C0", "borderRadius": 10, "padding": 5}))
}

func GenRadar(data map[string]any) any {
	storeCfg(keyRadar, genRadarCfg(data))
	return ac.Wrapper().Body(
		ac.Chart().Name("radar-out").GetData(func() (any, error) {
			return loadCfg(keyRadar), nil
		}),
		ac.Form().Mode("horizontal").Horizontal(ac.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			ac.Flex().ClassName("pb-4").Items(
				ac.Button().Icon("fa fa-arrow-up").Reload("radar-out").ActionType("submit"),
			),
			ac.InputKV().Name("rd").ValueType("input-number").Value(data),
		).Submit(func(d ac.Data) error {
			kvs := d.Get("rd").(map[string]any)
			storeCfg(keyRadar, genRadarCfg(kvs))
			return nil
		}),
	)
}

func genRadarCfg(kvs map[string]any) ac.ChartCfg {
	ind := make([]ac.Data, 0, len(kvs))
	values := make([]float64, 0, len(kvs))
	for k, v := range kvs {
		ind = append(ind, ac.Data{"name": k, "max": 100})
		values = append(values, v.(float64))
	}
	return ac.ChartConfig().
		Tooltip(ac.Schema{"trigger": "item"}).
		Radar(ac.Schema{"indicator": ind}).
		Series(
			ac.ChartSeries().
				Name("Athlete Performance").
				Type(keyRadar).
				Data([]ac.Schema{
					{
						"name":  "Athlete",
						"value": values,
					},
				}).
				AreaStyle(ac.Schema{"color": "rgba(255, 99, 132, 0.2)"}).
				LineStyle(ac.Schema{"color": "#FF6384"}))
}

func GenCommon(commCfg string) any {
	storeCfg(keyCommon, genCommonCfg(commCfg))
	return ac.Form().WrapWithPanel(false).ColumnCount(3).AutoFocus(true).Body(
		ac.Wrapper().Style(ac.Schema{"width": "50%"}).Body(
			comp.Editor(comp.EditorCfg{Lang: "json", Name: "in", Value: commCfg}),
		),
		ac.ButtonGroup().Vertical(true).Buttons(
			ac.Button().Icon("fa fa-arrow-right").Reload("diy-out").ActionType("submit"),
		),
		ac.Flex().Style(ac.Schema{"width": "40%"}).AlignItems("center").Items(
			ac.Chart().Name("diy-out").GetData(func() (any, error) {
				return loadCfg(keyCommon), nil
			}),
		),
	).Submit(func(d ac.Data) error {
		data := []byte(d.Get("in").(string))
		var cfg ac.ChartCfg
		err := json.Unmarshal(data, &cfg)
		if err != nil {
			return err
		}
		storeCfg(keyCommon, cfg)
		return nil
	})
}

func genCommonCfg(input string) ac.ChartCfg {
	var cfg ac.ChartCfg
	json.Unmarshal([]byte(input), &cfg)
	return cfg
}

var ConfigCache sync.Map

func loadCfg(key string) ac.ChartCfg {
	res, _ := ConfigCache.Load(key)
	return res.(ac.ChartCfg)
}

func storeCfg(key string, val ac.ChartCfg) {
	ConfigCache.Store(key, val)
}
