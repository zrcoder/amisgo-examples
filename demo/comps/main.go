package main

import (
	"github.com/zrcoder/amisgo-examples/demo/comps/amis"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
	"github.com/zrcoder/amisgo-examples/demo/comps/tableview"
)

func main() {
	app.Register("Amis", amis.Demos)
	app.Register("Table View", tableview.Demos)
	app.Run()
}
