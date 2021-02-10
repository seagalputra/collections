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

func DeleteAt(position int, node **Node) {

	if node == nil { return }

	// find node in the position
	var currentPos int
	temp := *node
	for temp != nil && currentPos != position {
		temp = temp.Next
		currentPos++
	}

	// find node in the previous position
	currentPos = 0
	prev := *node
	for prev != nil && currentPos != position - 1 {
		prev = prev.Next
		currentPos++
	}

	if prev != nil && temp != nil {
		(*prev).Next = (*temp).Next
	}

	if temp != nil {
		(*temp).Next = nil
	}
}

func DeleteList(head **Node) {
	*head = nil
}

func Length(head **Node) int {
	var counter int

	temp := *head
	for temp != nil {
		temp = temp.Next
		counter++
	}

	return counter
}

func Count(head *Node) int {
	if head == nil {
		return 0
	} else {
		return 1 + Count(head.Next)
	}
}

func Search(head **Node, target int) bool {
	current := *head
	for current != nil {
		if (*current).Data == target {
			return true
		}

		current = (*current).Next
	}

	return false
}
