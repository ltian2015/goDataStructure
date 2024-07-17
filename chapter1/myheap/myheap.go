package myheap

import "sort"

// The Interface type describes the requirements
// for a type using the routines in this package.
// !!! MHeap类型描述了这个包所提供的堆处理例程（routines）对可以作为堆的输入源的要求。
// Any type that implements it may be used as a
// min-heap with the following invariants (established after
// [Init] has been called or if the data is empty or sorted):
// !!! 任何实现了MHeap接口的类型,可以被当作一个带有下列不变性的“最小堆（min-heap）"来使用。
// !!! （最小堆的建立可以有三种情况：[Init]函数被调用，或数据是空的，或数据事先已排序）
//
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//
// Note that [Push] and [Pop] in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use [heap.Push] and [heap.Pop].
// !!! 注意：MHeap接口中的[Push] 和 [Pop]接口方法用于本包heap实现的调用。
type MHeap interface {
	sort.Interface
	Push(x any) // add x as element Len() //!!!将x作为最后的元素追加进来。
	Pop() any   // remove and return element Len() - 1.//!!!将最后的元素移除，并返回。
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
// !!!Init建立本包中除了Init方法之外其他堆处理方法（Push、Pop、Fix、Remove）所能处理的堆（head）不变量
func Init(h MHeap) {
	// heapify //!!!堆化处理，按照less函数所设定的规则，使最小或最大的元素成为根节点。
	n := h.Len()
	//!!! n/2 - 1是完整二叉树按照先根访问顺序将树节点依次放在列表中后，
	//!!! 该完整二叉树最后一棵最小子树的根节点的序号。
	//!!! 从该节点向前遍历，就是遍历各级子树的根节点，直到整个二叉树的根节点。
	/*!!! 有关二叉树的知识详见 file://./数据结构学习-二叉树.pdf */
	for i := n/2 - 1; i >= 0; i-- {
		//!!! down操作就是将给定根节点为i的二叉树的根节点的值作为整个二叉
		//!!! 树的最小值（也可以是最大值，取决于Less函数的比较语义）。
		//!!! 由于down操作是从最后一棵最小子树开始，不断先前\向上遍历各级子树，
		//!!! 所以，当处理完整棵二叉树后，整个树的根节点就是所有节点的最小值(也可能是最大值)
		down(h, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
// !!! push操作将追加元素到堆中。
func Push(h MHeap, x any) {
	h.Push(x)        //将元素追加为堆的数据源的最后一个元素。
	up(h, h.Len()-1) //将最后追加的元素作为树的叶子节点，按规则上浮。
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to [Remove](h, 0).
// !!! Pop操作从堆中弹出并返回最小的元素（取决于less函数的具体实现逻辑）。
// !!! 算法的复杂度是O(log n)，这里n=h.len()
// !!! Pop 等效于 Remove(h, 0)
func Pop(h MHeap) any {
	n := h.Len() - 1
	h.Swap(0, n)   //!!! 把最后一个节点与当前的根节点交换。
	down(h, 0, n)  //!!! 对当前的根节点进行下沉处理。确保最小的节点在根节点。
	return h.Pop() //!!!将最后的元素移除，并返回最后的元素。
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
// !!! Remove操作从堆中移除并返回位于序号i处的元素。
// !!! 算法的复杂度是O(log n)，这里n=h.len()
func Remove(h MHeap, i int) any {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)        //!!! 将要被移除的元素与最后一个元素交换。
		if !down(h, i, n) { //将新移到i位置处的元素（原来的最后一个元素）进行下沉处理，如果不能下沉就做上浮处理。
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling [Remove](h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
// !!! Fix操作能够子在堆的第i个元素的值发生变化时重新建立堆的顺序。
// 改变在位置i的元素的值然后调用Fix方法在效果上“等价于”先调用Remove(h, i)，
// 然后f再Push新值，但是Fix方法会比后者的代价更小，其复杂度为O(log n)，这里n=h.Len()
func Fix(h MHeap, i int) {
	if !down(h, i, h.Len()) { //如果i不能下沉，那么就对i进行上浮处理。
		up(h, i)
	}
}

//!!! up操作使得给定的子节点按照规则（小于或大于根节点，取决于less的具体实现），上浮为根节点。
// h表示”堆“的数据源
//j 表示参数h所给定的堆中的元素序号（先根遍历下，二叉树节点的序号）,up函数将其作为“子节点”进行处理。

func up(h MHeap, j int) {
	for { //反复循环处理
		//!!! i表示给子定节点j的根节点。
		i := (j - 1) / 2 // parent//!!!完整二叉树按先根顺序存储于列表中时，对于给定节点j，其根节点的位置为 （j-1）/2
		//!!!当给定节点j“不小于（具体含义取决于less的实现）” 其根节点，就退出循环。
		if i == j || !h.Less(j, i) {
			break
		}
		//!!!当给定节点小于其根节点时，就进行节点的位置交换，把位置为j处的子节点与位置为i出的节点交换，即将原位置j的节点上浮为根节点（放在位置i处）.
		h.Swap(i, j)
		j = i //!!! 再次将新放入值的根节点（i）作为子节点，继续进行上浮操作的循环。
	}
}

//!!! down操作使得给定节点i0作为根节点，按照规则（不是最大或最小就下沉，取决于less的具体实现）下沉为子节点。
//!!! down操作可以保证如果给定的根节点不是下级节点最小值时，就下沉一级，如果下沉后，仍不是最小节点，
//!!! 就继续下沉，直到下一级节点没有比它更小为止，但不能保证下二级节点（孙子节点）有比它还
//!!! 小的节点,为此，这就要求初始化堆要从最后一棵最小子树开始，向前、向上做down处理，这样才能保证
//!!! 整棵树的

// h表示”堆“的数据源
// i0表示 参数h所给定的堆中的元素序号（先根遍历下，二叉树节点的序号），down函数将其作为根节点进行处理。
//注意，i0所表示的节点都是有子节点的根节点
// n表示参数h所给定的堆中的数据长度

func down(h MHeap, i0, n int) bool {

	i := i0 //注意，i代表了当前循环所处理“根节点”的位置。
	for {
		j1 := 2*i + 1 //!!! 完整二叉树按先根顺序存储于列表中时，位置为i的节点的左子树节点位于2*i+1处。
		//!!! 判断左子树是否存在，如果左子树序号范围超界，则表示没有左子树，如果没有左子树，那么也没有右子树。
		//!!! 当给定的节点没有子树时，就结束循环，表明已经走到了树的某个叶子节点。
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}

		j := j1 // left child //!!! 位置j表示的是左子树节点
		//!!! j2表示右子树节点的位置，下面的代码找到左、右子树中最小值的节点序号。
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		//!!! 如果左右两个节点中，最小的子节点不小于根节点，那么根节点就无法下沉了，就结束循环。
		//!!! 但这就要求必须初始化时必须从最小子树向前\向上进行下沉（down）处理，
		//!!! 否则无法保证叶子结点一定比根节点小（或大）。
		//!!! 如果初始化从最小子树开始做下沉处理，那么就保证了所有根节点都是比子节点小，
		//!!! 在此前提下，对某个发生了变动的根节点做下沉处理时，只要最小的子节点都不小于根节点时，
		//!!! 就无需继续进行下沉处理了。
		if !h.Less(j, i) {
			break
		}
		//!!! 当子树的最小节点小于根节点时，将其作为根节点，此时根节点最小。
		h.Swap(i, j)
		//!!!  再次将从根节点下沉为子节点的作为根节点，继续循环，做下沉处理。
		i = j //注意，i表示新一轮要进行下沉处理的根节点的序号。J是本轮循环中子节点的序号。
	}
	// !!!返回值表明给定的节点是否下沉为子节点。ture表明下沉，false表明未下沉。
	return i > i0
}
