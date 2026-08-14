package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/hello", hello)

	serverPort := 8080
	log.Printf("Starting Simple Service on port %d\n", serverPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), serverMux); err != nil {
		log.Fatalf("HTTP ListenAndServe error: %v", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "hello world"})
}
