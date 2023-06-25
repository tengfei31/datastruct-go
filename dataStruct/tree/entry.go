package tree

import (
	"datastruct-go/dataStruct/listarr"
	"log"
)

// Entry 数据类型
type Entry[T int] struct {
	Key  T
	Data T
}

// SeqSearch 顺序搜索无序表
func SeqSearch[T int](lst listarr.List[T], k T, x *listarr.ElementT[T]) bool {
	if lst.Size > 0 {
		var i int
		for i = 0; i < lst.Size; i++ {
			if lst.Element[i].Key == k {
				*x = lst.Element[i]
				return true
			}
		}
	}
	return false
}

const MaxNum = 100

// SeqSearch2 顺序搜索有序表
func SeqSearch2[T int](lst listarr.List[T], k T, x *listarr.ElementT[T]) bool {
	var i T
	lst.Element[lst.Size].Key = MaxNum
	for i = 0; lst.Element[i].Key < k; i++ {
	}
	if lst.Element[i].Key == k {
		*x = lst.Element[i]
		return true
	}
	return false
}

// bSch 二分搜索(递归)
func bSch[T int](lst listarr.List[T], k T, low int, high int) int {
	var mid int = -1
	if low <= high {
		mid = (low + high) / 2
		if k < lst.Element[mid].Key {
			return bSch[T](lst, k, low, mid-1)
		} else if k > lst.Element[mid].Key {
			return bSch[T](lst, k, mid+1, high)
		} else {
			return mid
		}
	}
	return mid
}

// BSearch 二分查找(递归)
func BSearch[T int](lst listarr.List[T], k T, x *listarr.ElementT[T]) bool {
	var i int
	i = bSch[T](lst, k, 0, lst.Size)
	if i < 0 {
		return false
	}
	*x = lst.Element[i]
	return true
}

// BSearch2 二分搜索(迭代算法)
func BSearch2[T int](lst listarr.List[T], k T, x *listarr.ElementT[T]) bool {
	var mid int
	var low int = 0
	var high int = lst.Size
	for low <= high {
		mid = (low + high) / 2
		if k < lst.Element[mid].Key {
			high = mid - 1
		} else if k > lst.Element[mid].Key {
			low = mid + 1
		} else {
			*x = lst.Element[mid]
			return true
		}
	}
	return false
}

// bSearchTr 二叉搜索树(递归算法)
func btSearch[T int](node *BTNode[T], k T) *BnElement[T] {
	if node == nil {
		return nil
	}
	if node.Element.GetWeight() == k {
		return &node.Element
	} else if node.Element.GetWeight() < k {
		return btSearch[T](node.RChild, k)
	} else {
		return btSearch[T](node.LChild, k)
	}
}

// BtSearch 二叉搜索树(递归算法)
func BtSearch[T int](tree Btree[T], k T, x *BnElement[T]) bool {
	var res = btSearch[T](tree.Root, k)
	if res == nil {
		return false
	}
	*x = *res
	return true
}

// BtSearch2 二叉搜索树（迭代算法）
func BtSearch2[T int](tree Btree[T], k T, x *BnElement[T]) bool {
	var node = tree.Root
	if node == nil {
		return false
	}
	for node != nil {
		if node.Element.GetWeight() == k {
			*x = node.Element
			return true
		} else if node.Element.GetWeight() < k {
			node = node.LChild
		} else {
			node = node.RChild
		}
	}
	return false
}

// BtInsert 二叉搜索树的插入
func BtInsert[T int](tree *Btree[T], x BnElement[T]) bool {
	var q, p *BTNode[T]
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

// BtRemove 二叉搜索树的删除
func BtRemove[T int](tree *Btree[T], k T, x *BnElement[T]) bool {
	var c, r, s, p, q *BTNode[T]
	p = tree.Root
	for p != nil && p.Element.GetWeight() != k {
		q = p
		if k < p.Element.GetWeight() {
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

// BFSearch 斐波那契搜索
func BFSearch[T int](lst listarr.List[T], k T, x *listarr.ElementT[T]) bool {
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
		if lst.Element[i-1].Key == k {
			*x = lst.Element[i-1]
			return true
		} else if lst.Element[i-1].Key > k {
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
