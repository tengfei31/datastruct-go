package tree

import (
	"fmt"
	"sync"
	"testing"
)

func TestMinHeap_HeapSort(t *testing.T) {
	var hp = MakeMinHeap()
	t.Log("最小堆排序前", hp.Elements)
	hp.HeapSort()
	t.Log("最小堆排序后", hp.Elements)
}

// testHeap 测试最小堆
func TestHeap(t *testing.T) {
	var heapQueue *PQueue = new(PQueue)

	heapQueue.CreatePQ(10)
	heapQueue.Append(BnElement[int]{W: 71})
	heapQueue.Append(BnElement[int]{W: 74})
	heapQueue.Append(BnElement[int]{W: 2})
	heapQueue.Append(BnElement[int]{W: 72})
	heapQueue.Append(BnElement[int]{W: 54})
	heapQueue.Append(BnElement[int]{W: 93})
	heapQueue.Append(BnElement[int]{W: 52})
	heapQueue.Append(BnElement[int]{W: 28})
	fmt.Println(heapQueue)
	fmt.Println("1:", heapQueue.Serve())
	fmt.Println("2:", heapQueue.Serve())
}

// TestExchangeGo 两个协程交替打印
func TestExchangeGo(t *testing.T) {
	var notify1 = make(chan struct{}, 1)
	var notify2 = make(chan struct{}, 1)
	var wd sync.WaitGroup
	//1
	wd.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			select {
			case <-notify1:
				t.Log("go func 1")
				notify2 <- struct{}{}
			default:
				t.Log("go func 1 no sign")
			}
		}
		wd.Done()
	}()
	//2
	wd.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			select {
			case <-notify2:
				t.Log("go func 2")
				notify1 <- struct{}{}
			default:
				t.Log("go func 2 no sign")
			}
		}
		wd.Done()
	}()
	//出发func 1
	notify2 <- struct{}{}
	wd.Wait()
	t.Log("main exit")
}
