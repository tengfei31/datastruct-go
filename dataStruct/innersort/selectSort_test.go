package innersort

import "testing"

func TestSelectSort(t *testing.T) {
	var lst = makeList(10, t)
	t.Log("简单选择排序前", lst.Elements)
	lst.SelectSort()
	t.Log("简单选择排序后", lst.Elements)
}
