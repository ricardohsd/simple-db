package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/ricardohsd/simple-db/server"
)

var address *string

func init() {
	address = flag.String("port", ":3000", "Port to start server")
	flag.Parse()
}

func main() {
	log.Printf("Starting server on address %v", *address)

	stop := make(chan os.Signal)

	signal.Notify(stop, os.Interrupt)

	s, err := server.New(address)
	if err != nil {
		log.Fatalln(err)
	}

	go s.HandleConnections()

	<-stop

	log.Println("Quitting")
}
