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

	t.Run("should pop element from stack", func(t *testing.T) {
		stack := Stack{}
		stack.push(10)
		stack.push(11)
		stack.push(12)
		stack.push(13)
		stack.pop()

		want := Stack{Head: &StackNode{Data: 10, Next: &StackNode{Data: 11, Next: &StackNode{Data: 12}}}}

		assertEqual(t, stack, want)
	})

	t.Run("should raise error when try pop element from empty stack", func(t *testing.T) {
		stack := Stack{}
		err := stack.pop()

		assertError(t, err, ErrStackPopElement)
	})
}

func assertEqual(t testing.TB, got, want Stack) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatalf("expected an error but got nothing")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
