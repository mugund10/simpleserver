package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/mugund10/simpleserver/internal/checker"
	"github.com/mugund10/simpleserver/internal/readers"
	"github.com/mugund10/simpleserver/pkg/middlewares"
)

type Subdomain struct {
	Name string
}

func main() {
	proxies := make(map[string]*httputil.ReverseProxy)
	sd := readers.GetServerS()
	pd := readers.Getproxies()
	for _, dat := range pd {
		url, _ := url.Parse(fmt.Sprintf("http://172.17.0.1:%d", dat.Port))
		proxies[dat.Subdomain] = httputil.NewSingleHostReverseProxy(url)
	}
	port := fmt.Sprintf(":%v", sd[0].Port)

	// a custom middleware stack
	Mstack := middlewares.MakeStack(checker.CheckSubdomain)

	// var subs []Subdomain
	// for _, rpx := range pd {
	// 	full := fmt.Sprintf("%s.%s", rpx.Subdomain, sd[0].Domain)
	// 	subs = append(subs, Subdomain{Name: full})
	// }
	// tmpl, err := template.ParseFiles("templates/index.html")
	// if err != nil {
	// 	log.Println("Template error", 500)
	// 	return
	// }
	// custom multiplexer for routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Host == "mugund10.dev" {
			proxies["root"].ServeHTTP(w, r)
			//tmpl.Execute(w, subs)
		} else {
			sdom := strings.Split(r.Host, ".")
			proxies[sdom[0]].ServeHTTP(w, r)
		}
	})

	// custom server
	server := http.Server{
		Addr:    port,
		Handler: Mstack(mux),
	}

	// server starts
	log.Println("[INFO] server is running on port ", port)
	err := server.ListenAndServeTLS("fullchain.pem", "privkey.pem")
	if err != nil {
		log.Println("[ERROR] ", err)
	}

}

// a simple handler
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	// to do write reverse proxy
	fmt.Fprintf(w, "%v", "Default handler")
}
