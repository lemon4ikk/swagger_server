package main

import (
	"log"
	"net/http"
)

func main() {
	swaggerUIPath := "./dist"

	http.Handle("/", http.FileServer(http.Dir(swaggerUIPath)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
