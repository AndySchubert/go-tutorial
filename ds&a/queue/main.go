package main

import "fmt"

type Queue struct {
	items []int
}

// Enque adds a value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Deque removes a value at the front and returns the removed value
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := &Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
