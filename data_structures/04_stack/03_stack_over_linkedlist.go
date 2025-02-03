package stack

import "errors"

type StackOverLinkedList struct {
	top *Node
}

type Node struct {
	data interface{}
	next *Node
}

func (s *StackOverLinkedList) Push(item interface{}) {
	t := &Node{
		data: item,
		next: s.top,
	}
	s.top = t
}

func (s *StackOverLinkedList) Pop() interface{} {
	if s.top == nil {
		panic(errors.New("EmptyStackException"))
	}
	t := s.top

	s.top = t.next

	return t.data
}

func (s *StackOverLinkedList) Peek() interface{} {
	if s.top == nil {
		panic(errors.New("EmptyStackException"))
	}

	return s.top.data
}

func (s *StackOverLinkedList) IsEmpty() bool {
	return s.top == nil
}
