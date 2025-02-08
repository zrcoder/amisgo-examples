package comp

import (
	"github.com/zrcoder/amisgo"
)

func EditorImg(app *amisgo.App, lang, value string, transfor func(any) (any, error)) any {
	return app.Form().Title("").AutoFocus(true).ColumnCount(3).WrapWithPanel(false).Body(
		app.Wrapper().ClassName("w-1/2").Body(
			Editor(app, EditorCfg{Lang: lang, Name: "editor", Value: value}),
		),
		app.ButtonGroup().Vertical(true).Buttons(
			app.Action().Label("▶").Transform(transfor, "editor", "img"),
		),
		app.Wrapper().ClassName("w-2/5").Body(
			app.Flex().ClassName("h-full").AlignItems("center").Items(
				app.Image().Width("80%").Name("img").ImageMode("original").InnerClassName("no-border"),
			),
		),
	).Actions()
}
