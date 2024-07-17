package chapter1

import (
	"fmt"
	"testing"
)

// 部分选择排序法，其思想就是排出前k（0-k-1）个元素的大小。这样，就要把
// 二者就交换(替换第k个元素)
func partialSelectionSort(s []int, k int) int {
	quickSort(s[0 : k-1])
	for i := k; i < len(s); i++ {
		if s[i] < s[k-1] {
			swap(s, i, k-1)
		}
	}
	return s[k-1]
}
func TestPartialSelectionSort(t *testing.T) {
	s := []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10}
	fmt.Println(s)
	second := partialSelectionSort(s, 2)
	third := partialSelectionSort(s, 3)
	fourth := partialSelectionSort(s, 4)
	fmt.Printf("第二小：%d, 第三小：%d 第四小：%d /n", second, third, fourth)

}

// /////////////////////////////下面是一些特殊值的选择，最小值、最大值/////////////////////////////
// SlecxtMin找到给定数组中的最小元素及其位置，需要进行n-1次比较。
func SelectMin(s []int) (min int, loc int) {
	if s == nil || len(s) == 0 {
		panic("无法在空数组中找到最小元素")
	}
	min = s[0]
	loc = 0
	for i := 1; i < len(s); i++ {
		if min > s[i] { // 当前的min被 s[i]被所战胜（比min还是小），那么就用s[i]的值作为min，并更新位置
			min = s[i]
			loc = i
		}
	}
	return min, loc
}
func TestSelectMin(t *testing.T) {
	s := []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10}
	fmt.Println(s)
	min, loc := SelectMin(s)
	fmt.Printf("最小值是:%d,位置为：%d", min, loc)
}
func SelectMinAndMax(s []int) (min, max int) {
	if s == nil || len(s) == 0 {
		panic("无法在空数组中找到最小元素和最大元素")
	}
	evenCount := len(s)%2 == 0
	const setp = 2
	var startIndex int = 0
	if evenCount {
		startIndex = 2
		if s[0] > s[1] {
			min, max = s[1], s[0]
		} else {
			min, max = s[0], s[1]
		}
	} else {
		startIndex = 1
		min, max = s[0], s[0]
	}
	for i := startIndex; i < len(s); i += 2 {
		if s[i] > s[i+1] {
			if s[i] > max {
				max = s[i]
			}
			if s[i+1] < min {
				min = s[i+1]
			}
		} else {
			if s[i+1] > max {
				max = s[i+1]
			}
			if s[i] < min {
				min = s[i]
			}
		}
	}
	return min, max
}

func TestSelectMinAndMax(t *testing.T) {
	s := []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10, 122, 56, -3, 324}
	fmt.Println(s)
	min, max := SelectMinAndMax(s)
	fmt.Printf("最小值是:%d,最大值是：%d", min, max)
}

// SelectMinAndSecondMinBasic是最基本（易于理解）的选择最小值和次小值的算法，
// 这种算法的想法和同时找出最大值最小值一样，就是一次拿出两个元素与当前的最小值及次小值进行比较，
// 相当于在4个元素中找出最小值与次小值，这就需要7次比较，大概的比较次数为： n/2 * 7。
func SelectMinAndSecondMinBasic(s []int) (min, secondMin int) {
	if s == nil || len(s) <= 2 {
		panic("无法在空数组,或者元素过少的数组中找到最小与次小的元素")
	}
	evenCount := len(s)%2 == 0
	const setp = 2
	var startIndex int = 0
	if evenCount {
		startIndex = 2
		if s[0] > s[1] {
			secondMin, min = s[1], s[0]
		} else {
			secondMin, min = s[0], s[1]
		}
	} else {
		startIndex = 1
		min, secondMin = s[0], s[0]
	}
	//糟糕情况下，下面的循环进行了(n-1）/2 * 7次元素比较，即，3.5（n-1）次比较
	//理论上还存在向 n+lgn-2  次比较的优化算法。
	for i := startIndex; i < len(s); i += 2 {
		if s[i] > s[i+1] { //s[i+1]小   //!!! 为了进入分支，进行了1次比较
			if s[i+1] < min && s[i] < min { //!!!进行了两次比较
				min = s[i+1]
				secondMin = s[i]
			} else if s[i+1] < min && s[i] > secondMin { //!!!进行了两次比较
				secondMin = min
				min = s[i+1]
			} else if s[i+1] > min && s[i+1] < secondMin { //!!! 进行了两次比较
				secondMin = s[i+1]
			}
		} else { //s[i]小
			if s[i] < min && s[i+1] < min {
				min = s[i]
				secondMin = s[i+1]
			} else if s[i] < min && s[i+1] > min {
				secondMin = min
				min = s[i]
			} else if s[i] > min && s[i] < secondMin {
				secondMin = s[i]
			}
		}
	}
	return min, secondMin
}
func TestSelectMinAndSecondMin(t *testing.T) {
	s := []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10, 122, 56, -3, 324}
	fmt.Println(s)
	min, max := SelectMinAndSecondMinBasic(s)
	fmt.Printf("最小值是:%d, 次小值是：%d", min, max)
}
