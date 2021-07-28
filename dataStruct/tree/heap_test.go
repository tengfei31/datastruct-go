package tree

import (
	"sync"
	"testing"
)

func TestMinHeap_HeapSort(t *testing.T) {
	var hp = makeMinHeap()
	t.Log("最小堆排序前", hp.Elements)
	hp.HeapSort()
	t.Log("最小堆排序后", hp.Elements)
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
			}
		}
		wd.Done()
	}()
	//出发func 1
	notify2 <- struct{}{}
	wd.Wait()
	t.Log("main exit")
}
