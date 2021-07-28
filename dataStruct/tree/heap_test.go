package tree

import "testing"

func TestMinHeap_HeapSort(t *testing.T) {
	var hp = makeMinHeap()
	t.Log("最小堆排序前", hp.Elements)
	hp.HeapSort()
	t.Log("最小堆排序后", hp.Elements)

}
