package doublylinkedlist

import "fmt"

// List represents a singly linked list.
type List struct {
	first    *Node
	size     int
	circular bool
}

// Node represents a singly linked node.
type Node struct {
	data     interface{}
	next     *Node
	prev     *Node
	sentinel bool
	circular bool
}

// New returns a new Singly linked list.
// Implementational choice here was to use sentinel node as terminatory element.
func New(circular bool) *List {
	return &List{
		first: &Node{
			sentinel: true,
			circular: circular,
		},
		size:     0,
		circular: circular,
	}
}

// Append appends an element to end of list.
func (l *List) Append(data interface{}) {
	if l.size == 0 {
		l.first.next = &Node{
			data: data,
			prev: l.first,
			next: l.first,
		}

		l.first.prev = l.first.next
		l.size++
		return
	}

	node := &Node{
		data: data,
		next: l.first,
		prev: l.first.prev,
	}

	l.first.prev.next = node
	l.first.prev = node
	l.size++
}

// Prepend inserts an element at the beginning of the linked list.
func (l *List) Prepend(data interface{}) {
	if l.size == 0 {
		l.Append(data)
		return
	}

	node := &Node{
		data: data,
		prev: l.first,
		next: l.first.next,
	}

	l.first.next.prev = node
	l.first.next = node
	l.size++
}

// InsertAt inserts an element at desired index in list.
func (l *List) InsertAt(index int, data interface{}) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	switch index {
	case 0:
		l.Prepend(data)
		return nil
	case l.size:
		l.Append(data)
		return nil
	}

	node := &Node{
		data: data,
	}

	i := 0
	for elem := l.Iter(); elem != nil || (l.circular && elem.sentinel); elem = elem.next {
		if i == index-1 {
			node.next = elem.next
			node.prev = elem
			elem.next.prev = node
			elem.next = node
			break
		}
		i++
	}

	l.size++
	return nil
}

// DeleteAt removes an element at specified index.
func (l *List) DeleteAt(index int) error {
	if index < 0 || index > l.size-1 {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	switch index {
	case 0:
		l.first.next = l.first.next.next
	case l.size - 1:
		i := 0
		for elem := l.Iter(); i != index-1; elem = elem.next {
			if i == index-2 {
				elem.next = l.first
				l.first.prev = elem
			}
			i++
		}
	default:
		i := 0
		for elem := l.Iter(); elem != nil; elem = elem.Next() {
			if i == index-1 {
				elem.next = elem.next.next
				elem.next.prev = elem
				break
			}
			i++
		}
	}

	l.size--
	return nil
}

// Iter is a method to invoke iteration.
// If list is empty, nil is returned otherwise
// first node following the sentinel is returned.
func (l *List) Iter() *Node {
	if l.size == 0 {
		return nil
	}

	return l.first.next
}

// IterLast is a method to invoke iteration.
// If list is empty, nil is returned otherwise
// last node following the sentinel is returned.
func (l *List) IterLast() *Node {
	if l.size == 0 {
		return nil
	}

	return l.first.prev
}

func (l *List) String() string {
	s := ""
	if l.size > 0 {
		i := 0
		for node := l.Iter(); node != nil; node = node.next {
			if node.sentinel {
				break
			}
			if i != 0 {
				s += ", "
			}
			i++
			s += fmt.Sprintf("%s", node)
		}
	}

	return fmt.Sprintf("[%s]", s)
}

// Data returns data held by node.
func (n *Node) Data() interface{} {
	return n.data
}

// Next returns the next element in list.
func (n *Node) Next() *Node {
	if n.next == nil {
		return nil
	}

	if n.next.sentinel && n.next.circular {
		return n.next.next
	} else if n.next.sentinel {
		return nil
	}

	return n.next
}

// Prev returns the previous element in list.
func (n *Node) Prev() *Node {
	if n.prev == nil {
		return nil
	}

	if n.prev.sentinel && n.prev.circular {
		return n.prev.prev
	} else if n.prev.sentinel {
		return nil
	}

	return n.prev
}

func (n *Node) String() string {
	s := ""

	s += fmt.Sprintf("%v", n.data)

	return s
}
