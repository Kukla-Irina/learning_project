package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	lastNonZero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZero] = nums[i]
			lastNonZero++
		}
	}

	for i := lastNonZero; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums2 := []int{0}
	moveZeroes(nums2)
	fmt.Println(nums2)
}
