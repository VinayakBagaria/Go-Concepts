package synchronization

import "fmt"

func DoWork() {
	fmt.Println("mutex")
	MutexImplementation()
	fmt.Println("semaphore")
	SempahoreImplementation()
	fmt.Println("channel")
	ChannelImplementation()
	fmt.Println("crawler")
	CrawlerImplementation()
}
