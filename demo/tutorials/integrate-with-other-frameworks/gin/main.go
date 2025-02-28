package main

import (
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
	// 将 amisgo 包装为 gin 的 HandlerFunc
	g.GET("/", func(c *gin.Context) {
		app.ServeHTTP(c.Writer, c.Request)
	})
	// 启动 gin
	panic(g.Run(":8888"))
}
