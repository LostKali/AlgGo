package stack

import(
	"container/list"
	"fmt"
)

type StackOverList
[T any] struct {
	stack list.List
}

func (s *StackOverList
	[T]) IsEmpty() bool {
	return s.stack.Len() == 0
}

func (s *StackOverList
	[T]) Push(data T) {
	s.stack.PushBack(data)
}

func (s *StackOverList
	[T]) Pop() (T, error) {
	var zero T
	
	if s.IsEmpty() {
		return zero, fmt.Errorf("EmptyStackException")
	}

	back := s.stack.Back()
	s.stack.Remove(back)
	return back.Value.(T), nil
}

func (s *StackOverList
	[T]) Peek() (T, error) {
	var zero T

	if s.IsEmpty() {
		return zero, fmt.Errorf("EmptyStackException")
	}

	back := s.stack.Back()
	return back.Value.(T), nil
}