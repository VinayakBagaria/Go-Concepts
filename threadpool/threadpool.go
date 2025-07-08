package threadpool

import (
	"fmt"
	"sync"
	"time"
)

type Job func()

type Pool struct {
	workQueue chan Job
	wg        sync.WaitGroup
}

func NewPool(workerCount int) *Pool {
	wg := sync.WaitGroup{}
	pool := &Pool{
		workQueue: make(chan Job),
		wg:        wg,
	}

	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for job := range pool.workQueue {
				job()
			}
		}()
	}

	return pool
}

func (p *Pool) Add(job Job) {
	p.workQueue <- job
}

func (p *Pool) Wait() {
	p.wg.Wait()
	close(p.workQueue)
}

func DoWork() {
	pool := NewPool(5)
	for i := 0; i < 30; i++ {
		job := func() {
			time.Sleep(1 * time.Second)
			fmt.Println("job completed")
		}

		pool.Add(job)
	}

	pool.Wait()
}
