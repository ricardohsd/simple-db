package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

func TestStatistics(t *testing.T) {
	s := &statistics{
		data: make(map[string]int),
	}

	g := &errgroup.Group{}

	g.Go(func() error {
		for i := 0; i < 10; i++ {
			s.Incr("GET")
		}
		return nil
	})

	g.Go(func() error {
		for i := 0; i < 5; i++ {
			s.Incr("SET")
		}
		return nil
	})

	g.Go(func() error {
		for i := 0; i < 7; i++ {
			s.Incr("DEL")
		}
		return nil
	})

	g.Wait()

	assert.Equal(t, map[string]int{
		"GET": 10,
		"SET": 5,
		"DEL": 7,
	}, s.GetAll())
}
