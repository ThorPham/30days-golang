package main

import "fmt"

// d
type Stack struct {
	stack []int
}

func NewStack() Stack {
	return Stack{stack: []int{}}
}

func (s *Stack) addToStack(item int) {
	s.stack = append(s.stack, item)
}
func (s *Stack) removeItemFromStack() (int, error) {
	if len(s.stack) == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}
	tmp := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return tmp, nil

}
func main() {
	s := NewStack()
	s.addToStack(1)
	s.addToStack(2)
	s.addToStack(3)
	s.addToStack(4)
	fmt.Println(s.stack)
	s.removeItemFromStack()
	fmt.Println(s.stack)

}
