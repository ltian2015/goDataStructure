package basic

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"
	"time"
)

type Deque[T any] interface {
	InsertFront(item T)
	InsertBack(item T)
	RemoveFirst() T
	RemoveLast() T
	First() T
	Last() T
	IsEmpty() bool
	Size() int
}

type SliceDeque[T any] struct {
	items []T
}

func (sdq *SliceDeque[T]) InsertFront(item T) {
	if IsNil(item) {
		panic("空值不允许插入到队列")
	}
	sdq.items = append(sdq.items, item)
	l := len(sdq.items)
	for i := l - 1; i > 0; i-- {
		sdq.items[i] = sdq.items[i-1]
	}
	sdq.items[0] = item
}
func (sdq *SliceDeque[T]) InsertBack(item T) {
	if IsNil(item) {
		panic("空值不允许插入到队列")
	}
	sdq.items = append(sdq.items, item)
}
func (sdq *SliceDeque[T]) RemoveFirst() T {
	if len(sdq.items) == 0 {
		panic("队列已空，不能再删除元素")
	}
	result := sdq.items[0]
	sdq.items = sdq.items[1:]
	return result

}
func (sdq *SliceDeque[T]) RemoveLast() T {
	l := len(sdq.items)
	if l == 0 {
		panic("队列已空，不能再删除元素")
	}
	result := sdq.items[l-1]
	sdq.items = sdq.items[0 : l-1]
	return result

}
func (sdq *SliceDeque[T]) First() T {
	if len(sdq.items) == 0 {
		panic("队列已空，无法读取第一个元素")
	}
	return sdq.items[0]
}
func (sdq *SliceDeque[T]) Last() T {
	if len(sdq.items) == 0 {
		panic("队列已空，无法读取最后一个元素")
	}
	l := len(sdq.items)
	return sdq.items[l-1]
}
func (sdq *SliceDeque[T]) IsEmpty() bool {
	return len(sdq.items) == 0

}
func (sdq *SliceDeque[T]) Size() int {
	return len(sdq.items)
}

func TestDeque(t *testing.T) {
	var myDeque Deque[int] = &SliceDeque[int]{}
	myDeque.InsertFront(5)
	myDeque.InsertBack(10)
	myDeque.InsertFront(2)
	myDeque.InsertBack(12) // 2 5 10 12
	fmt.Println("myDeque.First() = ", myDeque.First())
	fmt.Println("myDeque.Last() = ", myDeque.Last())
	myDeque.RemoveLast()
	myDeque.RemoveFirst()
	fmt.Println("myDeque.First() = ", myDeque.First())
	fmt.Println("myDeque.Last() = ", myDeque.Last())
}

// 使用暴力算法求解元素个数为k的所有连续子数组中每个子数组中最大元素所组成的数组
// 比如， input []int{9, 1, 1, 0, 0, 0, 1, 0, 6, 8} ,k=3  则
// 子数组{9,1,1} 最大值为9
// 子数组{1,1，0}最大值为1
// 子数组{1,0，0}最大值为1
// 子数组{0,0，0}最大值为0
// 子数组{0,0，1}最大值为1
// 子数组{0,1，0}最大值为1
// 子数组{1,0，6}最大值为6
// 子数组{0,6，8}最大值为8
// 最终的输出数组为{9,1,1,0,1,1,6,8}
// 暴力算法的逻辑很简单，算法的时间复杂度约为(n*K),当K和n比较大时，就比较耗费时间
// 比如，当n为1百万，而k为1000的时候，就要循环比较约10亿次
// !!! 这种处理的也可以理解为，一次让每个元素与其后的K-1个元素进行比较，如果不是最大的就淘汰。
// !!! 如果后续元素不足K-1个，那就直接淘汰。
func MaxSubarryBruteForce(input []int, k int) (output []int) {
	for i := 0; i <= len(input)-k; i++ { //i表示每个子数组的在大数组中的起始坐标
		subArray := input[i : i+k]
		max := subArray[0]
		for j := 1; j < k; j++ { //j代表子数组的起始
			if subArray[j] > max {
				max = subArray[j]
			}
		}
		output = append(output, max)
	}
	return output
}
func TestMaxSubarry(t *testing.T) {
	input := []int{3, 1, 6, 4, 2, 10, 5, 9}
	output := MaxSubarryBruteForce(input, 3)
	fmt.Println("Output = ", output)
	input = []int{9, 1, 1, 0, 0, 0, 1, 0, 6, 8}
	output = MaxSubarryBruteForce(input, 3)
	fmt.Println("Output = ", output)
}

