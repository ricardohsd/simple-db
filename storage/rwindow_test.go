package storage

import (
	"testing"

	"github.com/ricardohsd/simple-db/protocol"
	"github.com/stretchr/testify/assert"
)

func TestRWindowExecute(t *testing.T) {
	rw := &rwindow{
		db:       make(map[string]rollingStatistics),
		protocol: &protocol.RWindow{},
	}

	resp, err := rw.Execute("RWAVG transactions")
	assert.NotNil(t, err, "key not found")
	assert.Equal(t, "", resp)

	resp, err = rw.Execute("RWSET transactions notANumber")
	assert.Equal(t, err.Error(), "value must be a valid integer")
	assert.Equal(t, "ERROR", resp)

	resp, err = rw.Execute("RWSET transactions 60")
	assert.Nil(t, err)
	assert.Equal(t, "OK", resp)

	resp, err = rw.Execute("RWSET transactions 100")
	assert.Nil(t, err)
	assert.Equal(t, "NACK", resp)

	resp, err = rw.Execute("RWADD transactions 600.60")
	assert.Nil(t, err)
	assert.Equal(t, "OK", resp)

	resp, err = rw.Execute("RWADD transactions notANumber")
	assert.Equal(t, err.Error(), "value must be a valid float")
	assert.Equal(t, "ERROR", resp)

	resp, err = rw.Execute("RWADD other 10.30")
	assert.NotNil(t, err, "key not found")
	assert.Equal(t, "", resp)

	resp, err = rw.Execute("RWAVG transactions")
	assert.Nil(t, err)
	assert.Equal(t, "10.01", resp)

	resp, err = rw.Execute("RWDEL transactions")
	assert.Nil(t, err)
	assert.Equal(t, "OK", resp)

	resp, err = rw.Execute("RWDEL transactions")
	assert.NotNil(t, err, "key not found")
	assert.Equal(t, "", resp)

	resp, err = rw.Execute("")
	assert.Equal(t, protocol.ErrMalformedCommand, err)
	assert.Equal(t, "", resp)
}
