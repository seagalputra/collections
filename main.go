package main

func main() {
	head := &Node{Data: 12}
	Push(100, &head)

	printList(head)
}
