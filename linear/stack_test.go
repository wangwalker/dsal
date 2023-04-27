package linear

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()
	if stack.Top() != nil {
		t.Error("Top() should return nil")
	}
	if stack.Pop() != nil {
		t.Error("Pop() should return nil")
	}
	if !stack.Empty() {
		t.Error("Empty() should return true")
	}
	stack.Push(1)
	if stack.Top() != 1 {
		t.Error("Top() should return 1")
	}
	stack.Push(2)
	if stack.Top() != 2 {
		t.Error("Top() should return 2")
	}
	if stack.Pop() != 2 {
		t.Error("Pop() should return 2")
	}
	if stack.Top() != 1 {
		t.Error("Top() should return 1")
	}
	if stack.Pop() != 1 {
		t.Error("Pop() should return 1")
	}
	if stack.Top() != nil {
		t.Error("Top() should return nil")
	}
	if stack.Pop() != nil {
		t.Error("Pop() should return nil")
	}
	if !stack.Empty() {
		t.Error("Empty() should return true")
	}
}
