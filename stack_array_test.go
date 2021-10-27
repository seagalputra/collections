package main

import (
	"reflect"
	"testing"
)

func TestStackArrayImpl(t *testing.T) {
	t.Run("should able to push element into empty stack", func(t *testing.T) {
		var got StackArray = StackArray{}
		got.Push(10)

		want := StackArray{Data: [5]int{10}}

		assertStackEqual(t, got, want)
	})

	t.Run("should able to push more element to empty stack", func(t *testing.T) {
		var got StackArray = StackArray{}
		got.Push(12)
		got.Push(14)

		want := StackArray{Data: [5]int{12, 14}}

		assertStackEqual(t, got, want)
	})

	t.Run("should able to pop element in stack", func(t *testing.T) {
		var got StackArray = StackArray{}
		got.Push(12)
		got.Push(14)
		got.Pop()

		want := StackArray{Data: [5]int{12}}

		assertStackEqual(t, got, want)
	})
}

func assertStackEqual(t testing.TB, got, want StackArray) {
	t.Helper()
	if !reflect.DeepEqual(got.Data, want.Data) {
		t.Errorf("got %v, want %v", got.Data, want.Data)
	}
}
