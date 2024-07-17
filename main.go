package main

import (
	"cmp"
	"fmt"
)

func main() {
	fmt.Println("GO数据结构第一章")
	methods := linearCombinationExample(41, 5, 3)
	for _, method := range methods {
		fmt.Println("(", method.Cx, " , ", method.Cy, ")")
	}
	i, j := 2, 3
	a := less(i, j)
	println(a)
}

type Method struct {
	Cx int
	Cy int
}

func less[T cmp.Ordered](a, b T) bool {
	return a > b
}

// 线性组合的简单例子，求m个x和n个y的和为n的所有可能性组合。
func linearCombinationExample(n int, x, y int) (methods []Method) {
	methods = make([]Method, 0)
	if n <= 0 || x <= 0 || y <= 0 {
		return methods
	}
	step := max(x, y)
	isXlessY := step == y
	for i := 0; i <= n/step; i++ {
		if isXlessY {
			if (n-i*step)%x == 0 {
				method := Method{Cx: (n - i*step) / x, Cy: i}
				methods = append(methods, method)
			} else {
				continue
			}
		} else {
			if (n-i*step)%y == 0 {
				method := Method{Cx: i, Cy: (n - i*step) / y}
				methods = append(methods, method)
			} else {
				continue
			}
		}

	}
	return methods
}
