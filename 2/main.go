package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for i := range nums {
		num := nums[i]
		wg.Go(func() {
			fmt.Println(num * num)
		})
	}

	wg.Wait()

}
