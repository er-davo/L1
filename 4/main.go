package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

// go run main.go -workers 10

func worker(ctx context.Context, ch chan int, workerID int) {
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %d: channal closed\n", workerID)
				return
			}
			fmt.Printf("Worker %d: %d\n", workerID, val)
		case <-ctx.Done():
			fmt.Printf("Worker %d: context done\n", workerID)
			return
		}
	}

}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	workers := flag.Int("workers", 1, "number of workers")
	numsCh := make(chan int)
	var wg sync.WaitGroup

	flag.Parse()

	for i := 0; i < *workers; i++ {
		wg.Add(1)
		id := i
		go func() {
			defer wg.Done()
			worker(ctx, numsCh, id)
			fmt.Println("go done", id)
		}()
	}

loop:
	for i := 0; i < 100000; i++ {
		// так как отправка в канал блокирующая, то используем select
		// и выходим из цикла по ctx.Done()
		select {
		case numsCh <- i + 1:
			// успешно отправили
		case <-ctx.Done():
			fmt.Println("context canceled, stop sending")
			break loop
		}
	}

	close(numsCh)

	wg.Wait()

	fmt.Println("done")
}
