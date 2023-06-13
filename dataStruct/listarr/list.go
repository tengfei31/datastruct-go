package listarr

import "log"

const MaxSize = 10

type ElementT[T int] struct {
	Key  T
	Data T
}

type List[T int] struct {
	Size    int
	MaxList int
	Element [MaxSize]ElementT[T]
}

// CreateList 初始化list
func (lst *List[T]) CreateList(maxList int) {
	//lst = new(List)
	lst.Size = 0
	lst.MaxList = maxList
}

// IsFull 是否满了
func (lst *List[T]) IsFull() bool {
	if lst.Size >= lst.MaxList {
		return true
	}
	return false
}

// Push 添加一个元素
func (lst *List[T]) Push(element ElementT[T]) {
	if lst.IsFull() == true {
		log.Fatalf("list is full")
	}
	lst.Element[lst.Size] = element
	lst.Size += 1
}
