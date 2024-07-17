package chapter1

import (
	"fmt"
	"testing"
)

//!!! indElementsWithSum找出一个包含 10 个元素的数组中元素的组合，其和等于 18。
//!!! findElementsWithSum 方法以递归方式尝试查找组合。每当和超出目标 k 时，它就会回溯：

func findElementsWithSum(arr [10]int, combinations [19]int, size int, k int,
	addValue int, l int, m int) int {
	var num int = 0
	if addValue > k {
		return -1
	}
	if addValue == k {
		num = num + 1
		var p int = 0
		for p = 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}
	var i int
	for i = l; i < size; i++ {
		//fmt.Println(" m", m)
		combinations[m] = l
		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l = l + 1
	}
	return num
}

// main method
func TestBackTracing(t *testing.T) {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var addedSum int = 18
	var combinations [19]int
	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
	// fmt.Println(check)//var check2 bool = findElement(arr,9)
	// fmt.Println(check2)
}
