package comp

import "github.com/zrcoder/amisgo/comp"

func (c *Comp) DualEditor(left, right EditorCfg, title string, action, reverseAction func(input any) (output any, err error)) comp.Form {
	left.Name = "input"
	right.Name = "output"
	left.ReadOnly = true
	right.ReadOnly = true
	right.Value = ""
	actions := make([]comp.Action, 0, 2)
	if action != nil {
		left.ReadOnly = false
		actions = append(actions,
			c.Action().Label("▶").Transform(action, left.Name, right.Name))
	}
	if reverseAction != nil {
		right.ReadOnly = false
		actions = append(actions,
			c.Action().Label("◀︎").Transform(reverseAction, right.Name, left.Name))
	}
	return c.Form().Title(title).ColumnCount(3).AutoFocus(true).WrapWithPanel(false).Body(
		c.Editor(left),
		c.ButtonGroup().Vertical(true).Buttons(
			actions...,
		),
		c.Editor(right),
	).Actions()
}
