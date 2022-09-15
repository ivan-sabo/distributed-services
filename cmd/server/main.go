package main

import (
	"log"

	"github.com/ivan-sabo/distributed-services/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
