package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/ricardohsd/simple-db/proxy"
)

var listenAddr *string
var backendAddr *string

func main() {
	listenAddr = flag.String("listen", ":3001", "Port to listen for connections")
	backendAddr = flag.String("backend", "127.0.0.1:3000", "Backend port to connect")
	flag.Parse()

	log.Printf("Started proxy on address %v\n", *listenAddr)

	stop := make(chan os.Signal)

	signal.Notify(stop, os.Interrupt)

	cl, err := proxy.NewProxy(*listenAddr, *backendAddr)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	go cl.HandleIncomingConnections()

	<-stop

	log.Println("Quitting")
}
