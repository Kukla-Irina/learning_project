package main

import "fmt"

func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	result := 0
	n := len(s)

	for i := 0; i < n; i++ {
		value := roman[s[i]]
		if i+1 < n && value < roman[s[i+1]] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
