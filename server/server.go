package server

import (
	"bufio"
	"io"
	"log"
	"net"

	"github.com/pkg/errors"
	"github.com/ricardohsd/simple-db/protocol"
	"github.com/ricardohsd/simple-db/storage"
)

type Server struct {
	address string
	ln      net.Listener
	storage storage.DB
}

func New(address string) (*Server, error) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, errors.Wrap(err, "starting server failed on address")
	}

	s := storage.New(&protocol.KV{})

	return &Server{address, ln, s}, nil
}

func (s *Server) HandleConnections() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Printf("ERROR Accept: %v", err)
			continue
		}

		go s.handleMessages(conn)
	}
}

func (s *Server) handleMessages(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		message, err := r.ReadString('\n')
		switch err {
		case nil:
			resp, err := s.storage.Execute(string(message))
			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
				break
			}

			conn.Write([]byte(resp + "\n"))
			break
		case io.EOF:
			log.Println("EOF. Closing connection")
			return
		default:
			log.Printf("Error: %v", err)
		}
	}
}
