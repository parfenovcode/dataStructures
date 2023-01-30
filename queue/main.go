package main

import "fmt"

type Queue struct {
	items []int
}

// Enqueue adds a value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue takes out a value from the beggining
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)

	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	myQueue.Enqueue(7)
	fmt.Println(myQueue)

	myQueue.Dequeue()
	fmt.Println(myQueue)

}
