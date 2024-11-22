package main

import (
	"net/http"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

const itemsRouter = "/items"

func main() {
	ag := amisgo.New().
		// Provide API endpoints for page calls.
		// In practice, consider using the advanced features of amisgo in the page (e.g., using GetData method instead of API) to eliminate the need for this API.
		HandleFunc(itemsRouter, func(w http.ResponseWriter, r *http.Request) {
			resp := comp.SuccessResponse("", comp.Data{"items": items})
			w.Write(resp.Json())
		}).
		Mount("/", page)

	panic(ag.Run())
}
