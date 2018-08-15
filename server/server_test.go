package server

import (
	"bufio"
	"io"
	"net"
	"strings"
	"testing"

	"github.com/ricardohsd/simple-db/protocol"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	address := ":3001"

	db := &fakeDB{}

	s, err := newTestServer(address, db)
	if err != nil {
		t.Errorf("failed to create listener")
	}

	go s.HandleConnections()

	assertReqResp(t, address, []string{
		"MESSAGE 1", "MESSAGE 2", "MESSAGE 3",
	})
}

func newTestServer(address string, db *fakeDB) (*Server, error) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Server{&address, ln, db}, nil
}

func assertReqResp(t *testing.T, address string, messages []string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	for _, message := range messages {
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			return err
		}

		resp := bufio.NewReader(conn)
		r, err := resp.ReadString('\n')
		switch err {
		case nil:
			r = strings.TrimSpace(r)
			r = strings.ToUpper(r)
			assert.Equal(t, message, r)
			break
		case io.EOF:
			return io.EOF
		default:
			return err
		}
	}

	return nil
}

type fakeDB struct{}

func (f *fakeDB) Execute(message string) (string, error) {
	return message, nil
}

func (f *fakeDB) Set(cmd *protocol.Command) (string, error) {
	return "", nil
}

func (f *fakeDB) Get(cmd *protocol.Command) (string, error) {
	return "", nil
}

func (f *fakeDB) Del(cmd *protocol.Command) (string, error) {
	return "", nil
}
