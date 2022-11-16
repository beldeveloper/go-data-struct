package datastruct

import "sync"

// NewQueue creates a new instance of the queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Queue implements a singly linked list (FIFO structure).
type Queue[T any] struct {
	mux   sync.Mutex
	first *queueItem[T]
}

// Push value into the queue.
func (q *Queue[T]) Push(val T) {
	q.mux.Lock()
	defer q.mux.Unlock()
	last := &q.first
	for *last != nil {
		last = &(*last).next
	}
	*last = &queueItem[T]{val: val}
}

// Pop value from the queue. It additionally returns a boolean flag that indicates if the queue was not empty.
func (q *Queue[T]) Pop() (T, bool) {
	q.mux.Lock()
	defer q.mux.Unlock()
	var val T
	if q.first == nil {
		return val, false
	}
	val = q.first.val
	q.first = q.first.next
	return val, true
}

// queueItem is a single item in the queue that contains a value and a pointer to the next value.
type queueItem[T any] struct {
	val  T
	next *queueItem[T]
}
