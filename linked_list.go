package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func printList(node *Node) {
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
