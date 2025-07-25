package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dist/task_traker.yaml")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
