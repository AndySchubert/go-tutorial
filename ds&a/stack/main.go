package main

import "fmt"

type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	newStack := &Stack{}
	fmt.Println(newStack)
	newStack.Push(100)
	newStack.Push(200)
	newStack.Push(300)
	fmt.Println(newStack)
	fmt.Println(newStack.Pop())
}
