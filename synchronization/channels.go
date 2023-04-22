package synchronization

import "fmt"

const size = 1000

var count = 0

func ChannelImplementation() {
	c := make(chan int)
	defer close(c)

	for i := 0; i < size; i++ {
		go func() {
			c <- i
		}()
	}

	for count < size {
		count += <-c
	}

	fmt.Println(count)
}
