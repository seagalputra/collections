package main

type Tree struct {
	Parent *TreeNode
}

type TreeNode struct {
	Data       int
	LeftChild  *TreeNode
	RightChild *TreeNode
}

func (t *Tree) Insert(element int) {

	if t.Parent == nil {
		t.Parent = &TreeNode{Data: element}
		return
	}

	p := t.Parent
	for p.LeftChild != nil && p.RightChild != nil {
		if element > p.Data {

			if p.RightChild != nil {
				p = p.RightChild
			}
		}

		if element < p.Data {
			if p.LeftChild != nil {
				p = p.LeftChild
			}
		}
	}

	node := &TreeNode{Data: element}
	if element > p.Data {
		p.RightChild = node
	}

	if element < p.Data {
		p.LeftChild = node
	}
}

func SearchTree(root *TreeNode, element int) TreeNode {
	if root == nil || root.Data == element {
		return *root
	}

	if element < root.Data {
		return SearchTree(root.LeftChild, element)
	}

	return SearchTree(root.RightChild, element)
}

func InsertRecur(root *TreeNode, element int) *TreeNode {
	if root == nil {
		root = &TreeNode{Data: element}
		return root
	}

	if root.Data < element {
		root.RightChild = InsertRecur(root.RightChild, element)
	} else {
		root.LeftChild = InsertRecur(root.LeftChild, element)
	}

	return root
}

func DeleteTree(root *TreeNode, element int) *TreeNode {
	if root == nil {
		return root
	}

	// search element and assign them into left or right node
	// to make root value not replaced by next element
	if element < root.Data {
		root.LeftChild = DeleteTree(root.LeftChild, element)
	} else if element > root.Data {
		root.RightChild = DeleteTree(root.RightChild, element)
	} else { // if the element found

		// check whether the element has child (leaf) node or not
		if root.LeftChild == nil || root.RightChild == nil {
			return nil
		} else if root.LeftChild == nil {
			return root.RightChild
		} else if root.RightChild == nil {
			return root.LeftChild
		} else {
			// find minimum value in the right leaf node
			minVal := minValue(root.RightChild)
			root.Data = minVal
			root.RightChild = DeleteTree(root.RightChild, minVal)
		}
	}

	return root
}

func minValue(root *TreeNode) int {
	min := root.Data

	for root.LeftChild != nil {
		min = root.LeftChild.Data
		root = root.LeftChild
	}

	return min
}
