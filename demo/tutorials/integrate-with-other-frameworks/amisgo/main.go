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
	app.Handle(apiPrefix, g) // 将 g 挂载到 /api/ 路径
	app.Mount("/", app.Page().InitApi(dateApi).Body("Now: ${date}"))
	// 启动服务
	panic(app.Run(":8888"))
}
