package main

import "fmt"

func main() {
	head := &Node{Data: 12}
	first := &Node{Data: 80}
	head.Next = first

	Push(100, &head)
	InsertAfter(&first, 500)
	Append(&head, 1_000_000)

	PrintList(head)
	fmt.Println("Size of List :", Length(&head))

	//Delete(&head, 80)
	//PrintList(head)

	DeleteAt(3, &head)
	PrintList(head)
	fmt.Println("Size of List :", Length(&head))

	DeleteList(&head)
	PrintList(head)
	fmt.Println("Size of List :", Length(&head))
}
