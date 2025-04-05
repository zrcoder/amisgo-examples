package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zrcoder/amisgo"
)

var (
	app      *amisgo.App
	upgrader = websocket.Upgrader{}
	wsConn   *websocket.Conn
)

func main() {
	app = amisgo.New()
	index := app.Page().Body(
		app.Service().Name("ui").Ws("/ws").Body(
			app.Amis().Name("main"),
		),
	)
	app.Mount("/", index)
	app.HandleFunc("/ws", wsHandler)
	app.Run(":8080")
}

func getDynamicUI() any {
	return app.Tpl().Tpl(fmt.Sprintf("Now: %s", time.Now().Format("15:04:05")))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	wsConn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer wsConn.Close()
	tk := time.NewTicker(1 * time.Second)
	for range tk.C {
		_ = wsConn.WriteJSON(map[string]any{
			"main": getDynamicUI(),
		})
	}
	select {}
}
