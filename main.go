package main

import (
	"go-concepts/bloomfilters"
	"go-concepts/consul"
	"go-concepts/designpatterns"
	"go-concepts/grpcsystem/grpcclient"
	"go-concepts/grpcsystem/grpcserver"
	"go-concepts/loadbalancer"
	"go-concepts/lrucache"
	"go-concepts/mediumremover"
	"go-concepts/middleware"
	"go-concepts/optionpattern"
	"go-concepts/pipefilter"
	"go-concepts/statemachine"
	"go-concepts/synchronization"
)

const decision = "middleware"

func main() {
	switch decision {
	case "synchronization":
		synchronization.DoWork()
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
	case "bloom_filters":
		bloomfilters.DoWork()
	case "medium_remover":
		mediumremover.DoWork()
	case "option_pattern":
		optionpattern.DoWork()
	case "middleware":
		middleware.DoWork()
	}
}
