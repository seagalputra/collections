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

func (s *Stack) Push(value int) {
	node := &StackNode{Data: value, Next: s.Head}
	s.Head = node
}

func (s *Stack) Pop() error {

	if s.Head == nil {
		return ErrStackPopElement
	}

	s.Head = s.Head.Next

	return nil
}
