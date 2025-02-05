package comp

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/model"
)

func EditorImg(app *amisgo.App, lang, value string, transfor func(any) (any, error)) any {
	return app.Form().Title("").AutoFocus(true).ColumnCount(3).WrapWithPanel(false).Body(
		app.Wrapper().Style(model.Schema{"width": "50%"}).Body(
			Editor(app, EditorCfg{Lang: lang, Name: "editor", Value: value}),
		),
		app.ButtonGroup().Vertical(true).Buttons(
			app.Action().Label("▶").Transform(transfor, "editor", "img"),
		),
		app.Flex().Style(model.Schema{"width": "40%"}).AlignItems("center").Items(
			app.Wrapper().Style(model.Schema{"width": "80%"}).Body(
				app.Image().Name("img").ImageMode("original").InnerClassName("no-border"),
			),
		),
	).Actions()
}
