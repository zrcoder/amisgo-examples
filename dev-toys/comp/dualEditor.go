package comp

import "github.com/zrcoder/amisgo/comp"

func DualEditor(left, right EditorCfg, title string, action, reverseAction func(input any) (output any, err error)) any {
	left.Name = "input"
	right.Name = "output"
	left.ReadOnly = true
	right.ReadOnly = true
	right.Value = ""
	actions := make([]any, 0, 2)
	if action != nil {
		left.ReadOnly = false
		actions = append(actions,
			comp.Action().Label("▶").Transform(action, left.Name, right.Name))
	}
	if reverseAction != nil {
		right.ReadOnly = false
		actions = append(actions,
			comp.Action().Label("◀︎").Transform(reverseAction, right.Name, left.Name))
	}
	return comp.Form().Title(title).ColumnCount(3).AutoFocus(true).WrapWithPanel(false).Body(
		Editor(left),
		comp.ButtonGroup().Vertical(true).Buttons(
			actions...,
		),
		Editor(right),
	).Actions()
}
