package main

import (
	"log"
	"net/http"

	"template/internal/router"
)

func main() {
	log.Println("Application successfully started")

	mux := http.NewServeMux()
	router.RegisterRoutes(mux)

	port := ":8080"
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
