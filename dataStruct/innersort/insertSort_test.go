package innersort

import "testing"

//makeLinkList 生成LinkList
func makeLinkList() *LinkList {
	lst := NewLinkList()
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

func TestInsertSortOrder(t *testing.T) {
	var lst = makeList(10, t)
	t.Log("顺序表上的直接插入排序前", lst.Elements)
	InsertSortOrder(lst)
	t.Log("顺序表上的直接插入排序后", lst.Elements)
}

func TestInsertSortLinkList(t *testing.T) {
	var lst = makeLinkList()
	t.Log("链表上的直接插入排序前", (*lst).handleLinkList())
	InsertSortLinkList(lst)
	t.Log("链表上的直接插入排序后", (*lst).handleLinkList())
}

func TestInsertSortShellSort(t *testing.T) {
	var lst = makeList(10, t)
	t.Log("希尔排序前", lst.Elements)
	InsertSortShellSort(lst)
	t.Log("希尔排序后", lst.Elements)
}
