package linear

// Stack is a linear data structure which follows a particular order in which
// the operations are performed. The order may be LIFO(Last In First Out) or
// FILO(First In Last Out).
type Stack struct {
	top   interface{}
	items []interface{}
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{top: nil, items: make([]interface{}, 0)}
}

// Push pushes an item onto the top of this stack.
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
	s.top = item
}

// Pop removes the object at the top of this stack and returns that object as
// the value of this function. If the stack is empty, the function returns nil.
func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.top
	s.items = s.items[:len(s.items)-1]
	if len(s.items) == 0 {
		s.top = nil
	} else {
		s.top = s.items[len(s.items)-1]
	}
	return item
}

// Top returns the object at the top of this stack without removing it from the
// stack. If the stack is empty, the function returns nil.
func (s *Stack) Top() interface{} {
	return s.top
}

// Empty returns true if this stack contains no elements.
func (s *Stack) Empty() bool {
	return len(s.items) == 0
}
