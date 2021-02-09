package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func PrintList(node *Node) {
	for node != nil {
		fmt.Printf("%d ", node.Data)
		node = node.Next
	}
}

func Push(data int, head **Node) {
	newNode := &Node{Data: data}
	newNode.Next = *head
	*head = newNode
}

func InsertAfter(prevNode **Node, data int) {
	newNode := &Node{Data: data}
	newNode.Next = (*prevNode).Next
	(*prevNode).Next = newNode
}

func Append(head **Node, data int) {
	newNode := &Node{Data: data}

	if (*head).Next == nil {
		(*head).Next = newNode
	}

	node := *head
	for node.Next != nil {
		node = node.Next
	}

	node.Next = newNode
}
