package main

import "net/http"

func handleReady(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, "Ready")
}
