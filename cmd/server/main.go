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
var engine *string

func init() {
	port = flag.String("port", "3000", "Port to start server")
	engine = flag.String("engine", "key-value", "Storage engine to use")
	flag.Parse()
}

func main() {
	var db storage.Engine

	switch *engine {
	case "key-value":
		db = storage.NewKV()
	case "rolling-window":
		db = storage.NewRWindow()
	default:
		log.Fatalf("value %v isn't a valid engine", *engine)
	}

	address := net.JoinHostPort("127.0.0.1", *port)

	log.Printf("Starting server on address %v, with engine %v", address, *engine)

	stop := make(chan os.Signal)

	signal.Notify(stop, os.Interrupt)

	s, err := server.New(address, db)
	if err != nil {
		log.Fatalln(err)
	}

	go s.HandleConnections()

	<-stop

	log.Println("Quitting")
}
