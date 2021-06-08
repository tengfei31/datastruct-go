package innersort

import (
	"testing"
)

//makeList 生成list
func makeList(maxList int, t *testing.T) *List {
	defer func() {
		if err := recover(); err != nil {
			t.Fatal(err)
		}
	}()
	var lst = NewList(maxList)
	lst.AddElement(T{K: 20})
	lst.AddElement(T{K: 10})
	lst.AddElement(T{K: 31})
	lst.AddElement(T{K: 34})
	lst.AddElement(T{K: 1})
	lst.AddElement(T{K: 4})
	lst.AddElement(T{K: 0})
	lst.AddElement(T{K: 15})
	lst.AddElement(T{K: 100})
	lst.AddElement(T{K: 89})
	return lst
}

func TestBubbleSort(t *testing.T) {
	var lst = makeList(10, t)
	t.Log("冒泡排序前", lst.Elements)
	BubbleSort(lst)
	t.Log("冒泡排序后", lst.Elements)
}

func TestBubbleSort1(t *testing.T) {
	var lst = makeList(10, t)
	t.Log("冒泡排序前", lst.Elements)
	BubbleSort1(lst)
	t.Log("冒泡排序后", lst.Elements)
}

func TestQuickSort(t *testing.T) {
	var lst = makeList(10, t)
	t.Log("快速排序前", lst.Elements)
	QuickSort(lst)
	t.Log("快速排序后", lst.Elements)
}
