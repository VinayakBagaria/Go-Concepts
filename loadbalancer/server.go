package loadbalancer

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

type server struct {
	url    string
	proxy  *httputil.ReverseProxy
	health bool
	mux    sync.RWMutex
}

type serverPool struct {
	servers []*server
	current uint64
}

func createServers(desiredCount int) []*server {
	serverList := []*server{}
	for i := 0; i < desiredCount; i++ {
		serverUrl := fmt.Sprintf("http://localhost:500%d", i)
		serverList = append(serverList, newServer(serverUrl))
	}

	return serverList
}

func newServer(rawURL string) *server {
	urlStruct, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	return &server{
		url:    rawURL,
		proxy:  httputil.NewSingleHostReverseProxy(urlStruct),
		health: true,
	}
}

func (s *server) checkHealth() {
	resp, err := http.Head(s.url)
	if err != nil {
		s.setHealthy(false)
		return
	}

	s.setHealthy(resp.StatusCode == http.StatusOK)
}

func (s *server) isHealthy() bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.health
}

func (s *server) setHealthy(health bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.health = health
}

func (s *serverPool) getNext() *server {
	nextIndex := s.getNextIndex()
	l := nextIndex + len(s.servers)

	for i := nextIndex; i < l; i++ {
		idx := i % len(s.servers)
		server := s.servers[idx]
		if server.isHealthy() {
			atomic.StoreUint64(&s.current, uint64(idx))
			return server
		}
	}

	return nil
}

func (s *serverPool) getNextIndex() int {
	nextIndex := atomic.AddUint64(&s.current, uint64(1))
	return int(nextIndex % uint64(len(s.servers)))
}
