package main

import (
	"math/rand"
	"time"
	"tryouts/designpatterns"
	"tryouts/linkedlist"
	"tryouts/lrucache"
	"tryouts/pipefilter"
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

const decision = "lru_cache"

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
	case "pipe_filter":
		pipefilter.DoWork()
	case "lru_cache":
		lrucache.DoWork()
	}
}
