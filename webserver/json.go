package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshall json response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(bytes)
	if err != nil {
		return
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}

	type errResponse struct {
		// Add JSON reflects tag so that we can tell the marshal function to marshal it with that key tag
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{Error: msg})
}
