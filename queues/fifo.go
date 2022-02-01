package queues

import "github.com/sonirico/container/types"

type FifoQueue[T types.Any] struct {
	size int

	head *spnode[T]
	tail *spnode[T]
}

func (q *FifoQueue[T]) Size() int {
	return q.size
}

func (q *FifoQueue[T]) Push(value T) {
	node := &spnode[T]{value: value}

	if q.size == 0 {
		q.head = node
		q.tail = node
		q.size = 1
		return
	}

	if q.size == 1 {
		q.tail = node
		q.tail.next = nil
		q.head.next = q.tail
		q.size = 2
		return
	}

	q.tail.next = node
	q.tail = node
	q.size++
}

func (q *FifoQueue[T]) Poll() (T, bool) {
	if q.size == 0 {
		var noop T
		return noop, false
	}

	if q.size == 1 {
		v := q.head.value
		q.head.next = nil
		q.tail.next = nil
		q.head = nil
		q.tail = nil
		q.size--

		return v, true
	}

	q.size--
	v := q.head.value
	q.head = q.head.next
	return v, true
}

func NewFifoQueue[T types.Any]() FifoQueue[T] {
	return FifoQueue[T]{}
}
