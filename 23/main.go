package main

import "fmt"

func removeAt[T any](s []T, i int) []T {
	copy(s[i:], s[i+1:])

	// обнуляем последний элемент
	// так как он остается в основном массиве
	var zeroValue T
	s[len(s)-1] = zeroValue

	return s[:len(s)-1]
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	nums = removeAt(nums, 5)

	fmt.Println(nums)
	fmt.Println(len(nums), cap(nums))
}
