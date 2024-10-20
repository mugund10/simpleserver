package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mugund10/simpleserver/pkg/checker"
	"github.com/mugund10/simpleserver/pkg/middlewares"
	"github.com/mugund10/simpleserver/pkg/readers"
)



func main() {
	sd := readers.GetServerS()
	port := fmt.Sprintf(":%v", sd[0].Port)

	// a custom middleware stack
	Mstack := middlewares.MakeStack(checker.CheckSubdomain)

	// custom multiplexer for routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", DefaultHandler)

	// custom server
	server := http.Server{
		Addr:    port,
		Handler: Mstack(mux),
	}

	// server starts
	log.Println("[INFO] server is running on port ", port)
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
