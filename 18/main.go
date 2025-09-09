package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) Count() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			counter.Increment()
		})
	}

	wg.Wait()

	fmt.Println(counter.Count())
}
