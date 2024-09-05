package main

import (
	"fmt"
	"io"
	"net/http"
)

const logApiPath = "/test_log/"

var logWriter io.Writer

func init() {
	http.HandleFunc(logApiPath, func(w http.ResponseWriter, r *http.Request) {
		logWriter = w
		for {
			w.(http.Flusher).Flush()
		}
	})
}

// 要打印日志时调这个函数，日志会渲染到 logview 中
func Logf(format string, v ...any) {
	if logWriter == nil {
		return
	}
	_, err := fmt.Fprintf(logWriter, format, v...)
	if err != nil {
		panic(err)
	}
}
