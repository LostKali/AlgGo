package stack

import "errors"

type Stack struct {
	stack []interface{}
	top   int
}

func NewStack(size int) *Stack {
	return &Stack{
		stack: make([]interface{}, size),
		top:   -1,
	}
}

func (s *Stack) Push(item interface{}) {
	if s.top == len(s.stack)-1 {
		panic(errors.New("Stack is full"))
	}
	s.top++
	s.stack[s.top] = item
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic(errors.New("EmptyStackException"))
	}
	item := s.stack[s.top]
	s.stack[s.top] = nil
	s.top--

	return item
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		panic(errors.New("EmptyStackException"))
	}

	return s.stack[s.top]
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}
