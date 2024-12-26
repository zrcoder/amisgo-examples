package api

import (
	"net/http"

	"github.com/zrcoder/amisgo/comp"
)

const (
	InitDateErrorPath = "/initDateError"
)

func InitDateError(w http.ResponseWriter, r *http.Request) {
	resp := comp.Response{
		Status: http.StatusInternalServerError,
		Msg:    "测试返回错误",
	}
	w.Write(resp.Json())
}
