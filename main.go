package main

import (
	"go-concepts/designpatterns"
	"go-concepts/linkedlist"
	"go-concepts/lrucache"
	"go-concepts/pipefilter"
	"go-concepts/queue"
	"go-concepts/sorting"
	"go-concepts/statemachine"
	"go-concepts/synchronization"
	"math/rand"
	"time"
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
