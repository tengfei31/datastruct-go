package order

import "log"

//Order 获取排序元素的接口
type Order interface {
	SetKey(int)
	GetKey() int
}

type T struct {
	Order
	K int
}

//SetKey 设置要比较的元素
func (t *T) SetKey(k int) {
	t.K = k
}

//GetKey 获取要比较的元素
func (t *T) GetKey() int {
	return t.K
}

type List struct {
	Size, MaxList int
	Elements      []T
}

//Node 链表节点
type Node struct {
	Element T
	Link    *Node
}

//LinkList 链表
type LinkList struct {
	First *Node
	Size  int
}

//InsertSortOrder 插入排序：顺序表上的直接插入排序
func InsertSortOrder(lst *List) {
	var (
		i, j int
		x    T
	)
	//执行n-1趟
	for i = 1; i < lst.Size; i++ {
		//待插入元素存入临时变量
		x = lst.Elements[i]
		//从后往前查找插入位置
		for j = i - 1; j >= 0 && x.GetKey() < lst.Elements[j].GetKey(); j-- {
			//元素后移，j指针前移
			lst.Elements[j+1].SetKey(lst.Elements[j].GetKey())
		}
		//待插入元素存入找到的插入位置
		lst.Elements[j+1].SetKey(x.GetKey())
	}
}

//InsertSortLinkList 插入排序：链表上的直接插入排序
func InsertSortLinkList(lst *LinkList) {
	//空链表
	if lst.First == nil {
		log.Fatalln("链表为空")
	}
	var unsorted, sorted, p, q *Node
	sorted = lst.First
	//至少一个节点
	for sorted.Link != nil {
		//unsorted 指示带插记录
		unsorted = sorted.Link
		//若待插记录小于第一个记录
		if unsorted.Element.GetKey() < lst.First.Element.GetKey() {
			//将待插记录从链表取下
			sorted.Link = unsorted.Link
			//将待插记录插在链表的最前面
			unsorted.Link = lst.First
			lst.First = unsorted
		} else { //若待插记录大于等于第一个记录
			q = lst.First
			p = q.Link
			//搜索待插记录的适当插入位置
			for unsorted.Element.GetKey() > p.Element.GetKey() {
				q = p
				p = p.Link
			}
			if unsorted == p {
				//将待插记录在有序子表末尾
				sorted = unsorted
			} else {
				//将待插记录插在节点*q之后
				sorted.Link = unsorted.Link
				unsorted.Link = p
				q.Link = unsorted
			}
		}
	}
}
