package main

import (
	"net/http"

	sdk "gitee.com/rdor/amis-sdk/v6"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

func main() {
	app := amisgo.New(conf.WithLocalSdk(http.FS(sdk.FS)))
	app.Mount("/", app.Page().Body("Hello, Amisgo!"))
	app.Run(":8080")
}
