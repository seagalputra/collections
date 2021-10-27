package main

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("should add element in queue", func(t *testing.T) {
		got := Queue{}
		got.Enqueue(10)

		node := &QueueNode{Data: 10}
		want := Queue{Front: node, Rear: node}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should enqueue more element in queue", func(t *testing.T) {
		got := Queue{}
		got.Enqueue(11)
		got.Enqueue(12)

		second := &QueueNode{
			Data: 12,
		}
		first := &QueueNode{
			Data: 11,
			Next: second,
		}

		want := Queue{Front: first, Rear: second}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should dequeue element in queue", func(t *testing.T) {
		got := Queue{}
		got.Enqueue(11)
		got.Enqueue(12)
		got.Dequeue()

		first := &QueueNode{
			Data: 12,
		}
		want := Queue{Front: first, Rear: first}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("tried dequeue element in empty queue", func(t *testing.T) {
		got := Queue{}
		err := got.Dequeue()

		if err == nil {
			t.Fatalf("should raise an error")
		}

		if err.Error() != "failed to dequeue, list is empty" {
			t.Errorf("got %s, want %s", "'"+err.Error()+"'", "'failed to dequeue, list is empty'")
		}
	})
}
