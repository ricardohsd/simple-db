package protocol

// Command wraps instruction, key and value that
// will be executed in the storage
type Command struct {
	Instruction string
	Key         string
	Value       interface{}
}

// Protocol defines interface for protocol communication
type Protocol interface {
	Parse(message string) (*Command, error)
}
