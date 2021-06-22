package innersort

import "testing"

func TestMergeSort(t *testing.T) {
	var lst = makeList(10, t)
	t.Log("合并排序前", lst.Elements)
	lst.MergeSort()
	t.Log("合并排序后", lst.Elements)
}

func TestLinkListMergeSort(t *testing.T) {
	var lst = makeLinkList()
	t.Log("单链表两路合并函数排序前", handleLinkList(*lst))
	RMergeSort(lst)
	t.Log("单链表两路合并函数排序后", handleLinkList(*lst))
}
