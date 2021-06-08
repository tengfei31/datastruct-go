package innersort

//交换排序

//Swap 交换x,y
func Swap(x, y, t *T) {
	t = x
	x = y
	y = t
}

//BubbleSort 冒泡排序
func BubbleSort(lst *List) {
	var (
		j int
		i = lst.Size - 1
		//temp T
		sorted = false
	)
	for i > 0 && !sorted {
		sorted = true
		for j = 0; j < i; j++ {
			if lst.Elements[j+1].GetKey() < lst.Elements[j].GetKey() {
				//交换位置
				//Swap(&lst.Elements[j+1], &lst.Elements[j], &temp)
				lst.Elements[j+1], lst.Elements[j] = lst.Elements[j], lst.Elements[j+1]
				sorted = false
			}
		}
		i--
	}
}

//BubbleSort1 改进冒泡排序
func BubbleSort1(lst *List) {
	var (
		j, last int
		i       = lst.Size - 1
	)
	//最多进行n-1趟
	for i > 0 {
		//进行循环就将last置成0
		last = 0
		//从前从后进行相邻元素的两两比较
		for j = 0; j < i; j++ {
			if lst.Elements[j+1].GetKey() < lst.Elements[j].GetKey() {
				//交换位置
				//Swap(&lst.Elements[j+1], &lst.Elements[j], &temp)
				lst.Elements[j+1], lst.Elements[j] = lst.Elements[j], lst.Elements[j+1]
				last = j
			}
		}
		i = last
	}
}

//partition 快速排序：分划函数
func partition(lst *List, left, right int) int {
	var (
		//确定分划序列的指针i, j
		i     = left
		j     = right + 1
		pivot = lst.Elements[left]
	)
	for {
		//i从左向右找第一个不小于pivot的元素
		for {
			i++
			if i >= lst.Size || lst.Elements[i].GetKey() >= pivot.GetKey() {
				break
			}
		}
		//j从右向左找第一个不大于pivot的元素
		for {
			j--
			if j < 0 || lst.Elements[j].GetKey() <= pivot.GetKey() {
				break
			}
		}
		if i < j {
			lst.Elements[i], lst.Elements[j] = lst.Elements[j], lst.Elements[i]
		} else {
			break
		}
	}
	//交换位于left和j的 元素
	lst.Elements[left], lst.Elements[j] = lst.Elements[j], lst.Elements[left]
	return j
}

//QSort 快速排序的递归函数
func QSort(lst *List, left, right int) {
	var k int
	if left < right {
		//对left和right之间的序列进行分划
		k = partition(lst, left, right)
		//对左子序列进行快速排序
		QSort(lst, left, k-1)
		//对右子序列进项快速排序
		QSort(lst, k+1, right)
	}
}

//QuickSort 快速排序：对顺序表实现的快速排序
func QuickSort(lst *List) {
	//以给定的初始序列lst调用快速排序函数
	QSort(lst, 0, lst.Size-1)
}
