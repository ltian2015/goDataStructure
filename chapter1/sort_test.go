package chapter1

/**
在计算机科学中，选择算法是一种用于在有序值（例如数字）集合中查找第 kc个最小值的算法。它找到的值称为
k阶统计量（kth order statistic）。作为特殊情况，选择包括在集合中查找最小、中值和最大元素的问题。
选择算法包括快速选择和中位数算法。当应用于n值的集合时，这些算法采用线性时间。如果是需要从一个数列中
多次选择元素，那么先进行排序在进行选择就比较划算。
**/
import (
	"fmt"
	"math/rand/v2"
	"testing"
)

const MaxUint = ^uint(0)

func getRandomNnumbers(size int, rng int) []int {

	result := make([]int, size, size+10)

	for i := 0; i < size; i++ {
		result[i] = rand.IntN(rng)
	}
	return result
}

func TestGetRandomNumbers(t *testing.T) {
	s := getRandomNnumbers(20, 100)
	fmt.Println(s)

}

func TestSelectionSort(t *testing.T) {
	s := getRandomNnumbers(30, 100)
	//s := []int{37, 0, 79, 71, 62, 91, 35, 47, 44, 19, 89, 38, 99, 57, 87, 56, 45, 25, 82, 88}
	fmt.Println(s)
	SelectionSort(s)
	fmt.Println(s)
}

// !!! SelectionSort对给定的切片进行选择排序
// !!! 选择排序的思路遍历切片全部数据，找到最小值元素作为第一个元素(与第一个元素交换)，
// !!! 然后再找次最小值作为第二个元素， 依次类推，完成排序
func SelectionSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] <= s[i] {
				//!!!将i处元素j处元素交换,使得i处元素仍旧当前最小
				temp := s[i]
				s[i] = s[j]
				s[j] = temp
			}
		}
	}
}
func TestBubbleSort(t *testing.T) {

}

func TestQuickSort(t *testing.T) {
	s := []int{37, 0, 79, 71, 62, 91, 35} //, 47, 44, 19, 89, 38, 99, 57, 87, 76, 56, 45, 25, 82, 88}
	fmt.Println(s)
	quickSort(s)
	fmt.Println(s)
}

// !!! quickSort（快速排序）是一种就地排序算法。其主要思想是找到任何一个元素作为基元，把所有比该元素小的元素都放在
// !!! 其左侧，把所有比该元素大的元素都放在其右侧，这样在对左右两侧的子数组进行同样的操作，如此
// !!! 反复递归下去，就会完成排序。
// !!! 在下面的算法中，为了便于操作，选取了最后一个元素作为基元。
func quickSort(s []int) {
	//!!! 下面的两种情况达到了不停与基元进行比较，形成左（小于）右（大于）两个子组数组的递归操作的“触底条件”
	l := len(s)
	if l == 1 || l == 0 {
		return
	}
	if l == 2 {
		if s[0] > s[1] {
			swap(s, 0, 1)
		}
		return
	}
	//!!! 选取最后一个元素作为基元，那么如果最后一个元素比其前一个元素小，那么,
	//!!! 二者就交换位置（这样就保证）,如果前一个元素小于或等于该元素，那么，
	//!!! 就把这个证明比之小的元素与数组中第一个未与之相比的元素交换（使之变为基元元素的前一个元素），
	//!!! 再重复上述操作，直至基元元素的序号与第一个未与之相比较的元素位置相同（相遇）就完成了左右
	//!!! 两侧数组的整理。

	baseIndex := l - 1                             //!!! 选取数组的最后一个元素基元
	firstUnComparedElementIndex := 0               //!!! 数组中第一个未与基元比较的元素位置,从0开始。
	for baseIndex != firstUnComparedElementIndex { //二者相遇
		BaseData := s[baseIndex]
		if s[baseIndex-1] > BaseData { //如果基元比前一个元素小，二者就交换位置
			swap(s, baseIndex-1, baseIndex)
			baseIndex = baseIndex - 1
		} else { //否则，就将这个小于基元的前一个元素与第一个未与基元比较的元素进行交换。
			swap(s, baseIndex-1, firstUnComparedElementIndex)
			firstUnComparedElementIndex += 1
		}
	}
	//对分治后的子数组继续进行同样的分治
	quickSort(s[:baseIndex])
	quickSort(s[baseIndex:])
}
func TestHeapSort(t *testing.T) {
	s := []int{37, 0, 79, 71, 62, 91, 35, 47, 44, 19, 89, 38, 99, 57, 87, 76, 56, 45, 25, 82, 88}
	fmt.Println(s)
	heapSort(s)
	fmt.Println(s)
}

