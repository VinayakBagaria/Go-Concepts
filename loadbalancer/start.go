package loadbalancer

import (
	"fmt"
	"log"
	"net/http"
)

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	s, err := getHealthyServer()
	if err != nil {
		http.Error(w, "Couldn't process request: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("X-Proxy", "golang-proxy")
	w.Header().Set("X-Origin", s.url)
	s.proxy.ServeHTTP(w, r)
}

func DoWork() {
	createServers(5)
	startHealthCheck()
	http.HandleFunc("/", forwardRequest)
	fmt.Println("Started server")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
