package protocol

import (
	"strings"

	"github.com/pkg/errors"
)

// KV is a basic text human readable protocol for a key value storage
type KV struct {
}

// Parse parses a message and validates it into GET, SET or DEL commands
func (s *KV) Parse(message string) (*Command, error) {
	args := strings.SplitN(message, " ", 3)
	if len(args) < 2 {
		return nil, ErrMalformedCommand
	}

	var inst, key, value string

	inst = strings.ToUpper(args[0])

	if len(args) >= 2 {
		key = strings.TrimSpace(args[1])
	}

	if len(args) == 3 {
		value = strings.TrimSpace(args[2])
	}

	switch inst {
	case "GET":
		return getCommand(key, value)
	case "SET":
		return setCommand(key, value)
	case "DEL":
		return delCommand(key, value)
	default:
		return nil, errors.Wrap(ErrUnknownCommand, inst)
	}
}

func getCommand(key string, value string) (*Command, error) {
	if len(key) == 0 {
		return nil, ErrMalformedCommand
	}

	return &Command{"GET", key, value}, nil
}

func setCommand(key string, value string) (*Command, error) {
	if len(key) == 0 || len(value) == 0 {
		return nil, ErrMalformedCommand
	}

	return &Command{"SET", key, value}, nil
}

func delCommand(key string, value string) (*Command, error) {
	if len(key) == 0 {
		return nil, ErrMalformedCommand
	}

	return &Command{"DEL", key, value}, nil
}
