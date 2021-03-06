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
	fmt.Println("Size of List (iterative method) :", Length(&head))
	fmt.Println("Size of list (recursive method) :", Count(head))

	//Delete(&head, 80)
	//PrintList(head)

	fmt.Println(Search(&head, 500))
	fmt.Println(Search(&head, 999))

	fmt.Println("Search (recursively) :", Find(head, 500))
	fmt.Println("Search (recursively) :", Find(head, 999))

	fmt.Println("Node at index 2 :", GetNth(&head, 2))
	fmt.Println("Node at index 2 (recursively) :", GetNthV2(head, 2))

	DeleteAt(3, &head)
	PrintList(head)
	fmt.Println("Size of List :", Length(&head))

	DeleteList(&head)
	PrintList(head)
	fmt.Println("Size of List :", Length(&head))
}
