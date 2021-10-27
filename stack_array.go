package main

type StackArray struct {
	Data [5]int
	Top  int
}

func (s *StackArray) Push(elm int) {
	// TODO: check whether the stack is full or not
	s.Data[s.Top] = elm
	s.Top++
}

func (s *StackArray) Pop() {
	// TODO: check whether stack is empty
	s.Top--
	s.Data[s.Top] = 0
}
