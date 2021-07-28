package innersort

import "testing"

func TestSList_RadixSort(t *testing.T) {
	var lst = MakeSList()
	t.Log("基数排序前", lst.handleSList())
	lst.RadixSort()
	t.Log("基数排序后", lst.handleSList())
}

//handleSList 打印里面的每一个元素
func (lst SList) handleSList() [][]int {
	var tmp = make([][]int, 10)
	for i := 0; i < lst.Size; i++ {
		var tmp1 = make([]int, 10)
		for j := 0; j < len(lst.Nodes[i].Key); j++ {
			tmp1[j] = lst.Nodes[i].Key[j].GetKey()
		}
		tmp[i] = tmp1
	}
	return tmp
}
