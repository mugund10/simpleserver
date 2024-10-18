package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mugund10/simpleserver/pkg/configs"
	"github.com/mugund10/simpleserver/pkg/filereaders"
)

func main() {
	var ServerConfig configs.ServersRoot

	file := filereaders.New()
	file.LoadServer(&ServerConfig)

	servers := ServerConfig.Servers
	fmt.Println("Number of servers:", len(servers))

	for i := 0; i < len(servers); i++ {
		svr := servers[i]
		go sooper(svr) 
	}

	select {}
}

func sooper(s configs.Server) {
	mux := http.NewServeMux()
	mux.HandleFunc(s.Server.Name, someHandler)

	server := &http.Server{
		Addr:    s.Server.Port,
		Handler: mux,
	}

	fmt.Printf("Starting server %s on port %s\n", s.Server.Name, s.Server.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error starting server:", err)
	}
}

func someHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handling path: %s", r.URL.Path)
}
