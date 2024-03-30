package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the client folder path
	clientFolder := "client" // Replace with your actual folder name

	// Create the file server handler, pointing to the client folder
	fileServer := http.FileServer(http.Dir(clientFolder))

	// Handle all requests with the file server
	http.Handle("/", fileServer)

	// Start the server on port 8080
	port := 9090
	fmt.Println("File server listening on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
