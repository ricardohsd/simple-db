package protocol

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// RWindow is a basic text human readable protocol for a rolling window storage
type RWindow struct {
}

// Parse parses a message and validates it into GET, SET or DEL commands
func (r *RWindow) Parse(message string) (*Command, error) {
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
	case "RWSET":
		return r.setCommand(key, value)
	case "RWADD":
		return r.addCommand(key, value)
	case "RWAVG":
		return r.avgCommand(key, value)
	case "RWDEL":
		return r.delCommand(key, value)
	default:
		return nil, errors.Wrap(ErrUnknownCommand, inst)
	}
}

func (r *RWindow) avgCommand(key string, value string) (*Command, error) {
	if len(key) == 0 {
		return nil, ErrMalformedCommand
	}

	return &Command{"RWAVG", key, value}, nil
}

func (r *RWindow) addCommand(key string, value string) (*Command, error) {
	if len(key) == 0 || len(value) == 0 {
		return nil, ErrMalformedCommand
	}

	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, errors.New("value must be a valid float")
	}

	return &Command{"RWADD", key, val}, nil
}

func (r *RWindow) setCommand(key string, value string) (*Command, error) {
	if len(key) == 0 || len(value) == 0 {
		return nil, ErrMalformedCommand
	}

	seconds, err := strconv.Atoi(value)
	if err != nil {
		return nil, errors.New("value must be a valid integer")
	}

	return &Command{"RWSET", key, seconds}, nil
}

func (r *RWindow) delCommand(key string, value string) (*Command, error) {
	if len(key) == 0 {
		return nil, ErrMalformedCommand
	}

	return &Command{"RWDEL", key, value}, nil
}
