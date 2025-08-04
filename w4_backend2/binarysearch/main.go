package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	numbers := []int{5, 2, 9, 1, 3}
	sort.Ints(numbers)
	fmt.Println("Отсортированный массив:", numbers)

	number := 2
	index := binarySearch(numbers, number)
	if index != -1 {
		fmt.Printf("Число %d найдено на индексе: %d\n", number, index)
	} else {
		fmt.Printf("Число %d не найдено", number)
	}
}
