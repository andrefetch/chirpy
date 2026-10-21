package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := ":8080"

	mux := http.NewServeMux()
	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Printf("Starting server on port: %s", port)
	log.Fatal(server.ListenAndServe())
}
