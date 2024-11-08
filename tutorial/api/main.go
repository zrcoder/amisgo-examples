package main

import (
	"encoding/json"
	"net/http"

	"github.com/zrcoder/amisgo"
)

const itemsRouter = "/items"

func main() {
	ag := amisgo.New().
		HandleFunc(itemsRouter, func(w http.ResponseWriter, r *http.Request) {
			data, _ := json.Marshal(Resp{Data: items})
			w.Write(data)
		}).
		Mount("/", page)
	panic(ag.Run())
}
