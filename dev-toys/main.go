package main

import (
	"log"
	"os"

	"github.com/zrcoder/amisgo-examples/dev-toys/routes"
)

func main() {
	app := routes.Setup()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	log.SetFlags(log.LstdFlags)
	log.Printf("Starting server on http://localhost:%s\n", port)
	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
