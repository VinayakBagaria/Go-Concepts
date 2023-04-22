package main

import (
	"math/rand"
	"time"
	"tryouts/designpatterns"
	"tryouts/linkedlist"
	"tryouts/queue"
	"tryouts/sorting"
	"tryouts/statemachine"
	"tryouts/synchronization"
)

func generateSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99)
	}
	return slice
}

const decision = "state_machine"

func main() {
	switch decision {
	case "sorting":
		sorting.DoWork()
	case "synchronization":
		synchronization.DoWork()
	case "linkedlist":
		linkedlist.DoWork()
	case "queue":
		queue.DoWork()
	case "design_patterns":
		designpatterns.DoWork()
	case "state_machine":
		statemachine.DoWork()
	}
}
