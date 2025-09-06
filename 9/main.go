package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

func x(ctx context.Context, ch chan int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			close(ch)
			return
		case ch <- i:
		}
	}
}

func x2(chFrom, chTo chan int) {
	for num := range chFrom {
		chTo <- 2 * num
	}
	close(chTo)
}

func printer(ch chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	chX := make(chan int)
	chX2 := make(chan int)

	var wg sync.WaitGroup

	go x(ctx, chX)
	go x2(chX, chX2)

	wg.Go(func() {
		printer(chX2)
	})

	<-ctx.Done()

	// ждем пока printer выведет все числа
	wg.Wait()
}
