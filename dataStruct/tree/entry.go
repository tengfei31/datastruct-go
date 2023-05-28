package tree

import (
	"datastruct-go/dataStruct/listarr"
	"log"
)

//KeyType 关键字类型
type KeyType int

//DataType 数据域类型
type DataType int

//Entry 数据类型
type Entry struct {
	Key  KeyType
	Data DataType
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
func btSearch(node *BTNode[int], k KeyType) *BnElement[int] {
	if node == nil {
		return nil
	}
	if KeyType(node.Element.GetWeight()) == k {
		return &node.Element
	} else if KeyType(node.Element.GetWeight()) < k {
		return btSearch(node.RChild, k)
	} else {
		return btSearch(node.LChild, k)
	}
}

//BtSearch 二叉搜索树(递归算法)
func BtSearch(tree Btree, k KeyType, x *BnElement[int]) bool {
	var res = btSearch(tree.Root, k)
	if res == nil {
		return false
	}
	*x = *res
	return true
}

//BtSearch2 二叉搜索树（迭代算法）
func BtSearch2(tree Btree, k KeyType, x *BnElement[int]) bool {
	var node = tree.Root
	if node == nil {
		return false
	}
	for node != nil {
		if KeyType(node.Element.GetWeight()) == k {
			*x = node.Element
			return true
		} else if KeyType(node.Element.GetWeight()) < k {
			node = node.LChild
		} else {
			node = node.RChild
		}
	}
	return false
}

//BtInsert 二叉搜索树的插入
func BtInsert(tree *Btree, x BnElement[int]) bool {
	var q, p *BTNode[int]
	p = tree.Root
	var k = x
	for p != nil {
		q = p
		if p.Element == k {
			log.Printf("Duplicate")
			return false
		} else if p.Element.GetWeight() > k.GetWeight() {
			p = p.LChild
		} else {
			p = p.RChild
		}
	}
	var r = NewNode2(x)
	if tree.Root != nil {
		if k.GetWeight() < q.Element.GetWeight() {
			q.LChild = r
		} else {
			q.RChild = r
		}
	} else {
		tree.Root = r
	}
	return true
}

//BtRemove 二叉搜索树的删除
func BtRemove(tree *Btree, k KeyType, x *BnElement[int]) bool {
	var c, r, s, p, q *BTNode[int]
	p = tree.Root
	for p != nil && KeyType(p.Element.GetWeight()) != k {
		q = p
		if k < KeyType(p.Element.GetWeight()) {
			p = p.LChild
		} else {
			p = p.RChild
		}
	}
	if p == nil {
		log.Printf("No element with k")
		return false
	}
	*x = p.Element
	if p.LChild != nil && p.RChild != nil {
		s = p.RChild
		r = p
		for s.LChild != nil {
			r = s
			s = s.LChild
		}
		p.Element = s.Element
		p = s
		q = r
	}
	if p.LChild != nil {
		c = p.LChild
	} else {
		c = p.RChild
	}
	if p == tree.Root {
		tree.Root = c
	} else if p == q.LChild {
		q.LChild = c
	} else {
		q.RChild = c
	}
	p = nil
	return true
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
