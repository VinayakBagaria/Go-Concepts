package loadbalancer

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type server struct {
	name   string
	url    string
	proxy  *httputil.ReverseProxy
	health bool
}

var (
	serverList      []*server
	lastServedIndex int
)

func createServers(desiredCount int) {
	for i := 0; i < desiredCount; i++ {
		serverName := fmt.Sprintf("server-%d", i)
		serverUrl := fmt.Sprintf("http://localhost:500%d", i)
		serverList = append(serverList, newServer(serverName, serverUrl))
	}
}

func newServer(name, rawURL string) *server {
	urlStruct, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	return &server{
		name:   name,
		url:    rawURL,
		proxy:  httputil.NewSingleHostReverseProxy(urlStruct),
		health: true,
	}
}

func (s *server) checkHealth() {
	resp, err := http.Head(s.url)
	if err != nil {
		s.health = false
		return
	}

	s.health = resp.StatusCode == http.StatusOK
}

func getHealthyServer() (*server, error) {
	for i := 0; i < len(serverList); i++ {
		s := getServer()
		if s.health {
			return s, nil
		}
	}

	return nil, fmt.Errorf("no healthy server")
}

func getServer() *server {
	s := serverList[lastServedIndex]
	lastServedIndex = (lastServedIndex + 1) % len(serverList)
	return s
}
