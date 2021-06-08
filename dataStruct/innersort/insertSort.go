package innersort

import "log"

//插入排序

//Order 获取排序元素的接口
type Order interface {
	SetKey(int)
	GetKey() int
}

type T struct {
	//Order
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

//NewList 创建List
func NewList(maxList int) *List {
	return &List{
		Size:     0,
		MaxList:  maxList,
		Elements: make([]T, 0, maxList),
	}
}

//AddElement 添加元素
func (lst *List) AddElement(element T) {
	var size = len(lst.Elements)
	if size >= lst.MaxList {
		panic("长度超出限制")
	}
	lst.Elements = append(lst.Elements, element)
	lst.Size = size + 1
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

//NewLinkList 创建linklist
func NewLinkList() *LinkList {
	return &LinkList{
		Size:  0,
		First: nil,
	}
}

//AddElement 往最前面塞element
func (lst *LinkList) AddElement(element T) {
	node := new(Node)
	node.Element = element
	node.Link = nil
	tmp := lst.First
	lst.First = node
	node.Link = tmp
	lst.Size += 1
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
		//unsorted 指示待插记录
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

//insertSortInsSort 插入排序：修改后的直接插入排序
func insertSortInsSort(lst *List, h int) {
	var (
		i, j int
		x    T
	)
	for i = h; i < lst.Size; i += h {
		x = lst.Elements[i]
		//位置j的同组前一个位置为j-h
		for j = i - h; j >= 0 && x.GetKey() < lst.Elements[j].GetKey(); j -= h {
			lst.Elements[j+h] = lst.Elements[j]
		}
		lst.Elements[j+h] = x
	}
}

//InsertSortShellSort 插入排序：希尔排序
func InsertSortShellSort(lst *List) {
	var (
		i    int
		incr = lst.Size
	)
	for {
		//计算增量
		incr = incr/3 + 1
		for i = 0; i < incr; i++ {
			insertSortInsSort(lst, incr)
		}
		if incr <= 1 {
			break
		}
	}
}
