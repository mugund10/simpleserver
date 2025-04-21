package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/mugund10/simpleserver/pkg/checker"
	"github.com/mugund10/simpleserver/pkg/middlewares"
	"github.com/mugund10/simpleserver/pkg/readers"
)

func main() {
	proxies := make(map[string]*httputil.ReverseProxy)
	sd := readers.GetServerS()
	pd := readers.Getproxies()
	for _, dat := range pd {
		url, _ := url.Parse(fmt.Sprintf("http://localhost:%d", dat.Port))
		proxies[dat.Subdomain] = httputil.NewSingleHostReverseProxy(url)
	}
	port := fmt.Sprintf(":%v", sd[0].Port)

	// a custom middleware stack
	Mstack := middlewares.MakeStack(checker.CheckSubdomain)

	// custom multiplexer for routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sdom := strings.Split(r.Host, ".")
		proxies[sdom[0]].ServeHTTP(w, r)
	})

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
	// to do write reverse proxy
	fmt.Fprintf(w, "%v", "Default handler")
}