// !!! heapSort（堆排序）是一种就地排序算法，主要是把数组看作了，完整二叉树，利用完整二叉树的性质
// !!! 进行排序，有关完整二叉树的知识详见 file://../myheap/数据结构学习-二叉树.pdf */
// !!! 核心思想就是，把被排序的数组看作完整二叉树，利用完整二叉其性质的将最小的元素交换到顶点，也就是
// !!! 序号为0的元素。然后，再把不包含0号元素的数组看作完整二叉树，继续重复上述交换操作，直到数组中的元素个
// !!! 数小于或等于2。 利用完整二叉树性质将最小的元素交换到定点的思路是，完整二叉树的最后一棵最小二叉数的顶点
// !!! 元素位置是 n/2,从该位置向前的元素依次是同级子树顶点和上级子树顶点，直至整个二叉树的顶点（位置为0的元素），
// !!! 这样，依次向前循环，把每一级子树的最小元素交换到顶点，那么整个二叉树的顶点（位置为0的元素）就是最小元素。
func heapSort(s []int) {
	//!!!当数组的长度小于或等于二的时候就达到了递归的“触底条件”，此时的排序操作十分简单。
	l := len(s)
	if l <= 2 {
		if l <= 1 {
			return
		}
		if s[0] > s[1] {
			swap(s, 0, 1)
		}
		return
	}
	//!!! 当数组元素个数大于2的时候，就处于递归条件。
	for i := l / 2; i >= 0; i-- { //l/2是最后的最小子树的顶点元素位置
		root := s[i]
		if 2*i >= l { //如果2*i>=l表明没有该最小子树没有子节点，无需操作
			continue
		}
		left := s[2*i] //找到给定二叉树的左子节点，如果比顶点小，二者就交换
		if root > left {
			swap(s, i, 2*i)
			root = left
		}
		if 2*i+1 < l { //找到给定二叉树的右子节点（如果存在的话），如果比顶点小，二者就交换
			right := s[2*i+1]
			if root > right {
				swap(s, i, 2*i+1)
			}
		}
	} //此循环结束后，二叉树的顶点变为最小值，也就是数组第一个元素一定是整个数组的最小值。
	s = s[1:] //将排除第一个元素的数组视图作为求解的数组，继续迭代
	heapSort(s)
}

// !!! swap函数对给定数组中的两个元素位置进行元素的交换操作
func swap(s []int, i, j int) {
	d := s[j]
	s[j] = s[i]
	s[i] = d
}

func TestInsertionSort(t *testing.T) {
	s := getRandomNnumbers(30, 100)
	//s := []int{37, 0, 79, 71, 62, 91, 35, 47, 44, 19, 89, 38, 99, 57, 87, 76, 56, 45, 25, 82, 88}
	fmt.Println(s)
	insertionSort(s)
	fmt.Println(s)
}

// !!! 插入排序就像对你手中的扑克牌进行排序。
// !!! 您将卡片分为两组：已排序的卡片和未排序的卡片。然后，从未排序的组中挑选一张卡，放在已排序的组中的最后位置，
// !!! 此时，从排序组的最后一个牌（待定位置的牌）作为当前处理的牌，与前一张牌比较，
// !!! 如果小于前一张牌，那么二者就交换顺序。再将当前牌的位置减1（前移，这样，当前牌还是待定位置的牌)，
// !!! 重复与前一张牌比较，直到当前的牌不再小于前面的牌为止。

