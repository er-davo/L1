package main

import (
	"flag"
	"fmt"
	"sync"
)

// go run main.go -workers 10

func worker(ch chan int, workerID int) {
	for val := range ch {
		fmt.Printf("Worker %d: %d\n", workerID, val)
	}
}

func main() {
	workers := flag.Int("workers", 1, "number of workers")
	numsCh := make(chan int)
	var wg sync.WaitGroup

	flag.Parse()

	for i := 0; i < *workers; i++ {
		id := i
		wg.Go(func() {
			worker(numsCh, id)
		})
	}

	for i := 0; i < 10; i++ {
		numsCh <- i + 1
	}

	close(numsCh)

	wg.Wait()
}
