package chapter1

import (
	"container/list"
	"datastructure/chapter1/myheap"
	"fmt"
	"testing"
	"time"
)

func TestUseList(t *testing.T) {
	useList()
}

// list.List是一个双向链表，该类型的”零值“是一个可用的空链表
func useList() {

	var intList list.List
	intList.PushBack(11)
	intList.PushBack(12)  //想列表的尾部加入一个元素
	intList.PushFront(10) //向列表的头部加入一个元素
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}

// 在GO语言中，函数的结果以元组方式呈现
func powerSeriers(a int) (int, int) {
	square := a * a
	cube := square * a
	return square, cube
}

type Node struct {
	id   int
	name string
}

func TestUseTuple(t *testing.T) {
	square, cube := powerSeriers(5)
	fmt.Println("square:", square, " cube:", cube)
	var node1 = Node{id: 1, name: "lant"}
	var node2 = node1 //Golang 中，赋值就是对值的拷贝。
	node2.id = 2
	fmt.Println("node1 id :", node1.id, " node2 id :", node2.id)
	var t1 time.Time = time.Now()

	println(t1.Format("YYYY-MM-dd:hh:mm:ss"))
	println(t1.Compare(time.Now()))
}

// ----------------------golang容器（container）库提供的“堆（heap）”数据结构----------------------------------
/**
// 包heap为任何实现heap.Interface的类型提供堆操作。
   堆是一棵完全二叉树，其特点是每个节点都是其子树中的最小值（或最大值）节点。
   树中的最小元素是根，索引为 0。
   堆是实现优先级队列的常用方式。
   要构建优先级队列，请实现 Heap 接口，并以（负）优先级作为 Less 方法的排序，
   这样 Push 会添加项目，而 Pop 会从队列中删除优先级最高的项目。
   示例包括这样的实现；
   文件 example_pq_test.go 具有完整的源代码。
**/

type IntegerHeap []int

// -------以下5个方法是"container/heap"包中，heap接口对实现者的要求
func (iheap IntegerHeap) Len() int {
	return len(iheap)
}
func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}
func (iheap IntegerHeap) Swap(i, j int) {
	temp := iheap[i]
	iheap[i] = iheap[j]
	iheap[j] = temp
}
func (iheap *IntegerHeap) Push(heapIntf any) {
	*iheap = append(*iheap, heapIntf.(int))
}

func (iheap *IntegerHeap) Pop() any {
	size := len(*iheap)
	if size == 0 {
		return nil
	}
	previousHeap := *iheap
	poped := previousHeap[size-1]
	*iheap = previousHeap[0 : size-1]
	return poped
}
func TestHeap(t *testing.T) {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	myheap.Init(intHeap)
	myheap.Push(intHeap, 2)
	myheap.Push(intHeap, 10)
	myheap.Push(intHeap, 3)
	fmt.Printf("最小的数字: %d\n", (*intHeap)[0])

	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", myheap.Pop(intHeap))
	}

}
