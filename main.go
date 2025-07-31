package main

import (
	"go-concepts/bloomfilters"
	"go-concepts/consul"
	"go-concepts/contextual"
	"go-concepts/designpatterns"
	"go-concepts/fanin"
	"go-concepts/grpcsystem/grpcclient"
	"go-concepts/grpcsystem/grpcserver"
	"go-concepts/internals"
	"go-concepts/loadbalancer"
	"go-concepts/lrucache"
	"go-concepts/mapwithexpiration"
	"go-concepts/mediumremover"
	"go-concepts/middleware"
	"go-concepts/optionpattern"
	"go-concepts/pipefilter"
	"go-concepts/signals"
	"go-concepts/statemachine"
	"go-concepts/synchronization"
	"go-concepts/threadpool"
	"go-concepts/tokenbucket"
	"go-concepts/tree"
	"go-concepts/workers"
)

const decision = "tree"

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
	case "token_bucket":
		tokenbucket.DoWork()
	case "context":
		contextual.DoWork()
	case "internals":
		internals.DoWork()
	case "signals":
		signals.DoWork()
	case "mapwithexpiration":
		mapwithexpiration.DoWork()
	case "threadpool":
		threadpool.DoWork()
	case "workers":
		workers.DoWork()
	case "fanIn":
		fanin.DoWork()
	case "tree":
		tree.DoWork()
	}
}
