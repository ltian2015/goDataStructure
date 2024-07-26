/*
*

本文件只涉及栈(Stack)、队列(Queue)、链表(List)和普通有根树(Tree)的数据结构和算法。
*
*/
package basic

import (
	"fmt"
	"reflect"
	"testing"
)

// !!! IsNil泛型函数判断给定的任何类型（any类型）值是否为nil。
// !!! golang中，所有类型都是语言所提供的基础类型做源类型或组合所衍生的，
// !!! 这些基本类型决定了被衍生类型的内存布局,也就决定其“零值”应该是nil还是0。
// !!! 因而，只要判断给定类型的值是否属于以下种类（kind），
// !!! 就可以通过该值调用IsZero（是否零值）判定其值是nil还是非nil。
// !!! 所有接口类型，包括interface{}，也就是any类型在内的接口零值,——nil最为特殊，用该值调用IsZero会抛出异常，但是该nil值的
// !!! kind是Invalid，因此可以用于判断是否为nil。（需要确定是否还有其他情况出现Invalid Kind的特殊值,但目前尚未发现）
func IsNil[T any](t T) bool {
	value := reflect.ValueOf(t)
	kind := value.Kind()
	switch kind {
	case reflect.Invalid:
		return true
	case reflect.Interface, reflect.Pointer, reflect.Chan,
		reflect.Func, reflect.UnsafePointer, reflect.Map, reflect.Slice:
		if value.IsZero() {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

func TestIsNil(t *testing.T) {
	var s string
	if s == "" {
		println("string类型的零值是\"\"")
	}
	var p *int

	if p == nil {
		println("f 为（==） nil")
	}
	if IsNil(p) {
		println("f 是（isNil） nil")
	}
	var f func(a string) string
	if f == nil {
		println("f 为（==） nil")
	}
	if IsNil(f) {
		println("f 是（isNil） nil")
	}
	var i any
	if i == nil {
		println("i 为（==） nil")
	}
	if IsNil(i) {
		println("i 是（isNil） nil")
	}
	var iv int
	if IsNil(iv) {
		println("it 是（isNil） nil")
	} else {
		fmt.Printf("%d  不是（isNil） nil\n", iv)
	}
	var strct = struct {
		id   int
		name string
	}{}
	if IsNil(strct) {
		println("strct 是（isNil） nil")
	} else {
		fmt.Printf("%v  不是（isNil） nil\n", strct)
	}
	empstrct := struct{}{}
	if IsNil(strct) {
		println("empstrct 是（isNil） nil")
	} else {
		fmt.Printf("%v  不是（isNil） nil\n", empstrct)
	}
}

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

type Node[T any] struct {
	value T
	next  *Node[T]
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