func insertionSort(s []int) {
	l := len(s)
	if l <= 1 {
		return
	}
	//!!! 初始情况下认为s[0]是已经排好的，故而从i=1处开始对未排序的元素进行插入排序处理。
	//!!! i表示当前处理的尚未排序的元素位置，
	for i := 1; i < l; i++ {
		//!!!由于s[0:i]已经排好序，从后向前处理sub，只要当前的元素比前一个元素小，二者就交换顺序
		sub := s[0 : i+1] //s[0:i+1]不包括第i+1个元素，但包括第i个元素

		for j := len(sub) - 1; j > 0 && sub[j] < sub[j-1]; j-- {
			temp := sub[j-1]
			sub[j-1] = sub[j]
			sub[j] = temp
		}
	}
}

// !!!插入排序的原理是总是将“未排序的元素”插入到“已排好序”的集合中。
// !!!该算法原理在现实生活中就是一边摸扑克牌，一边排序的算法。
// !!! 这是一种不太好的方式，增加了移动操作步骤。
func insertionSort2(s []int) {
	l := len(s)
	if l <= 1 {
		return
	}
	//!!! 初始情况下认为s[0]是已经排好的，故而从i=1处开始对未排序的元素进行插入排序处理。
	//!!! i表示当前处理的尚未排序的元素位置，
	for i := 1; i < l; i++ {
		unSortNum := s[i] //unSortNum表示未排序的元素
		sub := s[0 : i+1] //
		for j := 0; j < i; j++ {
			if unSortNum <= s[j] {
				moveBackOneElmentThenInsert(sub, j, unSortNum)
				break
			}
		}
	}
}

// !!! moveBackOneStepThenInsert将给定的切片s中位于位置的i元素依次向后
// !!!  移动一位，挤掉最后一个元素,并将位置i处的元素替换为d.
// !!! 该方法服务于插入排序法insertionSort
func moveBackOneElmentThenInsert(s []int, i int, d int) {
	l := len(s)
	if l <= 0 || i >= l {
		return
	}
	for j := l - 1; j > i; j-- {
		s[j] = s[j-1]
	}
	s[i] = d
}

func TestMoveBack(t *testing.T) {
	s := []int{0, 2, 4, 6, 8}
	fmt.Println(s)
	moveBackOneElmentThenInsert(s, 3, 3)
	fmt.Println(s)
}

func TestMergeSort(t *testing.T) {
	s := []int{37, 0, 79, 76, 62, 91, 35, 62, 44, 19, 89, 38, 99, 57, 100, 87, 76, 56, 4, 25, 82, 88}
	fmt.Println(len(s))
	fmt.Println(s)
	//result := mergeSort(s)
	//fmt.Println("-----------------------------------")
	//fmt.Println(result)
	s2 := mergeSort2(s, 4, SelectionSort)
	fmt.Println(s2)
}

// !!! mergeSort归并排序，归并排序的过程就是将数组分成两半，对每一半进行排序，然
// !!! 将”已排序”的两半合并在一起，重复这个过程，直到整个数组排序完毕。当然，这种
// !!! 不断二分的粒度可以根据具体数据规模加以控制，最小粒度的二分就相当于与每两个相邻
// !!! 的元素作为一个排序小组，排序之后，在对已排序的两个部分逐层向上合并（含排序）。
func mergeSort(s []int) []int {
	//切分排序。按照最小的粒度，将切片s的每两个相邻的元素进行排序。
	rstS := divide(s)
	//归并排序，从最小的粒度开始归并相邻的两组已排序的分组。
	merge(rstS)
	return rstS
}

