package storage

import (
	"github.com/pkg/errors"
)

var ErrWrongInstruction = errors.Errorf("wrong instruction")

// Engine defines required methods for database engine
type Engine interface {
	Execute(message string) (string, error)
}
