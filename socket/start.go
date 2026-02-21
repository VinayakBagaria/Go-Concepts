package socket

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func DoWork() {
	// Create a Unix domain socket and listen for incoming connections
	socket, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal(err)
	}

	// Cleanup sock file
	c := make(chan os.Signal, 1)
	// SIGINT -> ctrl + c
	// SIGTERM -> kill
	// SIGTSTP -> ctrl + z - terminal stop
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGTSTP)
	go func() {
		<-c
		log.Println()
		log.Println("exit")
		os.Remove("/tmp/echo.sock")
		os.Exit(1)
	}()

	for {
		// Accept an incoming connection
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Handle the connection in a separate goroutine
		go func(conn net.Conn) {
			// Create a buffer for incoming data
			buf := make([]byte, 4096)

			// Read data from the connection
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}

			// Echo the data back to the connection
			_, err = conn.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}(conn)
	}
}

// echo "Vinayak" | /usr/bin/nc -U /tmp/echo.sock
