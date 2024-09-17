package main

import (
	"fmt"

	"github.com/zrcoder/amisgo"
)

func main() {
	amisgo.Redirect("/", fmtPath)
	amisgo.Serve(fmtPath, page(formatters))
	amisgo.Serve(convPath, page(converters))
	amisgo.Serve(genPath, page(generaters))
	amisgo.Serve(chartPath, page(charts))
	amisgo.Serve(encDecPath, page(encoders))
	fmt.Println("Serving on http://localhost")
	panic(amisgo.ListenAndServe(":80", appConfig))
}