// !!! 以递归方式对给定数组s进行归并排序,递归是一种易于理解的分治策略实现方式。
//
//	!!! 分治策略要对问题进行分解，当子问题足够大，需要递归求解时，我们称为递归情况（recursivce case）。
//
// !!!  当子问题变得足够小，不需要递归时，我们说递归已经”触底“，进入了基本情况（base case）
// !!! granularity表示二分数组的粒度，数组在不断二分后，一旦规模小于给定的粒度后就不在进行二分，而是
// !!! 进行真正的排序，sortInLocal则表示对最小粒度的数组进行排序时所使用的就地排序算法。
// !!! granularity可以控制二分的嵌套层数，sortInLocal参数可以选择合适的就地排序算法。
func mergeSort2(s []int, granularity int, sortInLocal func([]int)) []int {
	l := len(s)
	rst := make([]int, l)
	if l <= granularity { //!!!当子问题达到了”触底“的基本情况（base case）就不再递归，而执行对最小子问题的处理。
		copy(rst, s)
		sortInLocal(rst)
		//问题分解，将数组二分为左右两个数组
		left := s[:l/2]
		right := s[l/2:]
		//对左数组进行归并排序
		sortedLeft := mergeSort2(left, granularity, sortInLocal)
		//对右侧数组进行归并排序
		sortedRight := mergeSort2(right, granularity, sortInLocal)
		//对已排好序的两侧数组进行归并排序
		sortedS := MergeSortedLists(sortedLeft, sortedRight)
		//将合并后的数组拷贝到结果数组中
		copy(rst, sortedS)
	}
	return rst
}

// !!!divide进行切分排序，将给定切片s相邻的两个元素进行分组，然后排序。
// !!!! 最后输出的切片中，两两分组的组内元素都已排好序（每组只有两个元素，而且是相邻的两个元素）
func divide(s []int) []int {
	slen := len(s)
	rstS := make([]int, slen)
	//相邻两个元素比较，排序后存储如结果数组。
	for i := 0; i < slen; i += 2 {
		if i+1 < slen {
			if s[i] > s[i+1] {
				rstS[i] = s[i+1]
				rstS[i+1] = s[i]
			} else {
				rstS[i] = s[i]
				rstS[i+1] = s[i+1]
			}
		} else {
			rstS[i] = s[i]
		}
	}
	return rstS
}

// !!! merge对已完成切分排序的素组进行就地的归并排序。
// !!! 思路就是以先以2个元素为分组单位，对每相邻的两个分组的进行合并排序，
// !!! 然后再以2*2个元素为分组单位，对每相邻的两个分组进行合并排序，
// !!! 如此，不断以2的n次方个元素来扩大分组的元素规模，直到将给定的切片分组为两个分组进行归并排序为止。
func merge(divS []int) {
	slen := len(divS)
	//!!! 在循环中，将分组的初始规模设置为2，对该规模的分组完成归并后，就将分组的规模翻倍，继续进行归并
	for groupSize := 2; groupSize < slen; groupSize *= 2 {
		var left, right []int
		//!!! 按照分组规模不停地将每相邻的两个分组（左、右两个分组）进行归并排序，直至遍历完整个切片。
		//!!! 主要是考虑按照当前的分组规模对切片不能完整分割的边界情况。
		for i := 0; i < slen; i += groupSize * 2 {
			//当遍历到当前位置，按照分组规模，划分左分组会超出切片边界时，右分组为空。
			if i+groupSize > slen {
				left = divS[i:]
				right = divS[0:0]
			} else { //当做分组不超过切片边界时，右分组可能超界，也可能不超界。
				left = divS[i : i+groupSize]
				if i+groupSize*2 < slen { //右侧分组不超界超界情况
					right = divS[i+groupSize : i+groupSize*2]
				} else {
					right = divS[i+groupSize:] //右侧分组超界情况
				}
			}
			mergedGroup := MergeSortedLists(left, right)
			//!!!将归并后的分组拷贝至原切片，使得原切片可以继续进行下一个分组规模的归并排序。
			copy(divS[i:i+len(mergedGroup)], mergedGroup)
		}
	}
}

