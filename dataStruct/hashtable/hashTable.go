package hashtable

//HashTable 散列表

//NeverUsed 表示这个位置没有被使用
const NeverUsed KeyType = 1

//
const (
	UnderFlow int = iota
	OverFlow
	Success
	Duplicate
	NotPresent
)

//HashNode 节点
type HashNode struct {
	Empty   bool
	Element T
}

//HashTable 散列表
type HashTable struct {
	M int
	t []*HashNode
}

//构建线性探查散列表

//NewArray 创建空节点数组
func NewArray(len int) []*HashNode {
	return make([]*HashNode, 0, len)
}

//CreateHashTable 创建空的散列表
func CreateHashTable(htb *HashTable, divitor int) {
	var i int
	htb.M = divitor
	htb.t = NewArray(htb.M)
	for i = 0; i < htb.M; i++ {
		htb.t[i].Empty = true
		htb.t[i].Element.Key = NeverUsed
	}
}

//HSearch 线性探查散列表的搜索
func (htb *HashTable) HSearch(k KeyType, pos *int) int {
	var i, j int
	*pos = int(k) % htb.M
	//计算基地址
	if *pos < 0 {
		*pos = htb.M + *pos
	}
	i = *pos
	//表示未找到空值的位置
	j = -1
	for *pos != i {
		//首次遇到空值的位置
		if htb.t[*pos].Element.Key == NeverUsed && j == -1 {
			j = *pos
		}
		//表中没有关键字值为K的元素
		if htb.t[*pos].Empty == true {
			break
		}
		//搜索成功
		if htb.t[*pos].Element.Key == k {
			return Success
		}
		*pos = (*pos + 1) % htb.M
	}
	//散列表满了
	if j == -1 {
		return OverFlow
	}
	//设置首次遇到的空值位置，并返回
	*pos = j
	return NotPresent
}

//Search  线性探查散列表的搜索
func (htb *HashTable) Search(k KeyType, x *T) bool {
	var pos *int
	pos = new(int)
	result := htb.HSearch(k, pos)
	if result == Success {
		*x = htb.t[*pos].Element
		return true
	}
	return false
}

//Insert 线性探查散列表的插入
func (htb *HashTable) Insert(x T) bool {
	var pos *int = new(int)
	result := htb.HSearch(x.Key, pos)
	//如果原表未满且不包含重复元素
	if result == NotPresent {
		htb.t[*pos].Element = x
		htb.t[*pos].Empty = false
		//插入成功
		return true
	}
	//原表存在重复元素或散列表已满，插入失败
	return false
}

//Delete 线性探查散列表的删除
func (htb *HashTable) Delete(k KeyType, x *T) bool {
	var pos *int = new(int)
	result := htb.HSearch(k, pos)
	//搜索到，删除成功
	if result == Success {
		*x = htb.t[*pos].Element
		htb.t[*pos].Element.Key = NeverUsed
		return true
	}
	//删除失败
	return false
}
