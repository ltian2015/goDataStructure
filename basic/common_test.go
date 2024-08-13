package basic

import (
	"fmt"
	"reflect"
	"testing"
)

const SIZE = 10_000_000

type Node[T any] struct {
	value T
	next  *Node[T]
}

type DNode[T any] struct {
	value T
	pre   *DNode[T]
	next  *DNode[T]
}

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

// Map函数把slice容器中所有元素按给定转换方法转换为另外一种元素的方法，并返回转换后的新slice。
// s是源slice容器，mapFunc是转换方法
func Map[A, B any, S ~[]A](s S, mapFunc func(a A) B) []B {
	l := len(s)
	if l == 0 {
		return []B{}
	}
	result := make([]B, l, l)
	for i, a := range s {
		b := mapFunc(a)
		result[i] = b
	}
	return result
}

// LeftFold函数对slice容器以左折叠方式进行缩减（reduce）求值。
// s是被折叠的slice，
// foldFunc是对两个值进行折叠得到一个值的折叠操作函数
func LeftFold[A any, S ~[]A](s S, foldFunc func(a1, a2 A) A) A {
	var result A //!!! 此时，result值是类型A的零值
	if len(s) == 0 {
		return result
	}
	for _, a := range s {
		result = foldFunc(result, a)
	}
	return result
}
