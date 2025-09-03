package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	first := arr[0]
	var less, greater []int

	for _, v := range arr[1:] {
		if v <= first {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	less = quickSort(less)
	greater = quickSort(greater)

	result := append(less, first)
	result = append(result, greater...)
	return result
}

func main() {
	arr := []int{5, 3, 8, 1, 2, 7}
	fmt.Println("Исходный:", arr)
	sorted := quickSort(arr)
	fmt.Println("Отсортированный:", sorted)
}
