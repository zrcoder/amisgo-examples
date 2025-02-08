package chart

import (
	_ "embed"
	"encoding/json"
	"strings"
	"sync"

	"github.com/zrcoder/amisgo-examples/dev-toys/comp"

	ac "github.com/zrcoder/amisgo/comp"
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

type Chart struct {
	*comp.Comp
	ConfigCache *sync.Map
}

func New(c *comp.Comp) *Chart {
	return &Chart{Comp: c, ConfigCache: &sync.Map{}}
}

func (c *Chart) GenLine(xAxis, values string) ac.Wrapper {
	c.storeCfg(keyLine, c.genCfg(xAxis, values, keyLine))
	return c.gen(xAxis, values, keyLine)
}

func (c *Chart) GenBar(xAxis, values string) ac.Wrapper {
	c.storeCfg(keyBar, c.genCfg(xAxis, values, keyBar))
	return c.gen(xAxis, values, keyBar)
}

func (c *Chart) gen(xAxis, values, cType string) ac.Wrapper {
	return c.Wrapper().Body(
		c.Chart().Name(cType).GetData(func() (any, error) {
			return c.loadCfg(cType), nil
		}),
		c.Form().Mode("horizontal").Horizontal(am.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			c.Flex().ClassName("pb-4").Items(
				c.Button().Label("▲").Reload(cType).ActionType("submit"),
			),
			c.InputText().Label("XAxis").Name("xAxisData").Value(xAxis),
			c.InputText().Label("Values").Name("values").Value(values),
		).Submit(func(d am.Schema) error {
			cfg := c.genCfg(d.Get("xAxisData").(string), d.Get("values").(string), cType)
			c.storeCfg(cType, cfg)
			return nil
		}),
	)
}

func (c *Chart) genCfg(xAxisData, values string, cType string) am.ChartCfg {
	return c.ChartConfig().
		Tooltip(am.Schema{"trigger": "axis"}).
		XAxis(c.ChartAxis().Type("category").Data(strings.Split(xAxisData, ","))).
		YAxis(c.ChartAxis().Type("value")).
		Series(
			c.ChartSeries().
				Type(cType).
				Data(strings.Split(values, ",")))
}

func (c *Chart) GenPolar(input1, input2 string) ac.Wrapper {
	c.storeCfg(keyPolar, c.genPolarCfg(input1, input2))
	return c.Wrapper().Body(
		c.Chart().Name("polar-out").GetData(func() (any, error) {
			return c.loadCfg(keyPolar), nil
		}),
		c.Form().Mode("horizontal").Horizontal(am.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			c.Flex().ClassName("pb-4").Items(
				c.Button().Label("▲").Reload("polar-out").ActionType("submit"),
			),
			c.InputText().Label("values1").Name("xAxisData").Value(input1),
			c.InputText().Label("Values2").Name("values").Value(input2),
		).Submit(func(d am.Schema) error {
			cfg := c.genPolarCfg(d.Get("xAxisData").(string), d.Get("values").(string))
			c.storeCfg(keyPolar, cfg)
			return nil
		}),
	)
}

func (c *Chart) genPolarCfg(input1, input2 string) am.ChartCfg {
	values1, values2 := strings.Split(input1, ","), strings.Split(input2, ",")
	data := make([][2]string, min(len(values1), len(values2)))
	for i := range data {
		data[i] = [2]string{values1[i], values2[i]}
	}
	return c.ChartConfig().
		Polar(am.Schema{"center": []string{"50%", "54%"}}).
		Tooltip(am.Schema{"trigger": "axis", "axisPointer": am.Schema{"type": "cross"}}).
		AngleAxis(c.ChartAxis().Type("value")).
		RadiusAxis(c.ChartAxis().Min(0)).
		AngleAxis(c.ChartAxis().Type("value").StartAngle(0)).
		Series(
			c.ChartSeries().
				CoordinateSystem(keyPolar).
				Name(keyLine).
				Type(keyLine).
				ShowSymbol(false).
				Data(data),
		)
}

func (c *Chart) GenScatter(input1, input2 string) ac.Wrapper {
	c.storeCfg(keyScatter, c.genScatterCfg(input1, input2))
	return c.Wrapper().Body(
		c.Chart().Name("scatter-out").GetData(func() (any, error) {
			return c.loadCfg(keyScatter), nil
		}),
		c.Form().Mode("horizontal").Horizontal(am.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			c.Flex().ClassName("pb-4").Items(
				c.Button().Label("▲").Reload("scatter-out").ActionType("submit"),
			),
			c.InputText().Label("X").Name("x").Value(input1),
			c.InputText().Label("Y").Name("y").Value(input2),
		).Submit(func(d am.Schema) error {
			cfg := c.genScatterCfg(d.Get("x").(string), d.Get("y").(string))
			c.storeCfg(keyScatter, cfg)
			return nil
		}),
	)
}

func (c *Chart) genScatterCfg(input1, input2 string) am.ChartCfg {
	arr1, arr2 := strings.Split(input1, ","), strings.Split(input2, ",")
	data := make([][2]string, min(len(arr1), len(arr2)))
	for i := range data {
		data[i] = [2]string{strings.TrimSpace(arr1[i]), strings.TrimSpace(arr2[i])}
	}
	return c.ChartConfig().
		XAxis(c.ChartAxis().Type("value")).
		YAxis(c.ChartAxis().Type("value")).
		Tooltip(am.Schema{"trigger": "item"}).
		Series(
			c.ChartSeries().
				Type(keyScatter).
				Data(data).
				ItemStyle(am.Schema{"color": "#4CAF50"}))
}

func (c *Chart) GenPie(data map[string]any) ac.Wrapper {
	c.storeCfg(keyPie, c.genPieCfg(data))
	return c.Wrapper().Body(
		c.Chart().Name("pie-out").GetData(func() (any, error) {
			return c.loadCfg(keyPie), nil
		}),
		c.Form().Mode("horizontal").Horizontal(am.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			c.Flex().ClassName("pb-4").Items(
				c.Button().Label("▲").Reload("pie-out").ActionType("submit"),
			),
			c.InputKV().Name("pd").ValueType("input-number").Value(data),
		).Submit(func(d am.Schema) error {
			kvs := d.Get("pd").(map[string]any)
			c.storeCfg(keyPie, c.genPieCfg(kvs))
			return nil
		}),
	)
}

func (c *Chart) genPieCfg(kvs map[string]any) am.ChartCfg {
	data := make([]am.Schema, 0, len(kvs))
	for k, v := range kvs {
		data = append(data, am.Schema{"name": k, "value": v})
	}
	return c.ChartConfig().
		Tooltip(am.Schema{"trigger": "item"}).
		Series(c.ChartSeries().Type(keyPie).
			Data(data).
			Label(am.Schema{"formatter": "{b}:  {d}%", "backgroundColor": "#5971C0", "borderRadius": 10, "padding": 5}))
}

func (c *Chart) GenRadar(data map[string]any) ac.Wrapper {
	c.storeCfg(keyRadar, c.genRadarCfg(data))
	return c.Wrapper().Body(
		c.Chart().Name("radar-out").GetData(func() (any, error) {
			return c.loadCfg(keyRadar), nil
		}),
		c.Form().Mode("horizontal").Horizontal(am.Schema{"justify": true}).WrapWithPanel(false).Actions().Body(
			c.Flex().ClassName("pb-4").Items(
				c.Button().Label("▲").Reload("radar-out").ActionType("submit"),
			),
			c.InputKV().Name("rd").ValueType("input-number").Value(data),
		).Submit(func(d am.Schema) error {
			kvs := d.Get("rd").(map[string]any)
			c.storeCfg(keyRadar, c.genRadarCfg(kvs))
			return nil
		}),
	)
}

func (c *Chart) genRadarCfg(kvs map[string]any) am.ChartCfg {
	ind := make([]am.Schema, 0, len(kvs))
	values := make([]float64, 0, len(kvs))
	for k, v := range kvs {
		ind = append(ind, am.Schema{"name": k, "max": 100})
		values = append(values, v.(float64))
	}
	return c.ChartConfig().
		Tooltip(am.Schema{"trigger": "item"}).
		Radar(am.Schema{"indicator": ind}).
		Series(
			c.ChartSeries().
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

func (c *Chart) GenCommon(commCfg string) ac.Form {
	c.storeCfg(keyCommon, c.genCommonCfg(commCfg))
	return c.EditorChart(
		commCfg,
		func() (any, error) {
			return c.loadCfg(keyCommon), nil
		},
		func(d am.Schema) error {
			data := []byte(d.Get("in").(string))
			var cfg am.ChartCfg
			err := json.Unmarshal(data, &cfg)
			if err != nil {
				return err
			}
			c.storeCfg(keyCommon, cfg)
			return nil
		},
	)
}

func (c *Chart) genCommonCfg(input string) am.ChartCfg {
	var cfg am.ChartCfg
	json.Unmarshal([]byte(input), &cfg)
	return cfg
}

func (c *Chart) loadCfg(key string) am.ChartCfg {
	res, _ := c.ConfigCache.Load(key)
	return res.(am.ChartCfg)
}

func (c *Chart) storeCfg(key string, val am.ChartCfg) {
	c.ConfigCache.Store(key, val)
}
