package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	page := comp.Page().Body(
		comp.ButtonToolbar().Buttons(
			comp.Button().Label("Dialog").ActionType("dialog").Dialog(comp.Dialog().Title("Test")),
			comp.Button().Label("Drawer").ActionType("drawer").Drawer(comp.Drawer().Title("Test")),
			comp.Button().Label("Toast").ActionType("toast").Toast(comp.Toast().Items(comp.Schema{"body": "Test"})),
		),
	)
	panic(amisgo.ListenAndServe(page))
}
