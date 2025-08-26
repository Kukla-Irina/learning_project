package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return p.Name
}

func printString(s Stringer) {
	fmt.Println(s.String())
}

func main() {

	p := Person{Name: "Bob"}
	printString(p)
	p2 := Person{Name: "Alice", Age: 30}
	fmt.Println(p2.Name)
}
