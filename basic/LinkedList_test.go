package basic

import (
	"testing"
)

type List[T any] interface {
	First() T             //Returns the first node in the list
	Size() int            //Returns the number of nodes in the list
	Insert(i int, item T) //Creates and inserts item in the ith node of the list
	RemoveAt(i int) T     //Removes and returns the item in the ith node of the list
	Append(item T)        //Creates and inserts item into the last node of the list
	Get(i int) T          // Returns the node position containing item in the list
	Items() []T           //Returns a slice of all the items in the list
}

// !!! 这个函数来自于go 1.21.0 开始发布的slices包。
// !!! SingleLinkedList[T any] 给出了任何类型的列表容器的通用操作，无法给出针对可比较类型（comparable）元素
// !!! 的查找（Find）操作，也无法给出可排序类型（ordered)元素的找到最大值(Max)、最小值(Min、排序(Sort)等操作.
// !!! 故而根据 [T any]可能是comparable或ordered类型，给出相应的辅助数据与行为分离的函数式编程思想的运用，即，
// !!! 根据数据类型的共性特征（由接口所代表的方法集）来给出独立的操作函数。这样可解决面向对象编程思想中的一些约束问题，
// !!! 比如，强制要求所操作元素的类型必须拥有特定的特征，比如要求元素必须是可比较（comparable）的或可排序的（ordered）。
func Find[E comparable](l List[E], e E) int {
	result := -1
	for index := 0; index < l.Size(); index++ {
		if l.Get(index) == e {
			result = index
			break
		}
	}
	return result
}

// ///////////////////////////////////以下是单向列表的操作
// 包外不可见

type SingleLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T] //方便Append操作，即，在尾部追加节点,通常的单向链表没有这个节点
	size int
}

func (sll *SingleLinkedList[T]) Size() int {
	return sll.size
}
func (sll *SingleLinkedList[T]) IsEmpty() bool {
	return sll.size == 0
}

