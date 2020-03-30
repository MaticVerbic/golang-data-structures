# golang-data-structures

## Singly linked list
Implementation of simple singly linked list.

### Example
```go
l := singlylinkedlist.New(false)
l.Append(5)
l.Append(6)
l.Append(7)
l.Append(8)
l.Append(9)
l.Prepend(4)
l.Prepend(2)
l.Prepend(1)
l.Prepend(0)
l.InsertAt(3, 3)

for elem := l.Iter(); elem != nil; elem = elem.Next() {
  fmt.Println(elem)
}

fmt.Println(l)
```
### Documentation
```go

// Returns a new List struct.
// If argument circular is true, the list behaves as a circular singly linked list.
New(circular bool) *List

// Add a new node to the head
(l *List) Append(data interface{})

// Add a new node to the tail
(l *List) Prepend(data interface{})

// Add a new node at specified index.
// Returns an out of bounds error if index is invalid.
(l *List) InsertAt(index int, data interface{}) error

// Remove a node from list.
// Returns an out of bounds error if index is invalid.
(l *List) DeleteAt(index int) error

// Returns the first element of the list.
(l *List) Iter() *Node

// String representation of the list.
(l *List) String() string

// Returns data from the node.
(n *Node) Data() interface{}

// Returns the next node in the list.
// Returns nil if element is the head.
(n *Node) Next() *Node

// Returns a string representation of the node.
(n *Node) String() string
```