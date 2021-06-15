package innersort

import "testing"

func TestMergeSort(t *testing.T) {
	var lst = makeList(10, t)
	t.Log("合并排序前", lst.Elements)
	MergeSort(lst)
	t.Log("合并排序后", lst.Elements)
}
