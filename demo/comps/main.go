package main

import (
	"github.com/zrcoder/amisgo-examples/demo/comps/amis"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
	"github.com/zrcoder/amisgo-examples/demo/comps/collapse"
	"github.com/zrcoder/amisgo-examples/demo/comps/divider"
	"github.com/zrcoder/amisgo-examples/demo/comps/flex"
	"github.com/zrcoder/amisgo-examples/demo/comps/tableview"
)

func main() {
	app.Register("Amis", amis.Demos)
	app.Register("Table View", tableview.Demos)
	app.Register("Collapse Goup", collapse.Demos)
	app.Register("Divider", divider.Demos)
	app.Register("Flex", flex.Demos)
	app.Run()
}
