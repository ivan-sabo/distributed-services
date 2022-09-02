package main

import (
	"log"

	"github.com/123sabo/distributed-services/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
