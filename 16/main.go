package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	// массив длиной 0 или 1 уже отсортирован
	if len(arr) < 2 {
		return arr
	}

	// выбираем опорный элемент - берём середину
	pivot := arr[len(arr)/2]

	// два среза для элементов меньше и больше (или равных) pivot
	var left, right []int

	// перебираем элементы массива
	for i, v := range arr {
		// пропускаем сам pivot
		if i == len(arr)/2 {
			continue
		}
		// распределяем элементы по сторонам
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	// рекурсивно сортируем левую и правую часть
	left = quickSort(left)
	right = quickSort(right)

	// отсортированный левый + pivot + отсортированный правый
	return append(append(left, pivot), right...)
}

func main() {
	data := []int{13, 512, 653, 4, 26, 8, 5, 632, 54, 3}

	sorted := quickSort(data)

	fmt.Println(sorted)
}
