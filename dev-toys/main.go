package main

import (
	"fmt"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	app := comp.App().Logo("/assets/gopher.svg").BrandName("Dev Toys").Header(header).Pages(pages)
	amisgo.Serve("/", app)

	fmt.Println("Serving on http://localhost")
	panic(amisgo.ListenAndServe(":80", appConfig))
}
