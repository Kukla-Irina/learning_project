package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	orig := x
	rev := 0

	for x > 0 {
		rev = rev*10 + x%10
		x = x / 10
	}

	return orig == rev

}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
