package queues

import (
	"reflect"
	"testing"
)

func TestFifoQueue(t *testing.T) {
	tests := []struct {
		name     string
		sequence []testFifoQueueAction
	}{
		{
			name: "poll when empty",
			sequence: []testFifoQueueAction{
				{
					opcode: poll,
					value:  0, // zero value for generic
					bsize:  0,
					asize:  0,
				},
			},
		},
		{
			name: "push & poll",
			sequence: []testFifoQueueAction{
				{
					opcode: push,
					value:  1,
					bsize:  0,
					asize:  1,
				},
				{
					opcode: poll,
					value:  1,
					bsize:  1,
					asize:  0,
				},
			},
		},
		{
			name: "push & poll & poll",
			sequence: []testFifoQueueAction{
				{
					opcode: push,
					value:  1,
					bsize:  0,
					asize:  1,
				},
				{
					opcode: poll,
					value:  1,
					bsize:  1,
					asize:  0,
				},
				{
					opcode: poll,
					value:  0,
					bsize:  0,
					asize:  0,
				},
			},
		},
		{
			name: "back and forth",
			sequence: []testFifoQueueAction{
				{
					opcode: push,
					value:  1,
					bsize:  0,
					asize:  1,
				},
				{
					opcode: push,
					value:  2,
					bsize:  1,
					asize:  2,
				},
				{
					opcode: push,
					value:  3,
					bsize:  2,
					asize:  3,
				},
				{
					opcode: poll,
					value:  1,
					bsize:  3,
					asize:  2,
				},
				{
					opcode: push,
					value:  4,
					bsize:  2,
					asize:  3,
				},
				{
					opcode: poll,
					value:  2,
					bsize:  3,
					asize:  2,
				},
				{
					opcode: poll,
					value:  3,
					bsize:  2,
					asize:  1,
				},
				{
					opcode: poll,
					value:  4,
					bsize:  1,
					asize:  0,
				},
				{
					opcode: poll,
					value:  0,
					bsize:  0,
					asize:  0,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := NewFifoQueue[int]()

			for _, seq := range test.sequence {
				if q.Size() != seq.bsize {
					t.Errorf("unexpected before size, want %d, have %d",
						seq.bsize, q.Size())
				}
				switch seq.opcode {
				case push:
					q.Push(seq.value)
				case poll:
					actual, _ := q.Poll()
					if !reflect.DeepEqual(actual, seq.value) {
						t.Fatalf("unexpected poll value, want %v, have %v, ctx (%v)",
							seq.value, actual, seq)
					}
				}

				if q.Size() != seq.asize {
					t.Errorf("unexpected after size, want %d, have %d",
						seq.asize, q.Size())
				}
			}
		})
	}
}
