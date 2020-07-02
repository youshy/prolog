package main

import (
	"log"

	"github.com/youshy/prolog/internal/server"
)

func main() {
	port := ":8080"
	server := server.NewHTTPServer(port)
	log.Printf("server listening on port %v\n", port)
	log.Fatal(server.ListenAndServe())
}
