package main

import "fmt"

type Queue struct {
	queue []int
}

func newQueue() Queue {
	return Queue{queue: []int{}}
}
func (q *Queue) enqueue(item int) {
	q.queue = append(q.queue, item)
}
func (q *Queue) dequeue() (int, error) {
	if len(q.queue) == 0 {
		return 0, fmt.Errorf("queue is empty")
	} else if len(q.queue) == 1 {
		item := q.queue[0]
		q.queue = []int{}
		return item, nil

	} else {
		item := q.queue[0]
		q.queue = q.queue[1:]
		return item, nil
	}

}

func main() {
	q := newQueue()
	q.enqueue(4)
	q.enqueue(3)
	q.enqueue(2)
	fmt.Printf("Queue: %s\n", q.queue)
	q.dequeue()
	fmt.Printf("Queue: %s\n", q.queue)
}
