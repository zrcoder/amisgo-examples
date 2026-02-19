package tableview

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

func Demos(a *amisgo.App) []app.Demo {
	return []app.Demo{
		{Name: "Base", View: base(a)},
		{Name: "Cell Style", View: cell(a)},
		{Name: "Col Style", View: col(a)},
		{Name: "Caption", View: caption(a)},
		{Name: "Condition", View: condition(a)},
		{Name: "Visible on", View: visibleOn(a)},
		{Name: "Layout", View: layout(a)},
		{Name: "FormBody", View: hanoi(a)},
	}
}

func base(app *amisgo.App) comp.TableView {
	return app.TableView().Trs(
		app.Tr().Background("#F7F7F7").Tds(
			app.Td().Body(app.Tpl().Tpl("地区")),
			app.Td().Body(app.Tpl().Tpl("城市")),
			app.Td().Body(app.Tpl().Tpl("销量")),
		),
		app.Tr().Tds(
			app.Td().Body(app.Tpl().Tpl("华北")).Rowspan(2),
			app.Td().Body(app.Tpl().Tpl("北京")),
			app.Td().Body(app.Tpl().Tpl("20")),
		),
		app.Tr().Tds(
			app.Td().Body(app.Tpl().Tpl("天津")),
			app.Td().Body(app.Tpl().Tpl("10")),
		),
	)
}

func cell(app *amisgo.App) comp.TableView {
	return app.TableView().Trs(
		app.Tr().Background("#F7F7F7").Tds(
			app.Td().Body(app.Tpl().Tpl("地区")),
			app.Td().Body(app.Tpl().Tpl("城市")),
			app.Td().Body(app.Tpl().Tpl("销量")),
		),
		app.Tr().Tds(
			app.Td().Body(app.Tpl().Tpl("华北")).Style(schema.Schema{"borderBottomWidth": 0, "borderLeftWidth": 0}),
			app.Td().Body(app.Tpl().Tpl("北京")),
			app.Td().Body(app.Tpl().Tpl("20")).Style((schema.Schema{"borderBottomWidth": 0, "borderRightWidth": 0})),
		),
	)
}

func col(app *amisgo.App) comp.TableView {
	return app.TableView().
		Cols(
			app.Tcol().Span(2),
			app.Tcol().Style(schema.Schema{"background": "#00F7F7"}),
		).
		Trs(
			app.Tr().Background("#F7F7F7").Tds(
				app.Td().Body(app.Tpl().Tpl("地区")),
				app.Td().Body(app.Tpl().Tpl("城市")),
				app.Td().Body(app.Tpl().Tpl("销量")),
			),
			app.Tr().Tds(
				app.Td().Body(app.Tpl().Tpl("华北")).Rowspan(2),
				app.Td().Body(app.Tpl().Tpl("北京")),
				app.Td().Body(app.Tpl().Tpl("20")),
			),
			app.Tr().Tds(
				app.Td().Body(app.Tpl().Tpl("天津")),
				app.Td().Body(app.Tpl().Tpl("10")),
			),
		)
}

func caption(app *amisgo.App) comp.TableView {
	return app.TableView().
		Caption("标题").
		// CaptionSide("bottom").
		Cols(
			app.Tcol().Span(2),
			app.Tcol().Style(schema.Schema{"background": "#00F7F7"}),
		).
		Trs(
			app.Tr().Background("#F7F7F7").Tds(
				app.Td().Body(app.Tpl().Tpl("地区")),
				app.Td().Body(app.Tpl().Tpl("城市")),
				app.Td().Body(app.Tpl().Tpl("销量")),
			),
			app.Tr().Tds(
				app.Td().Body(app.Tpl().Tpl("华北")).Rowspan(2),
				app.Td().Body(app.Tpl().Tpl("北京")),
				app.Td().Body(app.Tpl().Tpl("20")),
			),
			app.Tr().Tds(
				app.Td().Body(app.Tpl().Tpl("天津")),
				app.Td().Body(app.Tpl().Tpl("10")),
			),
		)
}

func condition(app *amisgo.App) comp.Service {
	return app.Service().
		Data(
			schema.Schema{"score": 40},
		).
		Body(
			app.TableView().Trs(
				app.Tr().Tds(
					app.Td().Background("${score > 50 ? 'green' : 'red'}").Body("score ${score} > 50 ?"),
					app.Td().Background("${score<100 ? 'green': 'red'}").Body("score ${score} < 100?"),
				),
			),
		)
}

func visibleOn(app *amisgo.App) comp.Page {
	return app.Page().Body(
		app.Switch().Label("显示第一行").Name("row1").Value(true),
		app.Switch().Label("显示北京单元格").Name("beijing").Value(true),
		app.TableView().Trs(
			app.Tr().Background("#f7f7f7").VisibleOn("row1").
				Tds(app.Td().Body("地区"), app.Td().Body("城市"), app.Td().Body("销量")),
			app.Tr().Tds(app.Td().Body("华北"), app.Td().VisibleOn("beijing").Body("北京")),
		),
	)
}

func layout(app *amisgo.App) comp.TableView {
	return app.TableView().Border(true).Trs(
		app.Tr().Background("orange").Tds(app.Td().Colspan(4).Align("center").Body("Header")),
		app.Tr().Tds(
			app.Td().Background("gray").Width(200).Body("Side").Rowspan(2),
			app.Td().Background("pink").Align("center").Body("Top Right").Colspan(3),
		),
		app.Tr().Background("lightgreen").Height(200).Tds(
			app.Td().Align("center").Body("C1"),
			app.Td().Align("center").Body("C2"),
			app.Td().Align("center").Body("C3"),
		),
		app.Tr().Background("lightblue").Tds(app.Td().Colspan(4).Align("center").Body("Footer")),
	)
}

func hanoi(app *amisgo.App) comp.TableView {
	trs := make([]comp.Tr, maxDiskCount)
	trs[0] = placeholder(app)
	for i := range trs[1:] {
		trs[i+1] = row(app, i)
	}
	return app.TableView().Border(false).Trs(trs...)
}

const maxDiskCount = 6

var colors = []string{
	"red", "green", "blue", "yellow", "brown", "pink",
}

func row(app *amisgo.App, disk int) comp.Tr {
	tds := make([]comp.Td, 0, 2*maxDiskCount)
	blanks := maxDiskCount - disk - 1
	for range blanks {
		tds = append(tds, app.Td())
	}
	for i := 0; i < 2*(disk+1); i++ {
		tds = append(tds, app.Td().Background(colors[disk]))
	}
	for range blanks {
		tds = append(tds, app.Td())
	}
	return app.Tr().Tds(tds...)
}

func placeholder(app *amisgo.App) comp.Tr {
	tds := make([]comp.Td, maxDiskCount*2)
	for i := range tds {
		tds[i] = app.Td()
	}
	return app.Tr().Tds(tds...)
}
