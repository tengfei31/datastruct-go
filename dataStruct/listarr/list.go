package listarr

import "log"

const MaxSize = 10

type T struct {
	Key  int
	Data int
}

type List struct {
	Size    int
	MaxList int
	Element [MaxSize]T
}

//CreateList 初始化list
func (lst *List) CreateList(maxList int) {
	//lst = new(List)
	lst.Size = 0
	lst.MaxList = maxList
}

//IsFull 是否满了
func (lst *List) IsFull() bool {
	if lst.Size >= lst.MaxList {
		return true
	}
	return false
}

//Push 添加一个元素
func (lst *List) Push(element T) {
	if lst.IsFull() == true {
		log.Fatalf("list is full")
	}
	lst.Element[lst.Size] = element
	lst.Size += 1
}
