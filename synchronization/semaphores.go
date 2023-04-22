package synchronization

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

const (
	iterations = 10
	goroutines = 2
)

var sm = semaphore.NewWeighted(goroutines)

func SempahoreImplementation() {
	ctx := context.Background()

	for i := 0; i < iterations; i++ {
		sm.Acquire(ctx, 1)
		go func() {
			defer sm.Release(1)
			time.Sleep(time.Second)
		}()
	}

	// print the number of running goroutines without the main goroutine
	fmt.Printf("currently running: %d goroutines\n", runtime.NumGoroutine()-1)
}
