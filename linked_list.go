package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func PrintList(list LinkedList) {
	var head *Node = list.Head
	fmt.Print("List element: ")
	for head != nil {
		fmt.Printf("%v ", head.Data)
		head = head.Next
	}
	fmt.Println()
}

// Insert value at last node in linked list
func (l *LinkedList) Insert(data int) {
	if l == nil {
		l.Head = &Node{Data: data}
		return
	}

	head := l.Head
	for head.Next != nil {
		head = head.Next
	}

	head.Next = &Node{Data: data}
}
