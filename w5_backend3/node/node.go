package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

// Добавление
func (l *LinkedList) Add(value int) {
	newNode := &Node{Value: value}
	newNode.Next = l.Head
	l.Head = newNode
}

// Удаление
func (l *LinkedList) Remove(value int) {
	if l.Head == nil {
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.Value == value {
			prev.Next = curr.Next
			return
		}
		prev = curr
		curr = curr.Next
	}
}

func (l *LinkedList) Print() {
	curr := l.Head
	for curr != nil {
		fmt.Print(curr.Value, " ")
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	list.Add(3)
	list.Add(2)
	list.Add(1)
	fmt.Print("После добавления: ")
	list.Print()

	list.Remove(2)
	fmt.Print("После удаления 2: ")
	list.Print()
}
