package main

// AddElement insert new value in defined position into array
func AddElement(arr []int, position int, value int) []int {
	var newArr []int = make([]int, len(arr)+1)

	for i := 0; i < position; i++ {
		newArr[i] = arr[i]
	}

	newArr[position] = value

	for i := position + 1; i < len(newArr); i++ {
		newArr[i] = arr[i-1]
	}

	return newArr
}

// Search element from given array
func Search(arr []int, value int) int {
	var foundIndex int = -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			foundIndex = i
			break
		}
	}

	return foundIndex
}

// DeleteElement remove value from specific index of array
func DeleteElement(arr []int, value int) []int {
	var position int = Search(arr, value) // if value not found returned -1
	if position >= 0 {
		for i := position + 1; i < len(arr); i++ {
			arr[i-1] = arr[i]
		}
	}
	return arr[:len(arr)-1]
}
