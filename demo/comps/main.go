package main

import (
	"github.com/zrcoder/amisgo-examples/demo/comps/amis"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
	"github.com/zrcoder/amisgo-examples/demo/comps/collapse"
	"github.com/zrcoder/amisgo-examples/demo/comps/divider"
	"github.com/zrcoder/amisgo-examples/demo/comps/flex"
	"github.com/zrcoder/amisgo-examples/demo/comps/grid"
	"github.com/zrcoder/amisgo-examples/demo/comps/grid2d"
	"github.com/zrcoder/amisgo-examples/demo/comps/hbox"
	"github.com/zrcoder/amisgo-examples/demo/comps/tableview"
)

func main() {
	app.Register("Amis", amis.Demos)
	app.Register("Table View", tableview.Demos)
	app.Register("Collapse Goup", collapse.Demos)
	app.Register("Divider", divider.Demos)
	app.Register("Flex", flex.Demos)
	app.Register("Grid", grid.Demos)
	app.Register("Grid2d", grid2d.Demos)
	app.Register("Hbox", hbox.Demos)
	app.Run()
}
