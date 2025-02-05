package comp

import (
	"github.com/zrcoder/amisgo"
)

func DualEditor(app *amisgo.App, left, right EditorCfg, title string, action, reverseAction func(input any) (output any, err error)) any {
	left.Name = "input"
	right.Name = "output"
	left.ReadOnly = true
	right.ReadOnly = true
	right.Value = ""
	actions := make([]any, 0, 2)
	if action != nil {
		left.ReadOnly = false
		actions = append(actions,
			app.Action().Label("▶").Transform(action, left.Name, right.Name))
	}
	if reverseAction != nil {
		right.ReadOnly = false
		actions = append(actions,
			app.Action().Label("◀︎").Transform(reverseAction, right.Name, left.Name))
	}
	return app.Form().Title(title).ColumnCount(3).AutoFocus(true).WrapWithPanel(false).Body(
		Editor(app, left),
		app.ButtonGroup().Vertical(true).Buttons(
			actions...,
		),
		Editor(app, right),
	).Actions()
}
