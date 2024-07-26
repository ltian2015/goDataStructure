package basic

import "testing"

// ///////////////////////////////////以下是单向列表的操作
// 包外不可见
type node[T any] struct {
	value T
	next  *node[T]
}
type SingleLinkedList[T any] struct {
	head *node[T]
	tail *node[T] //方便Append操作，即，在尾部追加节点
	size int
}

func (sll *SingleLinkedList[T]) Size() int {
	return sll.size
}
func (sll *SingleLinkedList[T]) IsEmpty() bool {
	return sll.size == 0
}
func (sll *SingleLinkedList[T]) addNodeToEmpyList(nd *node[T]) {
	if sll.size > 0 {
		panic("列表不空，操作错误")
	}
	sll.head = nd
	sll.tail = nd
	sll.size++
}
func (sll *SingleLinkedList[T]) Append(value T) {
	nd := &node[T]{
		value: value,
		next:  nil,
	}
	if sll.size == 0 {
		sll.addNodeToEmpyList(nd)
		return
	}
	sll.tail.next = nd
	sll.tail = nd
	sll.size++

}
func (sll *SingleLinkedList[T]) Insert(i int, value T) {
	newNd := &node[T]{value: value, next: nil}
	if sll.size == 0 {
		sll.addNodeToEmpyList(newNd)
		return
	}
	position := i
	if position < 0 {
		position = 0
	} else if position > sll.size {
		position = sll.size
	}
	if position == 0 {
		oldHead := sll.head
		sll.head = newNd
		newNd.next = oldHead
		sll.size++
		return
	}
	if position == sll.size {
		sll.tail.next = newNd
		sll.tail = newNd
		sll.size++
		return
	}
	preNode := sll.getNode(position - 1)
	curNodeAtPosition := sll.getNode(position)
	preNode.next = newNd
	newNd.next = curNodeAtPosition
	sll.size++
}
func (sll *SingleLinkedList[T]) Delete(i int) T {
	if sll.IsEmpty() {
		panic("试图从空列表中删除元素")
	}
	var ndTobeDelete *node[T]
	if i == 0 {
		ndTobeDelete = sll.head
		sll.head = ndTobeDelete.next
		sll.size--
		return ndTobeDelete.value
	}
	preNode := sll.getNode(i - 1)
	ndTobeDelete = sll.getNode(i)
	preNode.next = ndTobeDelete.next
	sll.size--
	return ndTobeDelete.value
}
func (sll *SingleLinkedList[T]) getNode(i int) *node[T] {
	if i < 0 || i > sll.size || sll.IsEmpty() {
		panic("无法获取非法序号的节点")
	}
	index := 0
	nd := sll.head
	for {
		if index == i {
			return nd
		} else {
			nd = nd.next
			index++
		}
	}
}
func (sll *SingleLinkedList[T]) Get(i int) T {
	nd := sll.getNode(i)
	return nd.value
}

// !!! 这个函数来自于go 1.21.0 开始发布的slices包。
// !!! SingleLinkedList[T any] 给出了任何类型的列表容器的通用操作，无法给出针对可比较类型（comparable）元素
// !!! 的查找（Find）操作，也无法给出可排序类型（ordered)元素的找到最大值(Max)、最小值(Min、排序(Sort)等操作.
// !!! 故而根据 [T any]可能是comparable或ordered类型，给出相应的辅助数据与行为分离的函数式编程思想的运用，即，
// !!! 根据数据类型的共性特征（由接口所代表的方法集）来给出独立的操作函数。这样可解决面向对象编程思想中的一些约束问题，
// !!! 比如，强制要求所操作元素的类型必须拥有特定的特征，比如要求元素必须是可比较（comparable）的或可排序的（ordered）。
func Find[E comparable](sll SingleLinkedList[E], e E) int {
	result := -1
	nd := sll.head
	index := 0
	for {
		if nd.value == e {
			result = index
			return result
		} else if nd.next == nil {
			return result
		} else {
			nd = nd.next
			index++
		}
	}
}

func TestSingleLinkedList(t *testing.T) {
	var slli SingleLinkedList[int]
	slli.Append(1)
	slli.Append(2)
	slli.Append(3)
	i := Find(slli, 0)
	println(i)
	var slls SingleLinkedList[string]
	slls.Append("One")
	slls.Append("Two")
	slls.Append("Three")
	i = Find(slls, "One")
	println(i)
	println(slls.size)
	slls.Delete(1)
	println(slls.size)
	println(Find(slls, "Two"))
	slls.Insert(1, "Two")
	println(Find(slls, "Two"))
}
