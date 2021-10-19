package main

import (
	"reflect"
	"testing"
)

func TestAddNewElementAtLinkedList(t *testing.T) {
	got := LinkedList{}
	want := LinkedList{
		Head: &Node{
			Data: 10,
			Next: &Node{
				Data: 11,
				Next: &Node{
					Data: 12,
				},
			},
		},
	}

	got.Insert(10)
	got.Insert(11)
	got.Insert(12)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRemoveElementAtLinkedList(t *testing.T) {
	got := LinkedList{}
	want := LinkedList{
		Head: &Node{
			Data: 10,
			Next: &Node{
				Data: 12,
			},
		},
	}

	got.Insert(10)
	got.Insert(11)
	got.Insert(12)

	got.DeleteByKey(11)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
