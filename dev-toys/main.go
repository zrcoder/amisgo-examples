package main

import (
	"log"

	"github.com/zrcoder/amisgo"
)

func main() {
	setupRoutes()
	startServer()
}

func startServer() {
	port := ":8888"
	log.Printf("Starting server on http://localhost%s\n", port)

	if err := amisgo.ListenAndServe(port, appConfig); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
