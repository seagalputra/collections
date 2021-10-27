package main

import "errors"

type QueueNode struct {
	Data int
	Next *QueueNode
}

type Queue struct {
	Front *QueueNode
	Rear  *QueueNode
}

func (q *Queue) Enqueue(data int) {

	node := &QueueNode{Data: data}
	head := q.Front

	if head == nil {
		q.Front = node
		q.Rear = node
		return
	}

	q.Rear.Next = node
	q.Rear = node
}

func (q *Queue) Dequeue() error {
	if q.Front == nil {
		return errors.New("failed to dequeue, list is empty")
	}

	q.Front = q.Front.Next
	return nil
}
