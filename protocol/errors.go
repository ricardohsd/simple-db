package protocol

import "github.com/pkg/errors"

var ErrMalformedCommand = errors.Errorf("malformed command")
var ErrUnknownCommand = errors.Errorf("unknown command")
