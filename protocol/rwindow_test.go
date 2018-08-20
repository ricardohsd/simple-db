package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRWindowParse(t *testing.T) {
	r := &RWindow{}

	cmd, err := r.Parse("RWAVG transactions")
	assert.Nil(t, err)
	assert.Equal(t, &Command{"RWAVG", "transactions", ""}, cmd)

	_, err = r.Parse("RWAVG ")
	assert.Equal(t, ErrMalformedCommand, err)

	cmd, err = r.Parse("RWSET transactions 10")
	assert.Nil(t, err)
	assert.Equal(t, &Command{"RWSET", "transactions", "10"}, cmd)

	_, err = r.Parse("RWSET transactions")
	assert.Equal(t, ErrMalformedCommand, err)

	_, err = r.Parse("RWSET")
	assert.Equal(t, ErrMalformedCommand, err)

	cmd, err = r.Parse("RWADD transactions 10.20")
	assert.Nil(t, err)
	assert.Equal(t, &Command{"RWADD", "transactions", "10.20"}, cmd)

	cmd, err = r.Parse("RWDEL transactions")
	assert.Nil(t, err)
	assert.Equal(t, &Command{"RWDEL", "transactions", ""}, cmd)

	_, err = r.Parse("RWDEL ")
	assert.Equal(t, ErrMalformedCommand, err)

	_, err = r.Parse("RWDEL")
	assert.Equal(t, ErrMalformedCommand, err)

	_, err = r.Parse("E ")
	assert.Equal(t, err.Error(), "E: unknown command")
}
