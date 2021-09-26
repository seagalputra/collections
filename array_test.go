package main

import (
	"reflect"
	"testing"
)

func TestAddNewElementAtArray(t *testing.T) {
	t.Skip()
	var arr []int = []int{1, 5, 2, 8, 7, 3}

	got := AddElement(arr, 0, 18)
	want := []int{18, 1, 5, 2, 8, 7, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
