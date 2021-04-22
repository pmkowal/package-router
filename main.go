package main

import (
	"log"
	"net/http"
	"packageRouter/internal/handlers"
)

const routesPath = "/routes"
const port = ":8080"

func main() {
	http.HandleFunc(routesPath, handlers.RoutesHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
