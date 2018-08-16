package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var address *string

func init() {
	address = flag.String("port", "127.0.0.1:3000", "Port to create connection")
	flag.Parse()
}

func main() {
	cl, err := NewClient(*address)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	cl.HandleConnection()
}

type client struct {
	address string
	conn    net.Conn
}

func NewClient(address string) (*client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("dialing %v failed", address)
	}

	return &client{
		address: address,
		conn:    conn,
	}, nil
}

func (c *client) HandleConnection() {
	for {
		resp := bufio.NewReader(c.conn)

		input := bufio.NewReader(os.Stdin)

		fmt.Print("> ")

		text, err := input.ReadBytes(byte('\n'))
		switch err {
		case nil:
			c.conn.Write(text)
			readMessage(resp)
		case io.EOF:
			os.Exit(1)
		default:
			log.Fatalf("Error %v", err)
		}
	}
}

func readMessage(resp *bufio.Reader) {
	message, err := resp.ReadString('\n')
	switch err {
	case nil:
		fmt.Print(" " + message)
		break
	case io.EOF:
		log.Println("EOF. Closing connection")
		os.Exit(1)
	default:
		log.Printf("Error: %v", err)
	}
}
