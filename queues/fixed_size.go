package queues

import (
	"github.com/sonirico/container/types"
)

type queue[T types.Any] interface {
	Push(value T)
	Poll() (T, bool)
	Size() int
}

// FixedSizeQueue ensures the inner queue maintains an upper limit of capacity
type FixedSizeQueue[T types.Any] struct {
	capacity int

	queue queue[T]
}

func (q *FixedSizeQueue[T]) Push(value T) {
	if q.queue.Size() == q.capacity {
		q.Poll()
	}

	q.queue.Push(value)
}

func (q *FixedSizeQueue[T]) PushEvict(value T) (item T, evicted bool) {
	if q.queue.Size() == q.capacity {
		item, evicted = q.Poll()
	}

	q.queue.Push(value)
	return
}

func (q *FixedSizeQueue[T]) Poll() (T, bool) {
	return q.queue.Poll()
}

func (q *FixedSizeQueue[T]) Size() int {
	return q.queue.Size()
}

func NewFixedSizeQueue[T types.Any](capacity int, inner queue[T]) FixedSizeQueue[T] {
	return FixedSizeQueue[T]{queue: inner, capacity: capacity}
}
