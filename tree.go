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
