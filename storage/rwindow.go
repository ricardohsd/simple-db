package storage

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/ricardohsd/horus"
	"github.com/ricardohsd/simple-db/protocol"
)

type rollingStatistics interface {
	Add(value float64)
	Average() float64
}

type rwindow struct {
	db       map[string]rollingStatistics
	protocol protocol.Protocol
}

// NewRWindow returns a storage that allows to create multiple rolling window metrics using Horus
func NewRWindow() *rwindow {
	p := &protocol.RWindow{}

	return &rwindow{
		db:       make(map[string]rollingStatistics),
		protocol: p,
	}
}

func (r *rwindow) Execute(message string) (string, error) {
	cmd, err := r.protocol.Parse(message)
	if err != nil {
		return "", err
	}

	switch cmd.Instruction {
	case "RWSET":
		return r.set(cmd)
	case "RWADD":
		return r.add(cmd)
	case "RWAVG":
		return r.average(cmd)
	case "RWDEL":
		return r.del(cmd)
	default:
		return "", errors.Wrap(ErrWrongInstruction, cmd.Instruction)
	}
}

func (r *rwindow) set(cmd *protocol.Command) (string, error) {
	log.Printf("Processing Set %v\n", cmd.Key)

	_, ok := r.db[cmd.Key]
	if ok {
		return "NACK", nil
	}

	seconds, err := strconv.Atoi(cmd.Value)
	if err != nil {
		return "ERROR", errors.New("value must be a valid integer")
	}

	rw, err := horus.NewRWindow(time.Duration(seconds)*time.Second, time.Second)
	if err != nil {
		return "", err
	}

	r.db[cmd.Key] = rw

	return "OK", nil
}

func (r *rwindow) add(cmd *protocol.Command) (string, error) {
	log.Printf("Processing Add %v\n", cmd.Key)

	rw, ok := r.db[cmd.Key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	val, err := strconv.ParseFloat(cmd.Value, 64)
	if err != nil {
		return "ERROR", errors.New("value must be a valid float")
	}

	rw.Add(val)

	return "OK", nil
}

func (r *rwindow) average(cmd *protocol.Command) (string, error) {
	log.Printf("Processing Average %v\n", cmd.Key)

	rw, ok := r.db[cmd.Key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	avg := rw.Average()

	return fmt.Sprintf("%.2f", avg), nil
}

func (r *rwindow) del(cmd *protocol.Command) (string, error) {
	log.Printf("Processing DEL %v\n", cmd.Key)

	_, ok := r.db[cmd.Key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	delete(r.db, cmd.Key)

	return "OK", nil
}
