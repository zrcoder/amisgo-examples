package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zrcoder/amisgo"
)

const (
	apiPrefix = "/api/"
	datePath  = "date"
	dateApi   = apiPrefix + datePath
)

func main() {
	// 初始化 gin
	g := gin.Default()
	g.GET(dateApi, func(c *gin.Context) {
		c.JSON(200, gin.H{"date": time.Now()})
	})
	// 初始化 amisgo
	app := amisgo.New()
	app.Mount("/", app.Page().InitApi(dateApi).Body("Now: ${date}"))
	// 使用标准库的 http 包
	http.Handle("/", app)
	http.Handle(apiPrefix, g)
	panic(http.ListenAndServe(":8888", nil))
}
