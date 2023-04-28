package main

import (
	"go-concepts/consul"
	"go-concepts/designpatterns"
	"go-concepts/grpcsystem/grpcclient"
	"go-concepts/grpcsystem/grpcserver"
	"go-concepts/linkedlist"
	"go-concepts/loadbalancer"
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

const decision = "consul"

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
	case "load_balancer":
		loadbalancer.DoWork()
	case "grpc":
		go grpcserver.DoWork()
		grpcclient.DoWork()
	case "consul":
		consul.DoWork()
	}
}
