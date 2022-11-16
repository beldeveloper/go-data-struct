package datastruct

import "testing"

func TestStack(t *testing.T) {
	s := NewStack[int]()
	val, exists := s.Pop()
	if exists {
		t.Fatalf("expected=empty; got=notEmpty")
	}
	n := 3
	for i := 0; i < n; i++ {
		s.Push(i)
	}
	for i := n - 1; i >= 0; i-- {
		val, exists = s.Pop()
		if !exists {
			t.Fatalf("expected=notEmpty; got=empty")
		}
		if val != i {
			t.Fatalf("expected=%d; got=%d", i, val)
		}
	}
	val, exists = s.Pop()
	if exists {
		t.Fatalf("expected=empty; got=notEmpty")
	}
}
