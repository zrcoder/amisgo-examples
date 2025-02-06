package comp

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/model"
)

func EditorJson(app *amisgo.App, value string) any {
	return app.Form().Title("").AutoFocus(true).ColumnCount(2).WrapWithPanel(false).Body(
		app.Wrapper().Style(model.Schema{"width": "50%"}).Body(
			Editor(app, EditorCfg{Lang: "json", Name: "editor", Value: value}),
		),
		app.Wrapper().Style(model.Schema{"width": "50%"}).Body(
			app.Json().Source("${DECODEJSON(editor)}"),
		),
	).Actions()
}
