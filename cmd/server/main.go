package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ricardohsd/simple-db/server"
	"github.com/ricardohsd/simple-db/storage"
)

var port *string

func init() {
	port = flag.String("port", "3000", "Port to start server")
	flag.Parse()
}

func main() {
	address := net.JoinHostPort("127.0.0.1", *port)

	log.Printf("Starting server on address %v", address)

	stop := make(chan os.Signal)

	signal.Notify(stop, os.Interrupt)

	kvs := storage.NewKV()
	s, err := server.New(address, kvs)
	if err != nil {
		log.Fatalln(err)
	}

	go s.HandleConnections()

	<-stop

	log.Println("Quitting")
}
