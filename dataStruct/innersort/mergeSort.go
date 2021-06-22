package innersort

//合并排序

//Merge 合并函数
func (lst *List) Merge(temp []T, i1, j1, i2, j2 int, k *int) {
	var (
		i = i1
		j = i2
	)
	//若两个子序列都不空，则循环
	for i <= j1 && j <= j2 {
		if lst.Elements[i].GetKey() <= lst.Elements[j].GetKey() {
			//将较小元素存入temp[*k]
			temp[*k] = lst.Elements[i]
			*k++
			i++
		} else {
			temp[*k] = lst.Elements[j]
			*k++
			j++
		}
	}
	//将子序列1中剩余的元素存入temp
	for i <= j1 {
		temp[*k] = lst.Elements[i]
		*k++
		i++
	}
	//将自学列2中剩余的元素存入temp
	for j <= j2 {
		temp[*k] = lst.Elements[j]
		*k++
		j++
	}
}

//MergeSort 合并排序
func (lst *List) MergeSort() {
	var temp []T = make([]T, lst.Size)
	//i1, j1和i2, j2分别是两个子序列的上、下界
	var i1, j1, i2, j2, i, k int
	var size = 1
	for size < lst.Size {
		i1 = 0
		k = 0
		//若i1+size<n，则说明存在两个子序列，需要两两合并
		for i1+size < lst.Size {
			//确定子序列2的下界和子序列1的上界
			i2 = i1 + size
			j1 = i2 - 1
			//设置子序列2的上界
			if i2+size-1 > lst.Size-1 {
				j2 = lst.Size - 1
			} else {
				j2 = i2 + size - 1
			}
			//合并相邻两个子序列
			lst.Merge(temp, i1, j1, i2, j2, &k)
			//确定下一次合并第一个子序列的下界
			i1 = j2 + 1
		}
		for i = 0; i < k; i++ {
			lst.Elements[i] = temp[i]
		}
		//子序列长度扩大一倍
		size *= 2
	}
}

//Divide 链表分割函数
func (node *Node) Divide(p *Node) *Node {
	var pos, mid, q *Node
	if p == nil {
		return nil
	}
	mid = p
	pos = mid.Link
	for pos != nil {
		pos = pos.Link
		if pos != nil {
			mid = mid.Link
			pos = pos.Link
		}
	}
	q = mid.Link
	mid.Link = nil
	return q
}

//Merge 单链表合并函数
func (node *Node) Merge(p *Node, q *Node) *Node {
	//head为哑节点，rear为指针变量
	var (
		rear = new(Node)
		head = Node{}
	)
	//rear指向head节点
	rear = &head
	//合并两个有序链表
	for p != nil && q != nil {
		if p.Element.GetKey() <= q.Element.GetKey() {
			rear.Link = p
			rear = p
			p = p.Link
		} else {
			rear.Link = q
			rear = q
			q = q.Link
		}
	}
	//将一个链表的剩余部分链至结果链表的尾部
	if p == nil {
		rear.Link = q
	} else {
		rear.Link = p
	}
	//返回结果链表的起始节点地址
	return head.Link
}

//RMSort 两路合并排序函数
func RMSort(sublst **Node) {
	if *sublst != nil && (*sublst).Link != nil {
		second := (*sublst).Divide(*sublst)
		RMSort(sublst)
		RMSort(&second)
		*sublst = (*sublst).Merge(*sublst, second)
	}
}

//RMergeSort 两路合并排序函数
func RMergeSort(lst *LinkList) {
	RMSort(&(lst.First))
}
