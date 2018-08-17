package proxy

import "sync"

type statistics struct {
	sync.RWMutex
	data map[string]int
}

// NewStatistics returns a datastruct that provides basic statistics
// over client/server communication
func NewStatistics() *statistics {
	return &statistics{
		data: make(map[string]int),
	}
}

func (s *statistics) Incr(key string) {
	s.Lock()
	defer s.Unlock()

	var count int

	count, ok := s.data[key]
	if !ok {
		count = 0
	}

	s.data[key] = count + 1
}

func (s *statistics) GetAll() map[string]int {
	s.RLock()
	defer s.RUnlock()

	return s.data
}
