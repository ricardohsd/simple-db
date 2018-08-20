package storage

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/ricardohsd/simple-db/protocol"
)

type storage struct {
	db       map[string]string
	protocol protocol.Protocol
}

// New returns a key value storage
func New() *storage {
	p := &protocol.KV{}
	return &storage{
		db:       make(map[string]string),
		protocol: p,
	}
}

func (s *storage) Execute(message string) (string, error) {
	cmd, error := s.protocol.Parse(message)
	if error != nil {
		return "", error
	}

	switch cmd.Instruction {
	case "GET":
		return s.get(cmd)
	case "SET":
		return s.set(cmd)
	case "DEL":
		return s.del(cmd)
	default:
		return "", errors.Wrap(ErrWrongInstruction, cmd.Instruction)
	}
}

func (s *storage) set(cmd *protocol.Command) (string, error) {
	log.Printf("Processing SET %v, %v\n", cmd.Key, cmd.Value)

	s.db[cmd.Key] = cmd.Value

	return "OK", nil
}

func (s *storage) get(cmd *protocol.Command) (string, error) {
	log.Printf("Processing GET %v\n", cmd.Key)

	m, ok := s.db[cmd.Key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	return m, nil
}

func (s *storage) del(cmd *protocol.Command) (string, error) {
	log.Printf("Processing DEL %v\n", cmd.Key)

	_, ok := s.db[cmd.Key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	delete(s.db, cmd.Key)

	return "OK", nil
}
