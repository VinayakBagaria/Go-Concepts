package loadbalancer

import (
	"fmt"
	"log"
	"net/http"
)

var pool *serverPool

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	server := pool.getNext()
	if server == nil {
		http.Error(w, "Service not available", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("X-Proxy", "golang-proxy")
	w.Header().Set("X-Origin", server.url)
	server.proxy.ServeHTTP(w, r)
}

func DoWork() {
	serverList := createServers(5)
	pool = &serverPool{servers: serverList, current: 0}
	startHealthCheck()
	http.HandleFunc("/", forwardRequest)
	fmt.Println("Started server")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
