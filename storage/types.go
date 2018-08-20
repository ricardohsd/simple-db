package storage

import (
	"github.com/pkg/errors"
	"github.com/ricardohsd/simple-db/protocol"
)

var ErrWrongInstruction = errors.Errorf("wrong instruction")

// DB defines required methods for database storage
type DB interface {
	Execute(message string) (string, error)
	Set(cmd *protocol.Command) (string, error)
	Get(cmd *protocol.Command) (string, error)
	Del(cmd *protocol.Command) (string, error)
}