// !!! MergeSortedLists对已经排好序的左右两个分组进行合并排序
// !!! 思路就是将左右两个分组中的第一个元素，也就是各自分组最小的元素拿出来进行比较，
// !!! 将最小的元素“移入（从源分组中删除）”结果切片中，这样，直到两个分组的元素都被取空，
// !!! 就完成了两个数组的归并排序。
func MergeSortedLists(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	//!!! 循环要达成的目标状态是两个分组中的元素都被取空，需要
	//!!! 考虑的是可能某一个分组可能先被取空的边界情况。

	for len(left) != 0 || len(right) != 0 {
		var minL, minR int //代表左右两个数组的最小元素（各自分组的第一个元素）

		//!!!左右分组都没取空的情况
		if len(left) > 0 && len(right) > 0 {
			minL = left[0]
			minR = right[0]
			if minL <= minR {
				result = append(result, minL)
				left = left[1:]
			} else {
				result = append(result, minR)
				right = right[1:]
			}
			continue
		}
		//!!! 左侧分组被取空的情况
		if len(left) == 0 && len(right) > 0 {
			minR = right[0]
			result = append(result, minR)
			right = right[1:]
			continue
		}
		//!!! 右侧分组被取空的情况
		if len(right) == 0 && len(left) > 0 {
			minL = left[0]
			result = append(result, minL)
			left = left[1:]
			continue
		}
	}
	return result
}

func TestCountingSort(t *testing.T) {
	//s := []int{37, 0, 79, 76, 62, 91, 35, 62, 44, 19, 89, 38, 99, 57, 87, 76, 56, 45, 25, 82, 88}
	//fmt.Println(s)
	//countingSort(s, 0, 99)
	//fmt.Println(s)
	//生成一个范围在取值范围为0到50，个数为100的随机数组
	s2 := getRandomNnumbers(100, 50)
	fmt.Println(s2)
	fmt.Println("------------------------------------------------")
	countingSort(s2, 0, 50)
	fmt.Println(s2)
}

//!!! 计数排序算法的是一种不基于比较的线性排序方法。
//!!! 工作原理是按照被排序元素的取值范围创建一个新的，
//!!! 所有元素初始值都为0的整数数组，该整数数组的“序号”与被排序元素取值范围中的
//!!! 每个可能元素建立进行一对一映射，这样计数素组的序号的大小既能映射会“源数值”的大小，也表明了其顺序。
//!!! 这个计数数组称为（可能出现元素的）计数数组。这样，遍历被排序的元素，然后找到该元素所映射的计数数组序号，
//!!!  使该序号下的数组元素值加1。这样，在输出的时候，按照计数数组的顺序，
//!!! 输出计数不为0（计数为0表示该元素没有出现过）的技术数组元素序号所映射的排序元素。
//!!! 计数排序适合有对有固定取值范围的元素进行排序，尤其是取值范围不大（内存开销不大），
//!!! 且元素重复出现次数较多的排序场景。

func countingSort(s []int, min, max int) {
	//!!! 根据被排序元素的取值范围制作计数素组
	countings := make([]int, (max - min + 1))
	//!!! 遍历被排序元素，按照被排序元素值与计数数组序号的映射关系增加被排序元素的出现次数
	for _, e := range s {
		countings[e-min] += 1 //!!! 被排序的元素e，映射为计数数组元素的序号为 e-min
	}
	var i = 0
	//根据出现的元素的计数，回写排序的结果
	for j, c := range countings { //j+min映射为s[i]的值,c代表s[i]的值出现的次数
		for ; c > 0; c-- {
			s[i] = j + min //j+min映射为S[i]的值
			i += 1
		}
	}
}

// !!! 基数排序
func RadixSort(s []int) {

}
