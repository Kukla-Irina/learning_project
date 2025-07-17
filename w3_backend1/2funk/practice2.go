package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5, 6}
	for i := len(num) - 1; i >= 0; i-- {
		fmt.Print(num[i], " ")
	}
}
