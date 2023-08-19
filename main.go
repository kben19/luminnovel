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

	log.Println("Listening on Port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