func (sll SingleLinkedList[T]) First() T {
	return sll.head.value
}
func (sll SingleLinkedList[T]) Items() []T {
	result := []T{}
	for node := sll.head; node.next != nil; node = node.next {
		result = append(result, node.value)
	}
	return result

}
func (sll *SingleLinkedList[T]) addNodeToEmpyList(nd *Node[T]) {
	if sll.size > 0 {
		panic("列表不空，操作错误")
	}
	sll.head = nd
	sll.tail = nd
	sll.size++
}
func (sll *SingleLinkedList[T]) Append(value T) {
	nd := &Node[T]{
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
	newNd := &Node[T]{value: value, next: nil}
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
func (sll *SingleLinkedList[T]) RemoveAt(i int) T {
	if sll.IsEmpty() {
		panic("试图从空列表中删除元素")
	}
	var ndTobeDelete *Node[T]
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
func (sll *SingleLinkedList[T]) getNode(i int) *Node[T] {
	if i < 0 || i >= sll.size || sll.IsEmpty() {
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

func TestSingleLinkedList(t *testing.T) {
	var slli List[int] = &SingleLinkedList[int]{}
	slli.Append(1)
	slli.Append(2)
	slli.Append(3)
	i := Find(slli, 0)
	println(i)
	var slls List[string] = &SingleLinkedList[string]{}
	slls.Append("One")
	slls.Append("Two")
	slls.Append("Three")
	i = Find(slls, "One")
	println(i)
	println(slls.Size())
	slls.RemoveAt(1)
	println(slls.Size())
	println(Find(slls, "Two"))
	slls.Insert(1, "Two")
	println(Find(slls, "Two"))
}

type DoubleLinkedList[T any] struct {
	head *DNode[T]
	tail *DNode[T]
	size int
}

func (dll DoubleLinkedList[T]) First() T { //Returns the first node in the list

	return dll.head.value
}

func (dll DoubleLinkedList[T]) Size() int { //Returns the number of nodes in the list
	return dll.size
}
func (dll *DoubleLinkedList[T]) addNodeToEmpyList(nd *DNode[T]) {
	if dll.size != 0 {
		panic("列表不空，操作错误")
	}
	dll.head = nd
	dll.tail = nd
	dll.size += 1
}
func (dll *DoubleLinkedList[T]) doInsert(oldNode, nd *DNode[T]) {
	if oldNode == nil {
		panic("插入所在位置的节点不存在")
	}
	preNode := oldNode.pre
	preNode.next = nd
	nd.pre = preNode
	nd.next = oldNode
	oldNode.pre = nd
	dll.size += 1
}
func (dll *DoubleLinkedList[T]) Insert(i int, item T) { //Creates and inserts item in the ith node of the list
	if IsNil(item) {
		panic("不允许向列表插入空对象！")
	}

	nd := &DNode[T]{value: item, pre: nil, next: nil}
	if dll.size == 0 {
		dll.addNodeToEmpyList(nd)
		return
	}
	if i < 0 {
		dll.doInsert(dll.head, nd)
		return
	}
	if i >= dll.size {
		dll.Append(item)
		return
	}
	if i >= dll.size/2 {
		oldNode := dll.tail
		for index := dll.size - 1; index > i; index-- {
			oldNode = oldNode.pre
		}
		dll.doInsert(oldNode, nd)
	} else {
		oldNode := dll.head
		for index := 0; index < i; index++ {
			oldNode = oldNode.next
		}
		dll.doInsert(oldNode, nd)
	}
}
func (dll *DoubleLinkedList[T]) doRemoveNode(node *DNode[T]) T {
	dll.size -= 1
	preNode := node.pre
	nextNode := node.next
	if preNode == nil {
		dll.head = nextNode
		return node.value
	}
	if nextNode == nil {
		dll.tail = node.pre
		return node.value
	}
	preNode.next = nextNode
	nextNode.pre = preNode
	return node.value
}
func (dll *DoubleLinkedList[T]) RemoveAt(i int) T { //Removes and returns the item in the ith node of the list
	if i < 0 || i >= dll.size {
		panic("给定的元素位置超界")
	}
	if i >= dll.size/2 {
		node := dll.tail
		for index := dll.size - 1; index > i; index-- {
			node = node.pre
		}
		return dll.doRemoveNode(node)
	} else {
		node := dll.head
		for index := 0; index < i; index++ {
			node = node.next
		}
		return dll.doRemoveNode(node)
	}
}

func (dll *DoubleLinkedList[T]) Append(item T) { //Creates and inserts item into the last node of the list
	if IsNil(item) {
		panic("不允许向列表插入空对象！")
	}
	nd := &DNode[T]{value: item, pre: nil, next: nil}
	if dll.size == 0 {
		dll.addNodeToEmpyList(nd)
		return
	}
	nd.pre = dll.tail
	dll.tail.next = nd
	dll.tail = nd
	dll.size += 1
}

// 双向链表获取制定位置的元素可以根据位置是靠近头节点还是尾结点来进行一些优化
func (dll DoubleLinkedList[T]) Get(i int) T { // Returns the node position containing item in the list
	if i < 0 || i >= dll.size || dll.size == 0 {
		panic("无法获取非法序号的节点")
	}
	if i <= dll.size/2 {
		index := 0
		nd := dll.head
		for {
			if index == i {
				return nd.value
			} else {
				nd = nd.next
				index++
			}
		}
	} else {
		index := dll.size - 1
		nd := dll.tail
		for {
			if index == i {
				return nd.value
			}
			nd = nd.pre
			index--
		}
	}
}
func (dll DoubleLinkedList[T]) Items() []T { //Returns a slice of all the items in the list
	var result []T
	if dll.size == 0 {
		return result
	}
	for nd := dll.head; nd.next != nil; nd = nd.next {
		result = append(result, nd.value)
	}
	return result
}
func TestDoubleLinkedList(t *testing.T) {
	TestSingleLinkedList(t)
	println("-------------------------------------------------")
	var idll List[int] = &DoubleLinkedList[int]{}
	idll.Append(1)
	idll.Append(2)
	idll.Append(3)
	i := Find(idll, 0)
	println(i)
	var sdll List[string] = &DoubleLinkedList[string]{}
	sdll.Append("One")
	sdll.Append("Two")
	sdll.Append("Three")
	i = Find(sdll, "One")
	println(i)
	println(sdll.Size())
	sdll.RemoveAt(1)
	println(sdll.Size())
	println(Find(sdll, "Two"))
	sdll.Insert(1, "Two")
	sdll.Insert(3, "Four")
	sdll.Insert(4, "Five")
	println(Find(sdll, "Two"))
	sdll.RemoveAt(3)
	println(Find(sdll, "Four"))
	println(Find(sdll, "Five"))

}
