package main

import (
	"log"
	"net/http"
)

func main() {
	//https://datatracker.ietf.org/doc/html/rfc6455

	setupAPI()
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func setupAPI() {

	var manager *Manager

	manager = NewManager()

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
}
