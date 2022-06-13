package main

import (
	"context"
	"log"
	"luminnovel/internal/app"
	"net/http"
)

func main() {
	ctx := context.Background()

	app.InitHTTP(ctx)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on Port 8000")
}
