package comp

import (
	"github.com/zrcoder/amisgo"
)

func EditorJson(app *amisgo.App, value string) any {
	return app.Form().Title("").AutoFocus(true).ColumnCount(2).WrapWithPanel(false).Body(
		app.Wrapper().ClassName("w-1/2").Body(
			Editor(app, EditorCfg{Lang: "json", Name: "editor", Value: value}),
		),
		app.Wrapper().ClassName("w-1/2").Body(
			app.Json().Source("${DECODEJSON(editor)}"),
		),
	).Actions()
}
