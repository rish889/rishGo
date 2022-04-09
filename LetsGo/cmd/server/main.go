package main

import (
	"log"

	"github.com/travisjeffery/proglog/internal/server"
)

func main() {
	log.Print("=============Starting Server===============")
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
