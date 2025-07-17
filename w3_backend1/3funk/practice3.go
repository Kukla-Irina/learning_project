package main

import "fmt"

func sumMap(m map[string]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	m := map[string]int{"мышь": 4, "компьютер": 6, "монитор": 8}
	fmt.Println("Sum:", sumMap(m))
}
