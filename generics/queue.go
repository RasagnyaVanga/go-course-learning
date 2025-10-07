package main

import (
	"errors"
	"fmt"
)

type constraintsSet interface {
	int | float32
}
type Queue[T constraintsSet] struct {
	data []T
}

func (q *Queue[T]) Enqueue(dataitem T) {
	q.data = append(q.data, dataitem)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var res T
	if len(q.data) == 0 {
		return res, errors.New("cannot delete from empty queue")
	}
	res = q.data[0]
	q.data = q.data[1:]
	return res, nil
}

func (q *Queue[T]) Peek() (T, error) {
	var res T
	if len(q.data) == 0 {
		return res, errors.New("no elements in queue")
	}
	res = q.data[0]
	return res, nil
}
func main() {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q)
	fmt.Println(q.Peek())
	val, _ := q.Peek()
	fmt.Println(val)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
}
