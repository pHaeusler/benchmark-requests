package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s: %s %s\n", r.RemoteAddr, r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "ok"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Starting server on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("ListenAndServe: %s", err)
	}
}
