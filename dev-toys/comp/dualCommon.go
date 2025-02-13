package comp

import (
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/model"
)

func (c *Comp) EditorImg(lang, value string, transfor func(any) (any, error)) comp.Form {
	return c.DualForm(
		c.Editor(EditorCfg{Lang: lang, Name: "editor", Value: value}),
		c.Image().Width("80%").Name("img").ImageMode("original").InnerClassName("no-border"),
		true,
		c.Action().Label("▶").Transform(transfor, "editor", "img"),
	)
}

func (c *Comp) EditorJson(value string) comp.Form {
	return c.DualForm(
		c.Editor(EditorCfg{Lang: "json", Name: "editor", Value: value}),
		c.Json().Source("${DECODEJSON(editor)}"),
		true,
	)
}

func (c *Comp) EditorQrCoder() comp.Form {
	return c.DualForm(
		c.Editor(EditorCfg{Lang: "text", Name: "editor"}),
		c.Wrapper().Body(
			c.QRCode().ID("qrcode").Value("${editor}").CodeSize(256).Level("M").BackgroundColor("white").ForegroundColor("#333"),
			c.Action().ClassName("w-full").Icon("fa fa-download").Label("Download").VisibleOn("${editor !== ''}").OnEvent(
				model.Schema{
					"click": model.Schema{
						"actions": []model.EventAction{
							c.EventAction().ActionType("saveAs").ComponentID("qrcode").Args(model.Schema{"name": "download.png"}),
						},
					},
				},
			),
		),
		true,
	)
}

func (c *Comp) EditorChart(commCfg string, getData func() (any, error), submit func(model.Schema) error) comp.Form {
	return c.Form().WrapWithPanel(false).ColumnCount(3).AutoFocus(true).Body(
		c.DualFormBody(
			c.Editor(EditorCfg{Lang: "json", Name: "in", Value: commCfg}),
			c.Chart().Name("diy-out").GetData(getData),
			true,
			c.SubmitAction().Label("▶︎").Reload("diy-out"),
		)...,
	).Submit(submit)
}

func (c *Comp) QrcodeEditor(action func([]byte) (path string, err error), getData func() (string, error)) comp.Form {
	return c.DualForm(
		c.InputImage().Name("img").Upload(int64(10*(1<<20)), action),
		c.Service().Name("out").GetData(func() (any, error) {
			res, err := getData()
			return model.Schema{"decqr": res}, err
		}).Body(
			c.Editor(EditorCfg{Name: "text", ReadOnly: true, Value: "${decqr}"}),
		),
		false,
		c.SubmitAction().Label("▶︎").Reload("out"),
	)
}

func (c *Comp) DualForm(left, right any, leftMain bool, buttons ...any) comp.Form {
	return c.Form().AutoFocus(true).ColumnCount(3).WrapWithPanel(false).Body(
		c.DualFormBody(left, right, leftMain, buttons...)...,
	).Actions()
}

func (c *Comp) DualFormBody(left, right any, leftMain bool, buttons ...any) []any {
	if leftMain {
		right = c.Flex().ClassName("h-full").AlignItems("center").Items(right)
	} else {
		left = c.Flex().ClassName("h-full").AlignItems("center").Items(left)
	}
	if len(buttons) == 0 {
		return []any{
			c.Wrapper().ClassName("w-1/2").Body(left),
			c.Wrapper().ClassName("w-1/2").Body(right),
		}
	}
	leftClass := "w-1/2"
	rightClass := "w-2/5"
	if !leftMain {
		leftClass, rightClass = rightClass, leftClass
	}
	return []any{
		c.Wrapper().ClassName(leftClass).Body(left),
		c.ButtonGroup().Vertical(true).Buttons(buttons...),
		c.Wrapper().ClassName(rightClass).Body(right),
	}
}
