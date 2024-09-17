package main

import (
	"fmt"

	"github.com/zrcoder/amisgo"
)

const (
	fmtPath    = "/fmt"
	convPath   = "/conv"
	genPath    = "/gen"
	chartPath  = "/chart"
	encDecPath = "/enc"
)

func main() {
	// app := comp.App().Logo("/assets/gopher.svg").BrandName("Dev Toys").Header(Header).Pages(pages)
	// amisgo.Serve("/", app)
	amisgo.Serve(fmtPath, page(formatters))
	amisgo.Serve(convPath, page(converters))
	amisgo.Serve(genPath, page(generaters))
	amisgo.Serve(chartPath, page(charts))
	amisgo.Serve(encDecPath, page(encoders))
	amisgo.Redirect("/", fmtPath)
	fmt.Println("Serving on http://localhost")
	panic(amisgo.ListenAndServe(":80", appConfig))
}
