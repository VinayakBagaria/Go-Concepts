package synchronization

import (
	"fmt"
	"sync"
)

var (
	n       = 1000
	counter = 0
	wg      = &sync.WaitGroup{}
	mutex   = &sync.Mutex{}
)

func MutexImplementation() {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer mutex.Unlock()
			defer wg.Done()
			mutex.Lock()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
