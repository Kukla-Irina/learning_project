package main

import (
	"fmt"
)

func reverse(x int) int {
	rev := 0

	for x != 0 {
		digit := x % 10
		rev = rev*10 + digit
		x = x / 10
	}

	return rev
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}
