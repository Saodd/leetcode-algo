package q0232

import (
	"testing"
)

func TestNewMyQueue(t *testing.T) {
	q := NewMyQueue()
	q.Push(1)
	if x := q.Peek(); x != 1 {
		t.Errorf("Peek failed. %d, %d", x, 1)
	}
	if x := q.Pop(); x != 1 {
		t.Errorf("Pop failed. %d, %d", x, 1)
	}
	if !q.Empty() {
		t.Errorf("Empty failed.")
	}
}
