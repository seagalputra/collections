package main

func main() {
	head := &Node{Data: 12}
	first := &Node{Data: 80}
	head.Next = first

	Push(100, &head)
	InsertAfter(&first, 500)
	Append(&head, 1_000_000)

	PrintList(head)
}
