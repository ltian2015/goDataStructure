package basic

import (
	"fmt"
	"testing"
)

type PriorityQueue[T any] struct {
	q    []SliceQueue[T] // 队列的切片,切片的序号（index）表示所含队列元素的优先级。
	size int
}

func NewPriorityQueue[T any](numberPriorities int) (pq PriorityQueue[T]) {
	pq.q = make([]SliceQueue[T], numberPriorities)
	return pq
	/** 上面的代码等价于以下代码
	pq = PriorityQueue[T]{
		q:    make([]SliceQueue[T], numberPriorities),
		size: 0}
	**/
}

func (pq *PriorityQueue[T]) Insert(item T, priority int) {
	pq.q[priority-1].Insert(item)
	pq.size++
}

func (pq *PriorityQueue[T]) Remove() T {
	if pq.size == 0 {
		panic("队列已空，不能再移除元素")
	}
	var result T
	for i := 0; i < len(pq.q); i++ {
		if pq.q[i].Size() > 0 {
			result = pq.q[i].Remove()
			pq.size--
			break
		}
	}
	return result
}

func (pq *PriorityQueue[T]) First() T {
	if pq.size == 0 {
		panic("队列已空，无法获得头元素")
	}
	var result T
	for i := 0; i < len(pq.q); i++ {
		q := pq.q[i]
		if q.Size() > 0 {
			result = q.First()
			break
		}
	}
	return result
}
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.size == 0
}

type Passenger struct {
	name     string
	priority int
}

func TestPriorityQueue(t *testing.T) {
	airlineQueue := NewPriorityQueue[Passenger](3)
	passengers := []Passenger{{"Erika", 3}, {"Robert", 3}, {"Danielle", 3},
		{"Madison", 1}, {"Frederik", 1}, {"James", 2},
		{"Dante", 2}, {"Shelley", 3}}
	fmt.Println("Passsengers: ", passengers)
	for i := 0; i < len(passengers); i++ {
		airlineQueue.Insert(passengers[i], passengers[i].priority)
	}
	fmt.Println("First passenger in line: ", airlineQueue.First())
	airlineQueue.Remove()
	airlineQueue.Remove()
	airlineQueue.Remove()
	fmt.Println("First passenger in line: ", airlineQueue.First())

}
