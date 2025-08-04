package main

import "fmt"

func bubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func main() {
	numbers := []int{5, 2, 9, 1, 3}
	bubbleSort(numbers)
	fmt.Println(numbers)
}
