package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] < target {
			// смещаем левую границу вправо
			left = mid + 1
		} else if arr[mid] > target {
			// смещаем правую границу влево
			right = mid - 1
		} else {
			// нашли
			return mid
		}
	}

	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(nums, 5))
}
