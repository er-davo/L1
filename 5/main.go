package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

// go run main.go -time=10

func consumer(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ch:
			//fmt.Printf("Consumer: %d\n", v)
		case <-ctx.Done():
			fmt.Println("Consumer stopped by contex")
			return
		}
	}
}

func producer(ctx context.Context, ch chan<- int) {
	i := 0
	for {
		select {
		case ch <- i:
		case <-ctx.Done():
			fmt.Println("Producer stopped by contex")
			return
		}
		i++
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	n := flag.Int("time", 3, "duration of program in seconds")
	ch := make(chan int)
	flag.Parse()

	go producer(ctx, ch)
	go consumer(ctx, ch)

	<-time.After(time.Second * time.Duration(*n))
	cancel()
}
