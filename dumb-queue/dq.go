package main

import (
	"fmt"
)

type Queue struct {
	data map[int]string
	head int
	tail int
}

func NewQueue() *Queue {
	var q Queue
	q.data = make(map[int]string)
	q.head = 0
	q.tail = 0
	return &q
}

func (q *Queue) enqueue(value string) map[int]string {
	if value == "" {
		return q.data
	}

	q.data[q.tail] = value
	q.tail++

	return q.data
}

func (q *Queue) dequeue() map[int]string {
	if q.head == q.tail {
		return q.data
	}

	delete(q.data, q.head)
	q.head++

	return q.data
}

func (q *Queue) size() int {
	return q.tail - q.head
}

func (q *Queue) show() map[int]string {
	return q.data
}

func main() {
	q := NewQueue()
	fmt.Printf("%+v\n", q)
	q.enqueue("cat")
	fmt.Println(q.show(), q.size())
	q.enqueue("dog")
	fmt.Println(q.show(), q.size())
	q.enqueue("starfish")
	fmt.Println(q.show(), q.size())
	q.dequeue()
	fmt.Println(q.show(), q.size())
	q.dequeue()
	fmt.Println(q.show(), q.size())
	q.dequeue()
	fmt.Println(q.show(), q.size())
	q.dequeue()
	fmt.Println(q.show(), q.size())
}
