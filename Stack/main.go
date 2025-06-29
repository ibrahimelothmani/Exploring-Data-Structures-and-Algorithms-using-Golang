package main

import "fmt"

type Stack struct {
	Items []int
}

func newStack() *Stack {
	return &Stack{Items: []int{}}
}

func (s *Stack) push(item int) {
	s.Items = append(s.Items, item)
}

func (s *Stack) pop() (int, bool) {
	if len(s.Items) == 0 {
		return 0, false
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item, true
}

func (s *Stack) peek() (int, bool) {
	if len(s.Items) == 0 {
		return 0, false
	}
	return s.Items[len(s.Items)-1], true
}
func (s *Stack) isEmpty() bool {
	return len(s.Items) == 0
}

func main() {
	stack := newStack()
	stack.push(1)
	stack.push(2)
	stack.push(3)

	fmt.Println("Stack contents:")
	for !stack.isEmpty() {
		if item, ok := stack.pop(); ok {
			fmt.Println(item)
		}
	}
}
