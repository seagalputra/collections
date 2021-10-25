package main

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
