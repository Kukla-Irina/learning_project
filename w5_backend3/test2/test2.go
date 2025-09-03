package main

import "fmt"

// Стек
type Stack struct {
	data []int
}

// Добавить
func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

// Убрать
func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last, true
}

// Очередь
type Queue struct {
	data []int
}

// Добавить
func (q *Queue) Enqueue(value int) {
	q.data = append(q.data, value)
}

// Убрать
func (q *Queue) Dequeue() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	first := q.data[0]
	q.data = q.data[1:]
	return first, true
}

func main() {
	// Стек
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Стек:")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	// Очередь
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("\nОчередь:")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
