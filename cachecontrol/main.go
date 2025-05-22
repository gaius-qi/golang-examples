package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/large", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request for /large")
		w.Header().Set("Cache-Control", "max-age=60")
		http.Redirect(w, r, "http://localhost:8081/large", http.StatusTemporaryRedirect)
	})

	log.Println("Server starting on :8082...")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
