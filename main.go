package main

import (
	"fmt"
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
	"os"
	"strings"
)

func main() {
	concepts := map[string]func(){
		"synchronization": synchronization.DoWork,
		"design_patterns": designpatterns.DoWork,
		"state_machine":   statemachine.DoWork,
		"pipe_filter":     pipefilter.DoWork,
		"lru_cache":       lrucache.DoWork,
		"load_balancer":   loadbalancer.DoWork,
		"grpc": func() {
			go grpcserver.DoWork()
			grpcclient.DoWork()
		},
		"consul":            consul.DoWork,
		"bloom_filters":     bloomfilters.DoWork,
		"medium_remover":    mediumremover.DoWork,
		"option_pattern":    optionpattern.DoWork,
		"middleware":        middleware.DoWork,
		"token_bucket":      tokenbucket.DoWork,
		"context":           contextual.DoWork,
		"internals":         internals.DoWork,
		"signals":           signals.DoWork,
		"mapwithexpiration": mapwithexpiration.DoWork,
		"threadpool":        threadpool.DoWork,
		"workers":           workers.DoWork,
		"fanIn":             fanin.DoWork,
		"tree":              tree.DoWork,
	}

	var conceptNames []string
	for conceptName := range concepts {
		conceptNames = append(conceptNames, conceptName)
	}

	if len(os.Args) == 1 {
		fmt.Println("Choose a concept:")
		fmt.Println(strings.Join(conceptNames, ", "))
		return
	}

	decision := os.Args[1]
	concepts[decision]()
}
