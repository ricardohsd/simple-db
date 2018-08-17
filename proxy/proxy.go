package proxy

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"

	"github.com/pkg/errors"
	"github.com/ricardohsd/simple-db/protocol"
)

type proxy struct {
	backendAddr string
	protocol    protocol.Protocol
	stats       *statistics
	ln          net.Listener
}

// NewProxy creates a tcp server and connects to a tcp backend.
// It accepts incoming requests from clients and forwards them to the backend.
// The proxy will collect metrics for every command being forwarded.
func NewProxy(listenAddr string, backendAddr string) (*proxy, error) {
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, errors.Wrap(err, "starting server failed on address")
	}

	return &proxy{
		backendAddr: backendAddr,
		protocol:    &protocol.KV{},
		stats:       NewStatistics(),
		ln:          ln,
	}, nil
}

func (p *proxy) HandleIncomingConnections() {
	for {
		conn, err := p.ln.Accept()
		if err != nil {
			log.Printf("ERROR Accept: %v", err)
			continue
		}

		go p.handleClientMessages(conn)
	}
}

func (p *proxy) handleClientMessages(clientConn net.Conn) {
	defer clientConn.Close()

	backendConn, err := net.Dial("tcp", p.backendAddr)
	if err != nil {
		log.Fatalf("dialing %v failed\n", p.backendAddr)
	}

	buff := bufio.NewReader(clientConn)

	for {
		log.Printf("Statistics %v\n", p.stats.GetAll())

		resp := bufio.NewReader(backendConn)

		message, err := buff.ReadString('\n')
		switch err {
		case nil:
			p.aggregateStats(message)

			backendConn.Write([]byte(message + "\n"))

			forwardMessageToClient(resp, clientConn)
			break
		case io.EOF:
			log.Println("EOF. Closing connection")
			return
		default:
			log.Printf("Error: %v", err)
		}
	}
}

func (p *proxy) aggregateStats(message string) {
	cmd, err := p.protocol.Parse(message)
	if err != nil {
		log.Println("Error when parsing protocol", err)
		return
	}

	p.stats.Incr(cmd.Instruction)
}

func forwardMessageToClient(resp *bufio.Reader, clientConn net.Conn) {
	message, err := resp.ReadString('\n')
	switch err {
	case nil:
		log.Printf("Forwarding message %v to client\n", strings.TrimSpace(message))

		clientConn.Write([]byte(message + "\n"))
		break
	case io.EOF:
		log.Fatalf("EOF. Closing connection")
	default:
		log.Printf("Error: %v", err)
	}
}
