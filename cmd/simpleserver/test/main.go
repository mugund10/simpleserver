package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request for domain: %s", r.Host)
		fmt.Fprintf(w, "Hello! You reached domain: %s\n", r.Host)
	})

	server := &http.Server{
		Addr:    ":443",
		Handler: handler,
	}

	certFile := "fullchain.pem"  // your local copy
	keyFile := "privkey.pem"     // your local copy

	log.Println("Starting HTTPS server on port 443...")
	if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
