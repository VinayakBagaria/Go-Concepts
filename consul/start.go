package consul

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/hashicorp/consul/api"
)

const (
	ttl     = time.Second * 5
	checkId = "check_health"
)

type Service struct {
	consulClient *api.Client
}

func NewService() *Service {
	client, err := api.NewClient(&api.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &Service{consulClient: client}
}

func (s *Service) registerService() {
	check := &api.AgentServiceCheck{
		DeregisterCriticalServiceAfter: ttl.String(),
		TLSSkipVerify:                  true,
		TTL:                            ttl.String(),
		CheckID:                        checkId,
	}

	reg := &api.AgentServiceRegistration{
		ID:      "auth_service",
		Name:    "dev_cluster",
		Tags:    []string{"auth"},
		Address: "127.0.0.1",
		Port:    3000,
		Check:   check,
	}

	if err := s.consulClient.Agent().ServiceRegister(reg); err != nil {
		log.Fatal(err)
	}
}

func (s *Service) Start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	s.registerService()
	go s.updateHealthCheck()
	s.acceptLoop(ln)
}

func (s *Service) acceptLoop(ln net.Listener) {
	for {
		_, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *Service) updateHealthCheck() {
	ticker := time.NewTicker(time.Second * 2)
	for {
		if err := s.consulClient.Agent().UpdateTTL(checkId, "online", api.HealthPassing); err != nil {
			log.Fatal(err)
		}
		<-ticker.C
	}
}

func DoWork() {
	s := NewService()
	fmt.Println("Starting server...")
	s.Start()
}