//
//MaxSubarryBruteForceImprove是对暴力算法的一种改进，
// 在随机一百万元素的数组中，对个数为1000的子数组求最大值列表的性能要比暴力算法提高5倍。
//其原理是：对于任何一个子数组[i,i+k]中的最大的元素,这里，0<=i<=len(input)，假定其在原始数组的序号为j
// i<=j<i+k，这意味着如果j在后续子数组[x,x+k]序号范围中，那么只需将序号从j+1到x+k-1的元素
// 序号为j的元素进行比较，就可再次确定子数组中最大元素的序号，这样不断迭代，直到所有子数组都遍历完毕。
//

func MaxSubarryBruteForceImprove(input []int, k int) (output []int) {

	var baseIndex int = 0
	//var max int = input[baseIndex]
	var maxIndex int = -1
	var maxRangeIndex = -1
	for baseIndex = 0; baseIndex <= len(input)-k; baseIndex++ {
		if maxIndex < baseIndex {
			//max = input[baseIndex]
			maxIndex = baseIndex
			maxRangeIndex = baseIndex
		}
		for i := maxRangeIndex - baseIndex + 1; i < k; i++ {
			if input[baseIndex+i] >= input[maxIndex] {
				maxIndex = baseIndex + i
				maxRangeIndex = baseIndex + k - 1
			}
		}
		output = append(output, input[maxIndex])
	}

	return output
}

// MaxSubarrayUsingDeque是一种与输入规模为线性关系的算法，
// 在随机一百万元素的数组中，对个数为1000的子数组求最大值列表的性能要比改进后的暴力算法提高30倍。
// 比暴力算法性能提高150倍。
// !!! 其原理是： 逐个访问元素，用一个“淘汰式”队列（保存元素序号即可）来保持当前元素到之前的K个元素
// !!! 按照从大到小顺序的排队。 也就是，每访问一个新元素就对队列进行一次从后先前的淘汰式刷新，这样就
// !!! 能始终使队列保持为前K个元素中，以最大元素序号为头部，大小依次排列的有序队列。当然，队列里元素
// !!! 数量取决于较大最大元素出现的先后，比如，一个窗口内的最大元素最后一个出现，就会把前面所有元素都淘汰。
// !!! 在移动元素访问序号i的时候，窗口[i-k+1,i) 也会随着移动，这里k<=i<len(input）。
// !!! 这样窗口[i-k+1,i)就一定会从队列头部寻找到自己的最大元素，如果队头的元素已经是窗口的左边界（i-k+1），
// !!! 添加为结果元素之后,就要从队列中移除,因为它不再属于后续的范围。

// !!! 通过这个范例，我们可以了解到，双向队列可以从队尾或队头进行访问和元素的增删操作，无法在其他位置进行
// !!! 增删操作，因此，比较适合完成这种“淘汰式”的排队——从队尾或队头向前或先后淘汰（不断移除）。

func MaxSubarrayUsingDeque(input []int, k int) (output []int) {
	deque := SliceDeque[int]{}
	var index int
	// 第一个元素窗口 [0,k]
	for index = 0; index < k; index++ {
		//用当前元素淘汰队尾元素（比之小，或同样大小但出现较早的元素），直到无法淘汰为止
		for !deque.IsEmpty() && input[index] >= input[deque.Last()] {
			deque.RemoveLast() //淘汰比当前元素小的所有队尾元素
		}
		deque.InsertBack(index) //经过淘汰后，将当前元素作为入选队列的元素插入到队列中
	}
	output = append(output, input[deque.First()])
	for ; index < len(input); index++ {
		//如果队头元素对当前窗口[index-k+1,index)失效，则移除
		if !deque.IsEmpty() && deque.First() <= index-k {
			deque.RemoveFirst()
		}
		// 不断用当前元素淘汰队尾元素，直到无法淘汰为止
		for !deque.IsEmpty() && input[index] >= input[deque.Last()] {
			deque.RemoveLast()
		}
		deque.InsertBack(index)
		//将新队头作为当前窗口[index-k+1,index)的最大值追加到结果中
		output = append(output, input[deque.First()])
	}
	return output
}

