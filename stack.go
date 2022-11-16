package datastruct

import "sync"

// NewStack creates a new instance of the stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Stack implements a LIFO structure.
type Stack[T any] struct {
	mux   sync.Mutex
	first *stackItem[T]
}

// Push value into the stack.
func (s *Stack[T]) Push(val T) {
	s.mux.Lock()
	defer s.mux.Unlock()
	first := s.first
	s.first = &stackItem[T]{val: val, prev: first}
}

// Pop value from the stack. It additionally returns a boolean flag that indicates if the stack was not empty.
func (s *Stack[T]) Pop() (T, bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	var val T
	if s.first == nil {
		return val, false
	}
	val = s.first.val
	s.first = s.first.prev
	return val, true
}

// stackItem is a single item in the stack that contains a value and a pointer to the prev value.
type stackItem[T any] struct {
	val  T
	prev *stackItem[T]
}
