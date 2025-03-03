package hbox

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
)

func Demos(a *amisgo.App) []app.Demo {
	return []app.Demo{
		{Name: "Hbox", View: a.Page().Body(
			a.HBox().ClassName("b-a bg-dark lter").Columns(
				a.Column(a.Plain().Text("Col A")).ColumnClassName("wrapper-xs b-r").Valign("cc"),
				a.Plain().Text("Col B"),
			),
			a.HBox().ClassName("b-a m-t bg-dark lter").Columns(
				a.Column(a.Plain().Text("Col A")).ColumnClassName("w-md wrapper-xs bg-primary b-r"),
				a.Plain().Text("..."),
			),
		)},
	}
}
