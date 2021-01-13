package tree

import "datastruct-go/dataStruct/listarr"

//KeyType 关键字类型
type KeyType int

//DataType 数据域类型
type DataType int

//Entry 数据类型
type Entry struct {
	Key  KeyType
	Data DataType
}

//Set 集合
type Set interface {
	//CreateList 创建空集合
	CreateList(maxSize int)
	//IsEmpty 集合是否为空
	IsEmpty() bool
	//IsFull 集合是否满了
	IsFull() bool
	//Search 在集合中搜索关键字值为k的元素，并将该元素放入x中，并返回true
	Search(k KeyType, x *T) bool
	//Insert 在集合中搜索关键字值为k的元素，如果不存在就插入到集合中，并返回true,否则返回false
	Insert(x T) bool
	//Remove 在集合中搜索关键字值为k的元素，如果存在，就将该元素赋值给*x, 并从集合中删除该元素，返回true，否则返回false
	Remove(k KeyType, x *T) bool
}

//SeqSearch 顺序搜索无序表
func SeqSearch(lst listarr.List, k KeyType, x *listarr.T) bool {
	if lst.Size > 0 {
		for i := 0; i < lst.Size; i++ {
			if KeyType(lst.Element[i].Key) == k {
				*x = lst.Element[i]
				return true
			}
		}
	}
	return false
}

const MaxNum = 100

//SeqSearch2 顺序搜索有序表
func SeqSearch2(lst listarr.List, k KeyType, x *listarr.T) bool {
	var i int
	lst.Element[lst.Size].Key = MaxNum
	for i = 0; KeyType(lst.Element[i].Key) < k; i++ {
	}
	if KeyType(lst.Element[i].Key) == k {
		*x = lst.Element[i]
		return true
	}
	return false
}

//bSch 二分搜索(递归)
func bSch(lst listarr.List, k KeyType, low int, high int) int {
	var mid int = -1
	if low <= high {
		mid = (low + high) / 2
		if k < KeyType(lst.Element[mid].Key) {
			return bSch(lst, k, low, mid-1)
		} else if k > KeyType(lst.Element[mid].Key) {
			return bSch(lst, k, mid+1, high)
		} else {
			return mid
		}
	}
	return mid
}

//BSearch 二分查找(递归)
func BSearch(lst listarr.List, k KeyType, x *listarr.T) bool {
	var i int
	i = bSch(lst, k, 0, lst.Size)
	if i < 0 {
		return false
	}
	*x = lst.Element[i]
	return true
}

//BSearch2 二分搜索(迭代算法)
func BSearch2(lst listarr.List, k KeyType, x *listarr.T) bool {
	var mid int
	var low int = 0
	var high int = lst.Size
	for low <= high {
		mid = (low + high) / 2
		if k < KeyType(lst.Element[mid].Key) {
			high = mid - 1
		} else if k > KeyType(lst.Element[mid].Key) {
			low = mid + 1
		} else {
			*x = lst.Element[mid]
			return true
		}
	}
	return false
}

//bSearchTr 二叉搜索树(递归算法)
func bSearchTr(node *BTNode, k KeyType) *T {
	if node == nil {
		return nil
	}
	if KeyType(node.Element) == k {
		return &node.Element
	} else if KeyType(node.Element) < k {
		return bSearchTr(node.LChild, k)
	} else {
		return bSearchTr(node.RChild, k)
	}
}

//BSearchTree 二叉搜索树(递归算法)
func BSearchTree(tree Btree, k KeyType, x *T) bool {
	var res = bSearchTr(tree.Root, k)
	if res == nil {
		return false
	}
	*x = *res
	return true
}

//BSearchTree2 二叉搜索树（迭代算法）
func BSearchTree2(tree Btree, k KeyType, x *T) bool {
	var node = tree.Root
	if node == nil {
		return false
	}
	for node != nil {
		if KeyType(node.Element) == k {
			*x = node.Element
			return true
		} else if KeyType(node.Element) < k {
			node = node.LChild
		} else {
			node = node.RChild
		}
	}
	return false
}

//BFSearch 斐波那契搜索
func BFSearch(lst listarr.List, k KeyType, x *listarr.T) bool {
	var (
		t int
		i int = 1
		p int = 1
		q int = 2
		n int = lst.Size
	)
	for q <= n {
		p = q
		q = i + q
		i = p
	}
	p = q - i
	q = i - p
	for {
		if KeyType(lst.Element[i-1].Key) == k {
			*x = lst.Element[i-1]
			return true
		} else if KeyType(lst.Element[i-1].Key) > k {
			if q == 0 {
				return false
			} else {
				i = i - q
				t = p
				p = q
				q = t - q
			}
		} else if p == 1 {
			return false
		} else {
			i = i - q
			p = p - q
			q = q - p
		}
	}
}
