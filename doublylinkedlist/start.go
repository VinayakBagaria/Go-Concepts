package doublylinkedlist

type Node struct {
	Value      any
	Next, Prev *Node
}

type List struct {
	head, tail *Node
	length     int
}

func (l *List) PushFront(node *Node) {
	if l.head != nil {
		node.Next = l.head
		l.head.Prev = node
	}
	l.head = node

	if l.tail == nil {
		l.tail = node
	}

	l.length += 1
}

func (l *List) Remove(node *Node) {
	if node == nil {
		return
	}

	prevItem := node.Prev
	nextItem := node.Next

	if prevItem != nil {
		prevItem.Next = nextItem
	}

	if nextItem != nil {
		nextItem.Prev = prevItem
	}

	if l.head == node {
		l.head = nextItem
	}

	if l.tail == node {
		l.tail = prevItem
	}

	node.Prev = nil
	node.Next = nil

	l.length -= 1
}

func (l *List) MoveToFront(node *Node) {
	l.Remove(node)
	l.PushFront(node)
}

func (l *List) Front() *Node {
	return l.head
}

func (l *List) Back() *Node {
	return l.tail
}

func (l *List) Length() int {
	return l.length
}
