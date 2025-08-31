package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(5)

	// выход по условию
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i == 5 {
				fmt.Println("Выход по условию")
				return
			}
		}
	}()

	// выход по контексту
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Выход по контексту")
				return
			default:
			}
		}
	}(ctx)
	cancel()

	// выход по таймауту

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Выход по таймауту")
				return
			default:
			}
		}
	}(ctx)

	// выход через канал
	ch := make(chan struct{})
	go func(ch <-chan struct{}) {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("Выход через канал")
				return
			default:
			}
		}
	}(ch)
	close(ch)

	// прекращение работы
	go func() {
		defer func() {
			fmt.Println("Прекращение работы")
			defer wg.Done()
		}()
		runtime.Goexit()
	}()

	wg.Wait()
}
