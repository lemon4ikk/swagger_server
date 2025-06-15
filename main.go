package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("dist/task_traker.yaml")
		if err != nil {
			http.Error(w, "file not found", 404)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
