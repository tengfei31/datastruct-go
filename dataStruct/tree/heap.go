package tree

import (
	"datastruct-go/dataStruct/base_interface"
	"log"
)

const MaxSize = 10

// MinHeap 最小堆结构
type MinHeap struct {
	Size     int
	MaxHeap  int
	Elements []BnElement[int]
}

// NewMinHeap 创建最小堆对象
func NewMinHeap(maxHeap int) *MinHeap {
	return &MinHeap{
		Size:     0,
		MaxHeap:  maxHeap,
		Elements: make([]BnElement[int], maxHeap),
	}
}

// HeapSort 最小堆排序
func (hp *MinHeap) HeapSort() {
	var i int
	//建堆
	for i = hp.Size / 2; i > 0; i-- {
		AdjustDown(hp.Elements, i, hp.Size)
	}
	//选择排序
	for i = hp.Size; i > 1; i-- {
		hp.Elements[1], hp.Elements[i] = hp.Elements[i], hp.Elements[1]
		AdjustDown(hp.Elements, 1, i-1)
	}
}

// AdjustDown 最小堆向下调整
func AdjustDown(heap []BnElement[int], r int, n int) {
	var child = r * 2
	var temp = heap[r]
	for child <= n {
		// T = int
		//if child < n && heap[child] > heap[child+1] {
		//	child++
		//}
		//if temp <= heap[child] {
		//	break
		//}
		// T = graph.EdgeNode
		if child < n && heap[child].GetWeight() > heap[child+1].GetWeight() {
			child++
		}
		if temp.GetWeight() < heap[child].GetWeight() {
			break
		}
		heap[child/2] = heap[child]
		child *= 2
	}
	heap[child/2] = temp
}

// AdjustUp 最小堆向上调整
func AdjustUp(heap []BnElement[int], n int) {
	var i = n
	var temp = heap[i]
	// T = int
	//for i != 1 && temp < heap[i/2] {
	//	heap[i] = heap[i/2]
	//	i /= 2
	//}
	// T = graph.EdgeNode
	for i != 1 && temp.GetWeight() < heap[i/2].GetWeight() {
		heap[i] = heap[i/2]
		i /= 2
	}
	heap[i] = temp
}

// CreateHeap 创建最小堆
func (hp *MinHeap) CreateHeap() {
	var i int
	var n = hp.Size
	for i = n / 2; i > 0; i-- {
		AdjustDown(hp.Elements, i, n)
	}
}

// Point 最高优先权队列值
type Point struct {
	base_interface.Weight
	//其实可以不用这个参数的，只要实现weight接口就可以了，现在我只是为了方便
	W   int
	Val int
}

// GetWeight 获取权重
func (p Point) GetWeight() int {
	return p.W
}

// SetWeight 设置权重
func (p *Point) SetWeight(w int) {
	p.W = w
}

// MakeMinHeap 制作最小堆
func MakeMinHeap() *MinHeap {
	var heap = NewMinHeap(20)
	var tmpSetElement = []int{503, 87, 512, 61, 908, 170, 275, 154, 509, 426, 523, 289, 456, 329, 77, 135, 500, 266}
	for val := range tmpSetElement {
		heap.Elements = append(heap.Elements, BnElement[int]{W: val})
	}
	heap.Size = len(heap.Elements) - 1
	heap.CreateHeap()
	return heap
}

// PQueue 优先权队列
type PQueue MinHeap

// CreatePQ 创建一个空的优先权队列
func (pq *PQueue) CreatePQ(maxSize int) {
	pq.MaxHeap = maxSize
	pq.Size = 0
	pq.Elements = make([]BnElement[int], pq.MaxHeap)
	// T = int
	//pq.Elements[0] = 0
	// T = graph.EdgeNode
	pq.Elements[0] = BnElement[int]{}
}

// IsEmpty 优先权队列是否为空
func (pq *PQueue) IsEmpty() bool {
	return pq.Size <= 0
}

// IsFull 优先权队列是否满了
func (pq *PQueue) IsFull() bool {
	return pq.Size >= pq.MaxHeap
}

// Append 将元素加入队列
func (pq *PQueue) Append(x BnElement[int]) {
	if pq.IsFull() {
		log.Fatalf("Overflow")
	}
	pq.Size++
	pq.Elements[pq.Size] = x
	//调整堆
	AdjustUp(pq.Elements, pq.Size)
}

// Serve 返回最高优先权的元素值，并从队列中删除该元素
func (pq *PQueue) Serve() base_interface.Weight {
	if pq.IsEmpty() {
		log.Fatalf("Underflow")
	}
	var x = pq.Elements[1]
	//调整堆
	pq.Elements[1] = pq.Elements[pq.Size]
	pq.Size--
	AdjustDown(pq.Elements, 1, pq.Size)
	return x
}
