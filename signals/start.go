package signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func DoWork() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		fmt.Printf("Received signal: %s\n", sig)
		close(quit)
	}
}
