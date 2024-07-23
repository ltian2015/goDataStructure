/*
*

本文件只涉及栈(Stack)、队列(Queue)、链表(List)和普通有根树(Tree)的数据结构和算法。
*
*/
package basic

import (
	"fmt"
	"testing"
)

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
	nameStack := SliceStack[string]{}
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
	intStack := SliceStack[int]{}
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
