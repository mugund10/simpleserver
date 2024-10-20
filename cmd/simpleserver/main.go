package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mugund10/simpleserver/pkg/checker"
	"github.com/mugund10/simpleserver/pkg/middlewares"
)

func main() {
	// a custom middleware stack
	Mstack := middlewares.MakeStack(checker.CheckSubdomain)

	// custom multiplexer for routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", DefaultHandler)

	// custom server
	server := http.Server{
		Addr:    ":443",
		Handler: Mstack(mux),
	}

	// server starts
	log.Println("[INFO] server is running on port 443")
	err := server.ListenAndServe()
	if err != nil {
		log.Println("[ERROR] ", err)
	}

}

// a simple handler
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	// write reverse proxy
	fmt.Fprintf(w, "%v", "Default handler")
}
