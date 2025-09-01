package main

import (
	"fmt"
	"sync"
)

const (
	stepCount    = 100000
	routineCount = 2
)

var counter int64

func main() {
	var wg sync.WaitGroup

	for i := 0; i < routineCount; i++ {
		wg.Add(1)
		go incr(&wg)
	}

	wg.Wait() // wait until all goroutines executed
	fmt.Printf("Step Count: %d\nLastValue: %d\nExpected: %d\n", stepCount, counter, stepCount*routineCount)
}

func incr(wg *sync.WaitGroup) {
	for i := 0; i < stepCount; i++ {
		counter++
	}

	wg.Done()
}
