package chapter1

/**
选择算法是一种在列表或数组中查找第 k 个最小（或最大）数字的算法。该数字称为 k 阶统计量。
它包括查找列表或数组中的最小、最大和中值元素的各种情况。选择算法通常有几种模式：
!!! 一：基于排序的选择算法
!!! 1.完全排序后再选择：这种算法对于在同一个列表或数组中反复多次进行不同的K阶统计量的选择是高效的。
!!! 2.无序的部分排序（Unordered Partial Sorting）：该算法排除前 k 个元素，其余元素按随机顺序排列，
!!!   这样就可以查找第 k 个最小（或最大）元素。可以先将列表按k个元素等分，分别排序，然后在分别与第一个k等分
!!!   进行比较和交换，使第一个k等分成为前k个最小（或最大）有序子序列。
**/
import (
	"fmt"
	"testing"
)

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
// !!! 本质上，这是一种部分排序选择法，只排出最小的前2个元素。这种算法就是把整个列表分成k(这里k=2）等分，
// !!! 每个等分进行排序，然后再与第一个等分进行比较和交换，这样第一个k等分就成了这个列表的前k个有序元素。
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
	for i := startIndex; i < len(s); i += 2 /**2等分**/ {
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

// ------------------------------------------------------------------------------
// 部分选择排序法，其思想就是排出前k个元素的顺序.这里的具体算法就是
// 把列表k等分，每一个k等分都与第一个k等分子集进行比较和交换，使第一个k等分始终保持是前k个元素的有序子集
func partialSelectionSort(s []int, k int) int {
	//!!!s1是已排序好的k个元素序列，相当于把s2的元素按照插入排序法插入到前k有序个元素中。
	insertAndSort := func(k int, s1, s2 []int) []int {
		for i := len(s2) - 1; i >= 0; i-- {
			if s2[i] < s1[k-1] { //比s1最大的小，则通过交换的方式插入到s1的最大位置处
				temp := s1[k-1]
				s1[k-1] = s2[i]
				s2[i] = temp
				//（试图）将s1的最大位置处新插入的元素前移，已保证顺序
				for j := k - 1; j >= 1; j-- {
					if s1[j] < s1[j-1] {
						swap(s1, j, j-1)
					} else {
						break
					}
				}
			}
		}
		return s1
	}
	s1 := s[0:k]
	quickSort(s1)
	l := len(s)
	for i := k; i < l; i += k {
		var s2 []int
		if i+k < l {
			s2 = s[i : i+k]
		} else {
			s2 = s[i:l]
		}
		s1 = insertAndSort(k, s1, s2)
	}
	return s1[k-1]
}
func TestPartialSelectionSort(t *testing.T) {
	s := []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10}
	fmt.Println(s)
	first := partialSelectionSort(s, 1)
	fmt.Printf("第1小： %d ", first)
	fmt.Println(s)
	s = []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10}
	second := partialSelectionSort(s, 2)
	fmt.Printf("第2小： %d ", second)
	fmt.Println(s)

	s = []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10}
	third := partialSelectionSort(s, 3)
	fmt.Printf("第3小： %d ", third)
	fmt.Println(s)

	s = []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10}
	fourth := partialSelectionSort(s, 4)
	fmt.Printf("第4小： %d ", fourth)
	fmt.Println(s)

	s = []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10}
	fifth := partialSelectionSort(s, 5)
	fmt.Printf("第5小： %d ", fifth)
	fmt.Println(s)

	s = []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10, 3, 4, 45}
	sixth := partialSelectionSort(s, 6)
	fmt.Printf("第6小： %d ", sixth)
	fmt.Println(s)

	s = []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10}
	seventh := partialSelectionSort(s, 7)
	fmt.Printf("第7小： %d ", seventh)
	fmt.Println(s)

	s = []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10}
	fmt.Println(s)
	eighth := partialSelectionSort(s, 8)
	fmt.Printf("第8小： %d ", eighth)
	fmt.Println(s)

	s = []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10}
	nineth := partialSelectionSort(s, 9)
	fmt.Printf("第9小： %d ", nineth)
	fmt.Println(s)

	s = []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10, 3, 4, 45}
	tenth := partialSelectionSort(s, 10)
	fmt.Printf("第10小： %d ", tenth)
	fmt.Println(s)

	s = []int{2, 34, 5, 6, 0, 8, 6, -1, 7, 9, 10}
	eleventh := partialSelectionSort(s, 11)
	fmt.Printf("第11小： %d ", eleventh)
	fmt.Println(s)

}
