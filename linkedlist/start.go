package linkedlist

import (
	"fmt"
)

type node struct {
	data string
	next *node
}

type linkedlist struct {
	head *node
}

func newLinkedlist() *linkedlist {
	return &linkedlist{}
}

func (l *linkedlist) addToFirst(data string) {
	newNode := &node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
}

func (l *linkedlist) addToLast(data string) {
	newNode := &node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

func (l *linkedlist) print() {
	if l.head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%s -> ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Printf("null\n")
}

func (l *linkedlist) deleteFirst() {
	if l.head == nil {
		return
	}

	l.head = l.head.next
}

func (l *linkedlist) deleteLast() {
	if l.head == nil {
		return
	}

	if l.head.next == nil {
		l.head = nil
		return
	}

	secondLast := l.head
	lastNode := l.head.next

	for lastNode.next != nil {
		secondLast = secondLast.next
		lastNode = lastNode.next
	}

	secondLast.next = nil
}

func (l *linkedlist) reverse() {
	if l.head == nil {
		return
	}

	previous := l.head
	current := l.head.next

	for current != nil {
		next := current.next
		current.next = previous

		previous = current
		current = next
	}
	l.head.next = nil
	l.head = previous
}

func reverseRecursive(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	newHead := reverseRecursive(head.next)
	head.next.next = head
	head.next = nil
	return newHead
}

func DoWork() {
	ll := newLinkedlist()
	ll.addToFirst("a")
	ll.addToFirst("is")
	ll.addToLast("list")
	ll.addToFirst("this")
	ll.print()
	fmt.Println("delete first ----")
	ll.deleteFirst()
	ll.print()
	fmt.Println("delete last ----")
	ll.deleteLast()
	ll.print()
	fmt.Println("insert again ----")
	ll.addToLast("list")
	ll.addToFirst("this")
	ll.print()
	fmt.Println("Reverse ----")
	ll.reverse()
	ll.print()
	fmt.Println("Reverse Recursive ----")
	ll.head = reverseRecursive(ll.head)
	ll.print()
}
