package tree

import "log"

const MaxSize = 10

//MinHeap 最小堆结构
type MinHeap struct {
	Size     int
	MaxHeap  int
	Elements []T
}

//AdjustDown 最小堆向下调整
func AdjustDown(heap []T, r int, n int) {
	var child int = r * 2
	var temp T = heap[r]
	for child <= n {
		if child < n && heap[child] > heap[child+1] {
			child++
		}
		if temp <= heap[child] {
			break
		}
		heap[child/2] = heap[child]
		child *= 2
	}
	heap[child/2] = temp
}

//AdjustUp 最小堆向上调整
func AdjustUp(heap []T, n int) {
	var i int = n
	var temp T = (heap)[i]
	for i != 1 && temp < (heap)[i/2] {
		(heap)[i] = (heap)[i/2]
		//hp.Elements[i/2] = temp
		i /= 2
	}
	(heap)[i] = temp
}

//CreateHeap 创建最小堆
func CreateHeap(hp *MinHeap) {
	var i int
	var n int = hp.Size
	for i = n / 2; i > 0; i-- {
		AdjustDown(hp.Elements, i, n)
	}
}

//PQueue 优先权队列
type PQueue MinHeap

//CreatePQ 创建一个空的优先权队列
func (pq *PQueue) CreatePQ(maxSize int) {
	pq.MaxHeap = maxSize
	pq.Size = 0
	pq.Elements = make([]T, 0, pq.MaxHeap)
	pq.Elements[0] = 0
}

//IsEmpty 优先权队列是否为空
func (pq *PQueue) IsEmpty() bool {
	return pq.Size <= 0
}

//IsFull 优先权队列是否满了
func (pq *PQueue) IsFull() bool {
	return pq.Size >= pq.MaxHeap
}

//Append 将元素加入队列
func (pq *PQueue) Append(x T) {
	if pq.IsFull() == true {
		log.Fatalf("Overflow")
	}
	pq.Size++
	pq.Elements[pq.Size] = x
	//调整堆
	AdjustUp(pq.Elements, pq.Size)
}

//Server 返回最高优先权的元素值，并从队列中删除该元素
func (pq *PQueue) Serve() T {
	if pq.IsEmpty() == true {
		log.Fatalf("Underflow")
	}
	var x T
	x = pq.Elements[1]
	//调整堆
	pq.Elements[1] = pq.Elements[pq.Size]
	pq.Size--
	AdjustDown(pq.Elements, 1, pq.Size)
	return x
}
