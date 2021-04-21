package main

import (
	"log"
	"net/http"
	"packageRouter/internal/handlers"
)

func main() {
	http.HandleFunc("/routes", handlers.RoutesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
