package fanin

import (
	"fmt"
	"sync"
)

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()

			for input := range ch {
				out <- input
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func DoWork() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 10; i <= 13; i++ {
			ch2 <- i
		}
	}()

	// merge channel
	out := fanIn(ch1, ch2)
	for val := range out {
		fmt.Println(val)
	}
}
