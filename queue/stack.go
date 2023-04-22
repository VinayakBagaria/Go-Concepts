package queue

import "fmt"

type queue struct {
	s1 []int
	s2 []int
}

func newQueue() queue {
	return queue{s1: []int{}, s2: []int{}}
}

func (q *queue) push(x int) {
	for _, element := range q.s1 {
		q.s2 = append(q.s2, element)
	}

	q.s1 = append([]int{}, x)
	for _, element := range q.s2 {
		q.s1 = append(q.s1, element)
	}

	q.s2 = []int{}
}

func (q *queue) pop() int {
	lastElement := q.peek()
	q.s1 = q.s1[:len(q.s1) - 1]
	return lastElement
}

func (q *queue) peek() int {
	return q.s1[len(q.s1) - 1]
}

func (q *queue) isEmpty() bool {
	return len(q.s1) == 0
}

func usingStack() {
	q := newQueue()
	q.push(1)
	q.push(2)
	fmt.Printf("Stack ---- %v\n", q.s1)
	fmt.Printf("Peek ---- %d\n", q.peek())
	fmt.Printf("Pop ---- %d\n", q.pop())
	fmt.Printf("Empty ---- %t\n", q.isEmpty())
	fmt.Printf("Peek ---- %d\n", q.peek())
	fmt.Printf("Pop ---- %d\n", q.pop())
	fmt.Printf("Empty ---- %t\n", q.isEmpty())
}