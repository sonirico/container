package queues

type op int

const (
	push op = iota
	poll
)

type testFifoQueueAction struct {
	opcode op
	value  int
	bsize  int
	asize  int
}
