package main

import (
	"log"

	"mycineden/backend/internal/server"
)

func main() {
	app, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
