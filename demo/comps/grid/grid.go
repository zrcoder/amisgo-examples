package grid

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
)

func Demos(a *amisgo.App) []app.Demo {
	return []app.Demo{
		{Name: "Base", View: a.Wrapper().Body(
			a.Grid().Columns(
				a.Column().Body(a.Plain().Text("Colunm 1")).ColumnClassName("bg-green-300"),
				a.Column().Body(a.Plain().Text("Colunm 2")).ColumnClassName("bg-blue-300"),
			),
			a.Grid().ClassName("m-t").Columns(
				a.Column().Body(a.Plain().Text("Colunm 1")).ColumnClassName("bg-green-300"),
				a.Column().Body(a.Plain().Text("Colunm 2")).ColumnClassName("bg-blue-300"),
				a.Column().Body(a.Plain().Text("Colunm 3")).ColumnClassName("bg-red-300"),
			),
		)},
		{Name: "md", View: a.Grid().Columns(
			a.Column().Body(a.Plain().Text("Colunm 1")).ColumnClassName("bg-green-300"),
			a.Column().Body(a.Plain().Text("Colunm 2")).ColumnClassName("bg-blue-300").Md("9"),
		)},
	}
}
