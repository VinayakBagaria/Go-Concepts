// https://dev.to/douglasmakey/understanding-unix-domain-sockets-in-golang-32n8
package httpsocket

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const socketPath = "/tmp/httpecho.sock"

func DoWork() {
	socket, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal(err)
	}

	// cleanup
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGTSTP)
	go func() {
		<-c
		os.Remove(socketPath)
		os.Exit(1)
	}()

	m := http.NewServeMux()
	m.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health check passed"))
	})

	server := http.Server{Handler: m}
	if err := server.Serve(socket); err != nil {
		log.Fatal(err)
	}
}

// curl --unix-socket /tmp/httpecho.sock http://localhost/

// connection over unix socket
// useful for local inter-process communication, no n/w-ing overhead involved, faster than TCP n/w,
// more secure as same machine only, local unix permissions of user & group
// like b/w gunicorn and nginx
