package main

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	assertEqual := func(t testing.TB, got, want Stack) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, but want %v", got, want)
		}
	}

	t.Run("should push element into top of stack", func(t *testing.T) {
		stack := Stack{}
		stack.push(10)

		want := Stack{Head: &StackNode{Data: 10}}

		assertEqual(t, stack, want)
	})

	t.Run("should push more element into stack", func(t *testing.T) {
		stack := Stack{}
		stack.push(11)
		stack.push(12)

		want := Stack{Head: &StackNode{Data: 11, Next: &StackNode{Data: 12}}}

		assertEqual(t, stack, want)
	})
}

func assertEqual(t testing.TB, got, want Stack) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}
