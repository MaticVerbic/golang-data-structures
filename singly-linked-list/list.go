package singlylinkedlist

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
	sentinel bool
}

// New returns a new Singly linked list.
// Implementational choice here was to use sentinel node as terminatory element.
func New(circular bool) *List {
	return &List{
		first: &Node{
			sentinel: true,
		},
		size:     0,
		circular: circular,
	}
}

// Append appends an element to end of list.
func (l *List) Append(data interface{}) {
	for node := l.first; node != nil; node = node.Next() {
		if node.next == nil || (l.circular && node.next.sentinel) {
			node.next = &Node{
				data: data,
			}

			if l.circular {
				node.next.next = l.first
			}

			l.size++
			break
		}
	}
}

// Prepend inserts an element at the beginning of the linked list.
func (l *List) Prepend(data interface{}) {
	if l.size == 0 {
		l.Append(data)
		return
	}

	node := &Node{
		data: data,
	}

	node.next = l.first.next
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
				if l.circular {
					elem.next = l.first
				} else {
					elem.next = nil
				}
			}
			i++
		}
	default:
		i := 0
		for elem := l.Iter(); elem != nil; elem = elem.Next() {
			if i == index-1 {
				elem.next = elem.next.next
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

	if n.next.sentinel {
		return n.next.next
	}

	return n.next
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.data)
}
