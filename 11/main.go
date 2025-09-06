package main

import "fmt"

func main() {
	var (
		overlap []int
		n, t    int
		nums    map[int]int
	)

	fmt.Printf("Number of elements for A: ")
	fmt.Scan(&n)
	nums = make(map[int]int, n)
	fmt.Println("Enter elements for A:")
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		nums[t] = 0
	}

	fmt.Printf("Number of elements for B: ")
	fmt.Scan(&n)
	fmt.Println("Enter elements for B:")
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		if _, ok := nums[t]; ok {
			overlap = append(overlap, t)
		}
	}

	fmt.Println(overlap)

}
