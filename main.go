package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{1, 5, 2, 8, 7, 3}

	// accessing array
	fmt.Println(arr[3])

	// adding new element
	// fmt.Println(addElement(arr, 3, 17))
	fmt.Println(AddElement(arr, 0, 18))

	// search element in array
	fmt.Printf("Element 8 found at index : %d\n", Search(arr, 8))

	// delete element at given position in array
	fmt.Println(DeleteElement(arr, 7))
	// fmt.Println(deleteElement(arr, 5))
}
