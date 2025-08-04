package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

func main() {
	example1 := []int{1, 2, 3, 1}
	fmt.Println("Пример 1:", containsDuplicate(example1))
	example2 := []int{1, 2, 3, 4}
	fmt.Println("Пример 2:", containsDuplicate(example2))
	example3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println("Пример 3:", containsDuplicate(example3))
	example4 := []int{9, 8, 7, 6, 5, 4}
	fmt.Println("Пример 4:", containsDuplicate(example4))
}
