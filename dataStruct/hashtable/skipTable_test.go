package hashtable

import (
	"testing"
)

func TestSkipTable(t *testing.T) {
	var maxKey = KeyType(10)
	var maxLev = 5
	var skipTable = NewSkipList(maxKey, maxLev)
	// 生成数据
	var i KeyType
	for i = 0; i < maxKey; i++ {
		var entry = Entry{
			Key:  i,
			Data: DataType(i),
		}
		skipTable.Insert(entry)
	}
	// 搜索一个数据
	var searchEntry = new(Entry)
	var searchRes = skipTable.Search(8, searchEntry)
	t.Log("搜索结果：", searchRes)
	if searchRes {
		t.Log("搜索成功：", searchEntry)
	}
	//删除一个数据
	var delEntry = new(Entry)
	if skipTable.Delete(9, delEntry) {
		t.Log("删除成功")
	} else {
		t.Log("删除失败")
	}
	//再次搜索
	if skipTable.Search(9, searchEntry) {
		t.Log("搜索成功：", searchEntry)
	} else {
		t.Log("搜索失败")
	}
}
