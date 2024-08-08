/*
*

本文件只涉及栈(Stack)、队列(Queue)、链表(List)和普通有根树(Tree)的数据结构和算法。
*
*/
package basic

import (
	"fmt"
	"testing"
	"time"
)

type StackPanic string

const PopEmptyStack StackPanic = "空栈弹出"
const TopEmptyStack StackPanic = "读取空栈"
const PushNilValue StackPanic = "空值入栈"

// Stack泛型接口提取了所有栈类型的共同性操作，及对操作的约定。
type Stack[T any] interface {
	//将一个元素压入栈，如果该元素的“零值”是nil,则不允许入栈，会抛出值为PushNilValue的 panic。
	Push(item T)
	//弹出栈中元素，如果栈已经为空，则会抛出值为PopEmptyStack的panic。
	Pop() T
	//读取栈顶端元素，如果栈已经为空，则会抛出值为TopEmptyStack的panic
	Top() T
	IsEmpty() bool
}

/*
*
SliceStack要求操作的元素的类型都是comparable的子类型，
这是因为为了要防止将类型的“零值”入栈，因此，需要将操作的元素与
"零值"进行比较，故而要求是comparable的子类型。
*
*/
type SliceStack[T comparable] struct {
	items []T
}

func getZero[T comparable]() T {
	var zeroValue T
	return zeroValue
}

// !!!这个实现的阻止了“零值”的入栈，对于空值为非nil的类型来说不合理
func (stack *SliceStack[T]) Push(item T) {
	if item != getZero[T]() {
		stack.items = append(stack.items, item)
	}
}

// !!! 这个实现认为当栈为空的时候，Pop操作返回类型的“零值”，这对于空值为非nil的类型来说不合理
func (stack *SliceStack[T]) Pop() T {
	var result T
	length := len(stack.items)
	if length > 0 {
		result = stack.items[length-1]
		stack.items = stack.items[:length-1]
	}
	return result
}

// !!! 这个实现认为当栈为空的时候，Top操作返回类型的“零值”，这对于空值为非nil的类型来说不合理
func (stack SliceStack[T]) Top() T {
	var result T
	length := len(stack.items)
	if length > 0 {
		result = stack.items[length-1]
	}
	return result
}
func (stack SliceStack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

func TestSliceStack(t *testing.T) {
	// Create a stack of names
	var nameStack Stack[string]
	nameStack = &SliceStack[string]{}

	nameStack.Push("Zachary")
	nameStack.Push("Adolf")
	topOfStack := nameStack.Top()
	if topOfStack != getZero[string]() {
		fmt.Printf("\nTop of stack is %s", topOfStack)
	}
	poppedFromStack := nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	poppedFromStack = nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}
	poppedFromStack = nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	poppedFromStack = nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	// Create a stack of integers
	var intStack Stack[int] = &SliceStack[int]{}

	intStack.Push(5)
	intStack.Push(10)
	intStack.Push(0) // Problem since 0 is the zero
	// value for int
	top := intStack.Top()
	if top != getZero[int]() {
		fmt.Printf("\nValue on top of intStack is %d", top)
	}
	popFromStack := intStack.Pop()
	if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
	popFromStack = intStack.Pop()
	if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
	popFromStack = intStack.Pop()
	if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
}

type SliceStackAny[T any] struct {
	items []T
}

func (stack *SliceStackAny[T]) Push(item T) {
	if IsNil(item) {
		panic(PushNilValue)
	} else {
		stack.items = append(stack.items, item)
	}
}
func (stack *SliceStackAny[T]) Pop() T {
	length := len(stack.items)
	if length == 0 {
		panic(PopEmptyStack)
	} else {
		item := stack.items[length-1]
		stack.items = stack.items[:length-1]
		return item
	}
}
func (stack *SliceStackAny[T]) Top() T {
	length := len(stack.items)
	if length == 0 {
		panic(TopEmptyStack)
	} else {
		return stack.items[length-1]
	}
}

