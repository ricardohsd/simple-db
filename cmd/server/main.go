package main

import (
	"flag"
	"log"

	"github.com/ricardohsd/simple-db/server"
)

var address *string

func init() {
	address = flag.String("port", ":3000", "Port to start server")
	flag.Parse()
}

func main() {
	log.Printf("Starting server on address %v", *address)

	s, err := server.New(address)
	if err != nil {
		log.Fatalln(err)
	}

	s.HandleConnections()
}
