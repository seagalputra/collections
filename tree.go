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
