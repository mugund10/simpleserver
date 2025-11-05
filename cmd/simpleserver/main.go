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
		url, _ := url.Parse(fmt.Sprintf("http://172.17.0.1:%d", dat.Port))
		proxies[dat.Subdomain] = httputil.NewSingleHostReverseProxy(url)
	}
	port := fmt.Sprintf(":%v", sd[0].Port)

	// a custom middleware stack
	Mstack := middlewares.MakeStack(checker.CheckSubdomain)

	// custom multiplexer for routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.Host == "mugund10.dev" {
			fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>mugund10.dev</title>
    <style>
      body {
        background-color: #fff;
        color: #000;
        font-family: Arial, sans-serif;
        margin: 0;
        padding: 40px;
      }
      .container {
        max-width: 600px;
        margin: auto;
        text-align: center;
      }
      a {
        color: #000;
        text-decoration: underline;
      }
      h1 {
        font-size: 2.5em;
        margin-bottom: 0.5em;
      }
      p {
        font-size: 1.1em;
        line-height: 1.6;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <h1>mugund10.dev</h1>
      <p>powered by <strong>SimpleServer</strong>.</p>
      <p>
        Source code is available on
        <a href="https://github.com/mugund10/simpleserver/" target="_blank">GitHub</a>.
      </p>
      <p>
        <a href="https://blog.of.mugund10.dev" target="_blank">blog @ blog.of.mugund10.dev</a>
      </p>
    </div>
  </body>
</html>`)

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
