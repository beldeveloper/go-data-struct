package datastruct

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	val, exists := q.Pop()
	if exists {
		t.Fatalf("expected=empty; got=notEmpty")
	}
	n := 3
	for i := 0; i < n; i++ {
		q.Push(i)
	}
	for i := 0; i < n; i++ {
		val, exists = q.Pop()
		if !exists {
			t.Fatalf("expected=notEmpty; got=empty")
		}
		if val != i {
			t.Fatalf("expected=%d; got=%d", i, val)
		}
	}
	val, exists = q.Pop()
	if exists {
		t.Fatalf("expected=empty; got=notEmpty")
	}
}