func (stack SliceStackAny[T]) IsEmpty() bool {
	return len(stack.items) == 0
}
func NewSliceStackAny[T any]() SliceStackAny[T] {
	return SliceStackAny[T]{}
}
func TestSliceStackAny(t *testing.T) {
	var add func(a, b int) int = func(a, b int) int { return a + b }
	var sub func(a, b int) int = func(a, b int) int { return a - b }
	//var mul func(a, b int) int
	var stk = NewSliceStackAny[func(int, int) int]()
	var opstack Stack[func(int, int) int] = &stk
	opstack.Push(add)
	opstack.Push(sub)
	println("working ok....")
	//	opstack.Push(mul)
	op := opstack.Pop()
	println(op(3, 6))

	nameStack := NewSliceStackAny[string]()
	nameStack.Push("郑健")
	nameStack.Push("刘飞")
	nameStack.Push("金正日")
	println(nameStack.Top())
	println(nameStack.Pop())
	println(nameStack.Pop())
	println(nameStack.Pop())
	println(nameStack.Pop())
}

type NodeStack[T any] struct {
	first *Node[T]
}

func (stack *NodeStack[T]) Push(item T) {

	if IsNil(item) {
		panic(PushNilValue)
	} else {
		nd := &Node[T]{value: item, next: nil}
		nd.next = stack.first
		stack.first = nd
	}
}
func (stack *NodeStack[T]) Pop() T {
	if stack.first == nil {
		panic(PopEmptyStack)
	} else {
		nd := stack.first
		stack.first = nd.next
		return nd.value
	}
}
func (stack *NodeStack[T]) Top() T {

	if stack.first == nil {
		panic(TopEmptyStack)
	} else {
		return stack.first.value
	}
}

func (stack NodeStack[T]) IsEmpty() bool {
	return stack.first == nil
}

func TestNodeStack(t *testing.T) {
	/**
	var add func(a, b int) int = func(a, b int) int { return a + b }
	var sub func(a, b int) int = func(a, b int) int { return a - b }
	//var mul func(a, b int) int
	var stk = NodeStack[func(int, int) int]{}
	var opstack Stack[func(int, int) int] = &stk
	opstack.Push(add)
	opstack.Push(sub)
	println("working ok....")
	//	opstack.Push(mul)
	op := opstack.Pop()
	println(op(3, 6))

	nameStack := NodeStack[string]{}
	nameStack.Push("郑健")
	nameStack.Push("刘飞")
	nameStack.Push("金正日")
	println(nameStack.Top())
	println(nameStack.Pop())
	println(nameStack.Pop())
	println(nameStack.Pop())
	//println(nameStack.Pop())
	**/
	anyStack := NodeStack[any]{}
	//var i interface{}
	//anyStack.Push(i)
	//var a any
	//anyStack.Push(a)
	//anyStack.Push(nil)
	fmt.Printf("nil 的类型是：%T\n", nil)
	var null any
	fmt.Printf("null 的类型是：%T\n", null)
	var nullStack Stack[string] = &NodeStack[string]{}
	fmt.Printf("nullStack 的类型是：%T\n", nullStack)
	anyStack.Push(nullStack)
}

func TestStackPerformance(t *testing.T) {
	nodeStack := NodeStack[int]{}
	sliceStack := SliceStackAny[int]{}
	// Benchmark nodeStack
	start := time.Now()
	for i := 0; i < SIZE; i++ {
		nodeStack.Push(i)
	}
	elapsed := time.Since(start)
	fmt.Println("\nTime for 10 million Push() operations on nodeStack: ", elapsed)
	start = time.Now()
	for i := 0; i < SIZE; i++ {
		nodeStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Pop() operations on nodeStack: ", elapsed)
	// Benchmark sliceStack
	start = time.Now()
	for i := 0; i < SIZE; i++ {
		sliceStack.Push(i)
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Push() operations on sliceStack: ", elapsed)
	start = time.Now()
	for i := 0; i < SIZE; i++ {
		sliceStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Pop() operations on sliceStack: ", elapsed)
}

type Data struct {
	id   int
	name string
}

func TestSlice(t *testing.T) {
	s := []Data{}
	d1 := Data{1, "lan"}
	s = append(s, d1)
	d1.id = 2
	fmt.Printf("d1 is %v s is %v", d1, s)
}
