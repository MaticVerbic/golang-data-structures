package stack

import "fmt"

// Stack is a simple implementation of a stack.
type Stack struct {
	top  *Node
	size int
}

// Node represents a stack node.
type Node struct {
	data interface{}
	prev *Node
}

// New returns a pointer to a new Stack.
func New() *Stack {
	return &Stack{}
}

// Push to the top of the stack.
func (s *Stack) Push(data interface{}) {
	s.top = &Node{
		data: data,
		prev: s.top,
	}

	s.size++
}

// Pop from the top of the stack.
func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}

	n := s.top
	s.top = s.top.prev

	s.size--
	return n.data
}

// Peek returns top item without removing it from stack.
func (s *Stack) Peek() interface{} {
	if s.top == nil {
		return nil
	}

	return s.top.data
}

// Implements stringify interface.
func (s *Stack) String() string {
	str := ""
	i := 0

	for node := s.top; node != nil; node = node.prev {
		if i != 0 {
			str += ", "
		}

		str += fmt.Sprintf("%s", node)
		i++
	}
	return fmt.Sprintf("[%s)", str)
}

// Size returns a stack size.
func (s *Stack) Size() int {
	return s.size
}

// Data returns data.
func (n *Node) Data() interface{} {
	return n.data
}

// Implements stringify interface.
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.data)
}
