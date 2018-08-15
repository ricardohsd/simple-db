package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	p := &KV{}

	cmd, err := p.Parse("GET name")
	assert.Nil(t, err)
	assert.Equal(t, &Command{"GET", "name", ""}, cmd)

	_, err = p.Parse("GET ")
	assert.Equal(t, ErrMalformedCommand, err)

	cmd, err = p.Parse("SET name john")
	assert.Nil(t, err)
	assert.Equal(t, &Command{"SET", "name", "john"}, cmd)

	_, err = p.Parse("SET name")
	assert.Equal(t, ErrMalformedCommand, err)

	_, err = p.Parse("SET")
	assert.Equal(t, ErrMalformedCommand, err)

	cmd, err = p.Parse("DEL name")
	assert.Nil(t, err)
	assert.Equal(t, &Command{"DEL", "name", ""}, cmd)

	_, err = p.Parse("DEL ")
	assert.Equal(t, ErrMalformedCommand, err)

	_, err = p.Parse("DEL")
	assert.Equal(t, ErrMalformedCommand, err)

	_, err = p.Parse("E ")
	assert.Equal(t, err.Error(), "E: unknown command")
}
