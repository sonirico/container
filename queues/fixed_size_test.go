package queues

import (
	"reflect"
	"testing"
)

type testFixedSizeQueueAction struct {
	opcode       op
	value        int
	bsize        int
	asize        int
	shouldEvict  bool
	evictedValue int
}

func TestFixedSizeQueue(t *testing.T) {
	tests := []struct {
		name     string
		sequence []testFixedSizeQueueAction
	}{
		{
			name: "poll when empty",
			sequence: []testFixedSizeQueueAction{
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
			sequence: []testFixedSizeQueueAction{
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
			sequence: []testFixedSizeQueueAction{
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
			sequence: []testFixedSizeQueueAction{
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
					opcode:       push,
					value:        3,
					bsize:        2,
					asize:        2,
					shouldEvict:  true,
					evictedValue: 1,
				},
				{
					opcode: poll,
					value:  2,
					bsize:  2,
					asize:  1,
				},
				{
					opcode: push,
					value:  4,
					bsize:  1,
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
			capacity := 2
			fifo := &FifoQueue[int]{}
			q := NewFixedSizeQueue[int](capacity, fifo)

			for _, seq := range test.sequence {
				if q.Size() != seq.bsize {
					t.Errorf("unexpected before size, want %d, have %d",
						seq.bsize, q.Size())
				}

				switch seq.opcode {
				case push:
					item, evicted := q.PushEvict(seq.value)
					if seq.shouldEvict != evicted {
						t.Errorf("unexpected evicted event, want %t, have %t",
							seq.shouldEvict, evicted)
					}
					if !reflect.DeepEqual(item, seq.evictedValue) {
						t.Errorf("unexpected evicted value, want %v, have %v",
							seq.evictedValue, item)
					}
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
