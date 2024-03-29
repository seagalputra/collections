package main

import (
	"reflect"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("Should add element into tree", func(t *testing.T) {
		got := Tree{}
		got.Insert(10)
		got.Insert(7)
		got.Insert(15)

		want := Tree{
			Parent: &TreeNode{
				Data: 10,
				LeftChild: &TreeNode{
					Data: 7,
				},
				RightChild: &TreeNode{
					Data: 15,
				},
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Should add 5 into existing tree", func(t *testing.T) {
		got := Tree{}
		got.Insert(10)
		got.Insert(7)
		got.Insert(15)
		got.Insert(5)

		want := Tree{
			Parent: &TreeNode{
				Data: 10,
				LeftChild: &TreeNode{
					Data: 7,
					LeftChild: &TreeNode{
						Data: 5,
					},
				},
				RightChild: &TreeNode{
					Data: 15,
				},
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Should add 12 into existing tree", func(t *testing.T) {
		got := Tree{}
		got.Insert(10)
		got.Insert(7)
		got.Insert(15)
		got.Insert(5)
		got.Insert(12)

		want := Tree{
			Parent: &TreeNode{
				Data: 10,
				LeftChild: &TreeNode{
					Data: 7,
					LeftChild: &TreeNode{
						Data: 5,
					},
				},
				RightChild: &TreeNode{
					Data: 15,
					LeftChild: &TreeNode{
						Data: 12,
					},
				},
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestBinarySearchTreeSearch(t *testing.T) {
	t.Run("should return the element node when search is found", func(t *testing.T) {
		tree := Tree{
			Parent: &TreeNode{
				Data: 10,
				LeftChild: &TreeNode{
					Data: 7,
				},
				RightChild: &TreeNode{
					Data: 15,
				},
			},
		}

		got := SearchTree(tree.Parent, 15)

		want := TreeNode{Data: 15}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestBinarySearchTreeDelete(t *testing.T) {
	t.Run("should return the element node when search is found", func(t *testing.T) {
		leftChild := &TreeNode{
			Data: 7,
		}

		rightChild := &TreeNode{
			Data: 15,
		}

		tree := Tree{
			Parent: &TreeNode{
				Data:       10,
				LeftChild:  leftChild,
				RightChild: rightChild,
			},
		}

		got := DeleteTree(tree.Parent, 10)

		want := &TreeNode{
			Data:      15,
			LeftChild: leftChild,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestBinarySearchTreeRecur(t *testing.T) {
	t.Run("Should add element into tree", func(t *testing.T) {
		var got *TreeNode
		got = InsertRecur(got, 10)
		got = InsertRecur(got, 7)
		got = InsertRecur(got, 15)

		want := &TreeNode{
			Data: 10,
			LeftChild: &TreeNode{
				Data: 7,
			},
			RightChild: &TreeNode{
				Data: 15,
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