// MaxSubarrayUsingDeque2与MaxSubarrayUsingDeque思路一样，
// 是书中原始代码，不易阅读.
func MaxSubarrayUsingDeque2(input []int, k int) (output []int) {
	deque := SliceDeque[int]{}
	var index int
	// First window
	for index = 0; index < k; index++ {
		for {
			if deque.IsEmpty() || input[index] < input[deque.Last()] {
				break
			}
			deque.RemoveLast()
		}
		deque.InsertBack(index)
	}
	for ; index < len(input); index++ {
		output = append(output, input[deque.First()])
		// Remove elements out of the window
		for {
			if deque.IsEmpty() || deque.First() > index-k {
				break
			}
			deque.RemoveFirst()
		}
		// Remove values smaller than the element currently being added
		for {
			if deque.IsEmpty() || input[index] < input[deque.Last()] {
				break
			}
			deque.RemoveLast()
		}
		deque.InsertBack(index)
	}
	output = append(output, input[deque.First()])
	return output
}
func TestMaxSubarryDequeVerify(t *testing.T) {
	input := []int{}
	for i := 0; i < SIZE; i++ {
		input = append(input, rand.IntN(1000))
	}
	output1 := MaxSubarrayUsingDeque(input, 10000)
	output2 := MaxSubarrayUsingDeque2(input, 10000)
	println(slices.Equal(output1, output2))
}

func TestMaxSubarryDeque(t *testing.T) {
	var input []int
	var output []int
	input = []int{3, 1, 6, 4, 2, 10, 5, 9, 8, 5, 7, 9, 2, 1, 3, 6}
	output = MaxSubarryBruteForce(input, 5)
	fmt.Println("Output = ", output)
	output = MaxSubarryBruteForceImprove(input, 5)
	fmt.Println("Output = ", output)
	output = MaxSubarrayUsingDeque(input, 5)
	fmt.Println("Output = ", output)

	input = []int{9, 1, 1, 0, 0, 0, 1, 0, 6, 8}
	output = MaxSubarryBruteForce(input, 3)
	fmt.Println("Output = ", output)
	output = MaxSubarryBruteForceImprove(input, 3)
	fmt.Println("Output = ", output)
	output = MaxSubarrayUsingDeque(input, 3)
	fmt.Println("Output = ", output)

	input = []int{9, 1, 1, 0, 0, 0, 1, 0, 6, 8, 11, 7, 3, 4, 6, 4, 5, 4, 1, 1, 0}
	output = MaxSubarryBruteForce(input, 3)
	fmt.Println("Output = ", output)
	output = MaxSubarryBruteForceImprove(input, 3)
	fmt.Println("Output = ", output)
	output = MaxSubarrayUsingDeque(input, 3)
	fmt.Println("Output = ", output)
}

func TestComparePerformance(t *testing.T) {
	input := []int{9, 1, 1, 0, 0, 0, 1, 0, 6, 8}
	output1 := MaxSubarryBruteForce(input, 3)
	fmt.Println("Output = ", output1)
	output2 := MaxSubarrayUsingDeque(input, 3)
	fmt.Println("Output = ", output2)
	// Benchmark performance of two algorithms
	input = []int{}
	//!!! SIZE  大小为1百万
	for i := 0; i < SIZE; i++ {
		input = append(input, rand.IntN(1000))
	}
	//!!! 注意，算法相同，写法稍有差异的两个函数，对于同一个大规模的输入序列，
	//!!! 谁后执行,谁的性能稍微占优，因为现代CPU有分支预测功能，后执行函数CPU分支预测效果更好.
	start := time.Now()
	MaxSubarrayUsingDeque2(input, 10000)
	elapsed := time.Since(start)
	fmt.Println("Using Deque2: ", elapsed)

	start = time.Now()
	MaxSubarrayUsingDeque(input, 10000)
	elapsed = time.Since(start)
	fmt.Println("Using Deque: ", elapsed)

	start = time.Now()
	MaxSubarryBruteForceImprove(input, 10000)
	elapsed = time.Since(start)
	fmt.Println("Using MaxSubarry Brute Force Improve: ", elapsed)

	start = time.Now()
	MaxSubarryBruteForce(input, 10000)
	elapsed = time.Since(start)
	fmt.Println("Using Brute Force: ", elapsed)

}
