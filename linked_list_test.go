package main

import (
	"testing"
)

func TestAddNewElementAtLinkedList(t *testing.T) {
	head := &Node{Data: 10} // this holds address of head node
	list := LinkedList{Head: head}

	list.Insert(18)
	PrintList(list)
}
