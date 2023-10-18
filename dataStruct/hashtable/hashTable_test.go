package hashtable

import (
	"log"
	"testing"
)

// TestHashTable 测试散列表
func TestHashTable(t *testing.T) {
	var hashTB = new(HashTable)
	var divitor = 6
	//初始化
	CreateHashTable(hashTB, divitor)
	var insertArr = []Entry{
		{Key: 1, Data: 1},
		{Key: 2, Data: 2},
		{Key: 3, Data: 3},
		{Key: 4, Data: 4},
		{Key: 5, Data: 5},
		{Key: 15, Data: 15},
	}
	for _, v := range insertArr {
		var insert = hashTB.Insert(v)
		if !insert {
			t.Error("插入失败", v)
			continue
		}
	}
	loopHashtable(hashTB)
	var x = new(Entry)
	//搜索
	var searchKey KeyType = 15
	var search = hashTB.Search(searchKey, x)
	if !search {
		t.Errorf("没有找到:%d", searchKey)
	}
	t.Log("散列表搜索结果:", x)

	//删除
	var del = hashTB.Delete(searchKey, x)
	if !del {
		t.Errorf("删除失败:%d", searchKey)
	}
	loopHashtable(hashTB)
	t.Log("散列表删除结果:", x)

	//插入
	var newInsert = hashTB.Insert(*x)
	if !newInsert {
		t.Error("插入失败", *x)
	}
	loopHashtable(hashTB)
}

func loopHashtable(hashTB *HashTable) {
	//遍历散列表
	for index, element := range hashTB.t {
		log.Printf("散列表index:%d, element:%v", index, element)
	}
}
