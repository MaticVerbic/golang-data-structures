package queue

import "fmt"

// Queue is a simple implementation of a queue.
type Queue struct {
	head *Node
	tail *Node
	size int
}

// Node represents a node in a Queue.
type Node struct {
	data interface{}
	next *Node
}

// New returns a point to a new Queue.
func New() *Queue {
	return &Queue{}
}

// Enqueue a new node at the end of the queue.
func (q *Queue) Enqueue(data interface{}) {
	if q.size == 0 {
		q.head = &Node{
			data: data,
		}
		q.tail = q.head
		q.size++
		return
	}

	q.tail.next = &Node{
		data: data,
	}

	q.tail = q.tail.next
	q.size++
}

// Dequeue an item from the front of the queue.
func (q *Queue) Dequeue() interface{} {
	switch q.size {
	case 0:
		return nil
	case 1:
		n := q.tail
		q.head = nil
		q.tail = nil
		q.size--
		return n.data
	case 2:
		n := q.head
		q.head = q.tail
		q.size--
		return n.data
	}

	n := q.head
	q.head = q.head.next
	q.size--
	return n.data
}

// Implements stringify interface.
func (q *Queue) String() string {
	s := ""
	i := 0
	for node := q.head; node != nil; node = node.next {
		if i != 0 {
			s += ", "
		}

		s += fmt.Sprintf("%s", node)
		i++
	}

	return fmt.Sprintf("[%s)", s)
}

// Implements stringify interface.
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.data)
}
