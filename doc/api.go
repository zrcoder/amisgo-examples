package main

import (
	"net/http"
	"strings"

	"github.com/zrcoder/amisgo-examples/doc/docs"
	"github.com/zrcoder/amisgo/comp"
)

const (
	docsApi  = "/docs"
	docQuery = "doc"
)

func getDoc(w http.ResponseWriter, r *http.Request) {
	docPath := r.URL.Query().Get(docQuery)
	if !strings.HasSuffix(docPath, ".md") {
		return
	}
	data, err := docs.FS.ReadFile(docPath)
	if err != nil {
		resp := comp.ErrorResponse(err.Error())
		w.Write(resp.Json())
		return
	}
	w.Header().Add("Content-Type", "text/markdown")
	w.Write(data)
}
