package storage

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/ricardohsd/simple-db/protocol"
)

type kv struct {
	db       map[string]string
	protocol protocol.Protocol
}

// NewKV returns a key value storage
func NewKV() *kv {
	p := &protocol.KV{}
	return &kv{
		db:       make(map[string]string),
		protocol: p,
	}
}

func (s *kv) Execute(message string) (string, error) {
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

func (s *kv) set(cmd *protocol.Command) (string, error) {
	log.Printf("Processing SET %v, %v\n", cmd.Key, cmd.Value)

	s.db[cmd.Key] = cmd.Value.(string)

	return "OK", nil
}

func (s *kv) get(cmd *protocol.Command) (string, error) {
	log.Printf("Processing GET %v\n", cmd.Key)

	m, ok := s.db[cmd.Key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	return m, nil
}

func (s *kv) del(cmd *protocol.Command) (string, error) {
	log.Printf("Processing DEL %v\n", cmd.Key)

	_, ok := s.db[cmd.Key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	delete(s.db, cmd.Key)

	return "OK", nil
}
