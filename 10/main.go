package main

import "fmt"

func main() {
	nums := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64, len(nums))

	for _, num := range nums {
		// округление до целого
		floor := int(num)
		// округление до десятков
		floor -= floor % 10

		groups[floor] = append(groups[floor], num)
	}

	fmt.Print(groups)
}
