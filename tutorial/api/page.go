package main

import (
	"github.com/zrcoder/amisgo/comp"
)

/*
Page

	├── Toolbar
	│  └─ Form 顶部表单项
	├── Grid // 用于水平布局
	│  ├─ Panel
	│  │  └─ Tabs
	│  │    └─ Chart
	│  └─ Panel
	│     └─ Chart
	└── CRUD
*/
var page = comp.Page().
	Toolbar(
		comp.Form().StaticLabelClassName("").
			PanelClassName("mb-0").
			Body(
				comp.Select().
					Label("区域").
					Name("businessLineId").
					SelectFirst(true).
					Mode("inline").
					Options(
						"北京",
						"上海",
					).
					CheckAll(false),
				comp.InputDateRange().
					Label("时间范围").
					Name("dateRange").
					Inline(true).
					Value("-1month,+0month").
					InputFormat("YYYY-MM-DD").
					CloseOnSelect(true).
					Clearable(false),
			).
			Mode("inline").
			Target("mainPage").
			SubmitOnChange(true).
			SubmitOnInit(true),
	).
	Body(
		comp.Grid().
			Columns(
				comp.Panel().
					ClassName("h-full").
					Body(
						comp.Tabs().
							Tabs(
								comp.Tab().
									Title("消费趋势").
									Tab(
										comp.Chart().
											Config(trendChartCOnfig),
									),
								comp.Tab().
									Title("账户余额").
									Tab("0.00"),
							),
					),
				comp.Panel().
					ClassName("h-full").
					Body("Test"),
			),
		comp.CrudTable().
			ClassName("m-t-sm").
			Api(itemsRouter). // 这里是演示 Api，也可以用 FetchData 方法，比 Api 简洁很多，不需要 ServeApi, 如下：
			// FetchData(func() any {
			// 	// 仅演示，直接给出数据，实际可以从数据库得到数据
			// 	return []Item{
			// 		{Engine: "e1", Browser: "chrome", Platform: "windows", Version: "1.0"},
			// 		{Browser: "safri", Platform: "macOS", Version: "2.0"},
			// 	}
			// }).
			Columns(
				comp.Column().Name("id").Label("ID"),
				comp.Column().Name("engine").Label("Rendering engine"),
				comp.Column().Name("browser").Label("Browser"),
				comp.Column().Name("platform").Label("Platform(s)"),
				comp.Column().Name("version").Label("Engine version"),
				comp.Column().Name("grade").Label("CSS grade"),
			),
	)
