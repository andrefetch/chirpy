package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := ":8080"

	fileServer := http.FileServer(http.Dir("."))

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Printf("Starting server on port: %s", port)
	log.Fatal(server.ListenAndServe())
}
