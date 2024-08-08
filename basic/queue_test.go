package basic

import (
	"fmt"
	"testing"
	"time"
)

type Queue[T any] interface {
	Insert(item T)         //队列插入元素只能在尾部追加
	Remove() T             //队列的移除元素要从头部移除
	First() T              //读取队列头部的元素
	Size() int             //读取队列元素的个数
	Iterator() Iterator[T] //以队列当前的状态创建一个迭代器
	IsEmpty() bool         //判断队列是否为空
}
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type SliceQueue[T any] struct {
	items []T //!!!注意，含有切片的数据结构，要注意这样的类型所绑定的方法最好用指针访问，否则会有大量数据拷贝
}

func (sq *SliceQueue[T]) Insert(item T) {
	if IsNil(item) {
		panic("空值不允许插入到队列")
	}
	sq.items = append(sq.items, item)
}
func (sq *SliceQueue[T]) Remove() T {
	l := len(sq.items)
	if l == 0 {
		panic("队列已空，不能再删除元素")
	}
	item := sq.items[0]
	sq.items = sq.items[1:]
	return item
}
func (sq *SliceQueue[T]) First() T {
	if len(sq.items) == 0 {
		panic("队列已空，无法读取第一个元素")
	}
	return sq.items[0]
}

func (sq *SliceQueue[T]) Size() int {
	return len(sq.items)
}
func (sq *SliceQueue[T]) IsEmpty() bool {
	return len(sq.items) == 0
}
func (sq *SliceQueue[T]) Iterator() Iterator[T] {
	itrt := queueIterator[T]{
		indexOfNext: 0,
		items:       sq.items,
	}
	return &itrt
}

type queueIterator[T any] struct {
	indexOfNext int
	items       []T //!!!此实现中，由于item不是指针(*[]T)，这会导致数据的拷贝
}

func (qi *queueIterator[T]) HasNext() bool {
	l := len(qi.items)
	return qi.indexOfNext <= l-1
}
func (qi *queueIterator[T]) Next() T {
	if !qi.HasNext() {
		panic("迭代器已经没有下一个元素了！")
	}
	next := qi.items[qi.indexOfNext]
	qi.indexOfNext += 1
	return next
}

func TestQueue(t *testing.T) {
	var myQueue Queue[int] = &SliceQueue[int]{}
	myQueue.Insert(15)
	myQueue.Insert(20)
	myQueue.Insert(30)
	myQueue.Remove()
	fmt.Println(myQueue.First())
	queue := SliceQueue[float64]{}
	for i := 0; i < 10; i++ {
		queue.Insert(float64(i))
	}
	iterator := queue.Iterator()
	for {
		if !iterator.HasNext() {
			break
		}
		fmt.Println(iterator.Next())
	}
	fmt.Println("queue.First() = ", queue.First())
}

type NodeQueue[T any] struct {
	first, last *Node[T]
	length      int
}

// 队列插入元素只能在尾部追加
func (nq *NodeQueue[T]) Insert(item T) {
	if IsNil(item) {
		panic("空值不允许插入到队列")
	}
	nd := &Node[T]{value: item, next: nil}
	if nq.first == nil {
		nq.first = nd
		nq.last = nd
	} else {
		nq.last.next = nd
		nq.last = nd
	}
	nq.length += 1
}

// 队列的移除元素要从头部移除
func (nq *NodeQueue[T]) Remove() T {
	if nq.length == 0 {
		panic("队列已空，不能再删除元素")
	}
	result := nq.first.value
	nq.first = nq.first.next
	if nq.first == nil {
		nq.last = nil
	}
	nq.length -= 1
	return result

}

// 读取队列头部的元素
func (nq NodeQueue[T]) First() T {
	if nq.length == 0 {
		panic("队列已空，无法读取第一个元素")
	}
	return nq.first.value
}

// 读取队列元素的个数
func (nq *NodeQueue[T]) Size() int {
	return nq.length
}

// 以队列当前的状态创建一个迭代器
func (nq *NodeQueue[T]) Iterator() Iterator[T] {
	return &nodeQueueIterator[T]{nq.first}
}

// 判断队列是否为空
func (nq *NodeQueue[T]) IsEmpty() bool {
	return nq.length == 0
}

type nodeQueueIterator[T any] struct {
	nextNode *Node[T]
}

func (nqi *nodeQueueIterator[T]) HasNext() bool {
	return nqi.nextNode != nil
}
func (nqi *nodeQueueIterator[T]) Next() T {
	if !nqi.HasNext() {
		panic("迭代器已经没有下一个元素了！")
	}
	result := nqi.nextNode.value
	nqi.nextNode = nqi.nextNode.next
	return result
}
func TestNodeQueue(t *testing.T) {
	var myQueue Queue[int] = &NodeQueue[int]{}
	myQueue.Insert(15)
	myQueue.Insert(20)
	myQueue.Insert(30)
	myQueue.Remove()
	myQueue.Remove()
	myQueue.Remove()
	myQueue.Insert(3)
	fmt.Println(myQueue.First())
	queue := NodeQueue[float64]{}
	for i := 0; i < 10; i++ {
		queue.Insert(float64(i))
	}
	iterator := queue.Iterator()
	for {
		if !iterator.HasNext() {
			break
		}
		fmt.Println(iterator.Next())
	}
	fmt.Println("queue.First() = ", queue.First())
}
func TestCompareQueuesPerformance(t *testing.T) {
	sliceQueue := SliceQueue[int]{}
	nodeQueue := NodeQueue[int]{}
	start := time.Now()
	for i := 0; i < SIZE; i++ {
		sliceQueue.Insert(i)
	}
	elapsed := time.Since(start)
	fmt.Println("Time for inserting 1 million ints in sliceQueue is", elapsed)
	start = time.Now()
	for i := 0; i < SIZE; i++ {
		nodeQueue.Insert(i)
	}
	elapsed = time.Since(start)
	fmt.Println("Time for inserting 1 million ints in nodeQueue is", elapsed)
	start = time.Now()
	for i := 0; i < SIZE; i++ {
		sliceQueue.Remove()
	}
	elapsed = time.Since(start)
	fmt.Println("Time for removing 1 million ints from sliceQueue is", elapsed)
	start = time.Now()
	for i := 0; i < SIZE; i++ {
		nodeQueue.Remove()
	}
	elapsed = time.Since(start)
	fmt.Println("Time for removing 1 million ints from nodeQueue is", elapsed)
}
