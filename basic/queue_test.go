package basic

import (
	"fmt"
	"testing"
)

type Queue[T any] interface {
	Insert(itme T)         //队列插入元素只能在尾部追加
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
	items []T
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
func (sq SliceQueue[T]) First() T {
	if len(sq.items) == 0 {
		panic("队列已空，无法读取第一个元素")
	}
	return sq.items[0]
}

func (sq SliceQueue[T]) Size() int {
	return len(sq.items)
}
func (sq SliceQueue[T]) IsEmpty() bool {
	return len(sq.items) == 0
}
func (sq SliceQueue[T]) Iterator() Iterator[T] {
	itrt := QueueIterator[T]{
		indexOfNext: 0,
		items:       sq.items,
	}
	return &itrt
}

type QueueIterator[T any] struct {
	indexOfNext int
	items       []T
}

func (qi *QueueIterator[T]) HasNext() bool {
	l := len(qi.items)
	return qi.indexOfNext <= l-1
}
func (qi *QueueIterator[T]) Next() T {
	if !qi.HasNext() {
		panic("迭代器已经没有下一个元素了！")
	}
	next := qi.items[qi.indexOfNext]
	qi.indexOfNext += 1
	return next
}

func TestQueue(t *testing.T) {
	myQueue := SliceQueue[int]{}
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
