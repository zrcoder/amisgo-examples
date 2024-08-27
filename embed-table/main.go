package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	page := comp.Page().Body(
		comp.Service().
			// Data(items). // 这一行直接指定数据，静态
			FetchData(getData). // 这一行指定数据获取方法，方法里边可以动态实现，比如从数据库读取
			Body(
				comp.Table().Source("$rows").ClassName("m-b-none").ColumnsTogglable(false).Columns(
					comp.Column().Name("engine").Label("Engine"),
					comp.Column().Name("grade").Label("Grade"),
					comp.Column().Name("version").Label("Version"),
					comp.Column().Name("browser").Label("Browser"),
					comp.Column().Name("id").Label("ID"),
					comp.Column().Name("platform").Label("Platform"),
				),
			),
	)

	panic(amisgo.ListenAndServe(page))
}

type Items struct {
	Rows []Item `json:"rows"`
}

type Item struct {
	Engine   string `json:"engine"`
	Browser  string `json:"browser"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
	Grade    string `json:"grade"`
	ID       int    `json:"id"`
	Children []Item `json:"children"`
}

var items = Items{
	Rows: []Item{
		{
			Engine:   "Trident",
			Browser:  "IE 4.0",
			Platform: "Win 95+",
			Version:  "4",
			Grade:    "X",
			ID:       1,
			Children: []Item{
				{
					Engine:   "Trident",
					Browser:  "IE 4.0",
					Platform: "Win 95+",
					Version:  "4",
					Grade:    "Y",
					ID:       1001,
				},
				{
					Engine:   "Trident",
					Browser:  "IE 4.0",
					Platform: "Win 95+",
					Version:  "4",
					Grade:    "Z",
					ID:       1002,
				},
			},
		},
		{
			Engine:   "Trident",
			Browser:  "Safari 6.0",
			Platform: "Win 95+",
			Version:  "4",
			Grade:    "A",
			ID:       2,
			Children: []Item{
				{
					Engine:   "Trident",
					Browser:  "IE 4.0",
					Platform: "Win 95+",
					Version:  "4",
					Grade:    "Y",
					ID:       2001,
				},
				{
					Engine:   "Trident",
					Browser:  "IE 4.0",
					Platform: "Win 95+",
					Version:  "4",
					Grade:    "Z",
					ID:       2002,
				},
			},
		},
	},
}

func getData() any {
	return items
}
