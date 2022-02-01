package queues

import "github.com/sonirico/container/types"

// spnode represents a single pointer queue node with just one pointer: to the next element
type spnode[T types.Any] struct {
	value T

	next *spnode[T]
}
