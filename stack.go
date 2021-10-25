package main

import "errors"

var ErrStackPopElement = errors.New("failed to pop the element, stack is empty")

type StackNode struct {
	Data int
	Next *StackNode
}

type Stack struct {
	Head *StackNode
}

func (s *Stack) push(value int) {
	if s.Head == nil {
		s.Head = &StackNode{Data: value}
		return
	}

	head := s.Head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &StackNode{Data: value}
}

func (s *Stack) pop() error {

	if s.Head == nil {
		return ErrStackPopElement
	}

	head := s.Head
	for head != nil {
		if head.Next.Next == nil {
			head.Next = nil
			return nil
		}
		head = head.Next
	}

	return nil
}
