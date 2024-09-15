package main

import (
	"github.com/zrcoder/amisgo"
)

func main() {
	ServeApi()

	amisgo.Serve("/", page)
	panic(amisgo.ListenAndServe(":80"))
}
