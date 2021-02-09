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
	fmt.Println()
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

func Delete(node **Node, key int) {

	if node == nil { return }

	// find node with given key
	temp := *node
	for temp != nil && temp.Data != key {
		temp = temp.Next
	}

	// find previous node
	prev := *node
	for prev != nil && prev.Next.Data != key {
		prev = prev.Next
	}

	// set prev node to next of deleted node
	if prev != nil && temp != nil {
		(*prev).Next = (*temp).Next
	}

	// set next deleted node to nil
	if temp != nil {
		(*temp).Next = nil
	}
}
