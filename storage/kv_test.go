package storage

import (
	"testing"

	"github.com/ricardohsd/simple-db/protocol"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	kv := &kv{
		db:       make(map[string]string),
		protocol: &protocol.KV{},
	}

	resp, err := kv.Execute("GET name")
	assert.NotNil(t, err, "key not found")
	assert.Equal(t, "", resp)

	resp, err = kv.Execute("SET name john")
	assert.Nil(t, err)
	assert.Equal(t, "OK", resp)

	resp, err = kv.Execute("GET name")
	assert.Nil(t, err)
	assert.Equal(t, "john", resp)

	resp, err = kv.Execute("DEL name")
	assert.Nil(t, err)
	assert.Equal(t, "OK", resp)

	resp, err = kv.Execute("DEL name")
	assert.NotNil(t, err, "key not found")
	assert.Equal(t, "", resp)

	resp, err = kv.Execute("")
	assert.Equal(t, protocol.ErrMalformedCommand, err)
	assert.Equal(t, "", resp)
}
