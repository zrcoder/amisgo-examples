package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	page := comp.Page().Body(
		comp.Service().
			// Data(items). // This line directly specifies the data, static
			GetData(getData). // This line specifies the data retrieval method, which can be implemented dynamically, such as reading from a database
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

	ag := amisgo.New().Mount("/", page)
	panic(ag.Run())
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

func getData() (any, error) {
	return items, nil
}
