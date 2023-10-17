package hashtable

//HashTable 散列表

// NeverUsed 表示这个位置没有被使用
const NeverUsed KeyType = 0

const (
	UnderFlow int = iota
	OverFlow
	Success
	Duplicate
	NotPresent
)

// HashNode 节点
type HashNode struct {
	Empty   bool  //该节点是否为空
	Element Entry //该节点具体值
}

// HashTable 散列表
type HashTable struct {
	M int         // 散列表大小
	t []*HashNode // 散列表集合
}

//构建线性探查散列表

// NewArray 创建空节点数组
func NewArray(len int) []*HashNode {
	return make([]*HashNode, len)
}

// CreateHashTable 创建空的线性探查散列表
func CreateHashTable(htb *HashTable, divitor int) {
	var i int
	if htb == nil {
		htb = new(HashTable)
	}
	htb.M = divitor
	htb.t = NewArray(htb.M)
	for i = 0; i < htb.M; i++ {
		var tmp = new(HashNode)
		tmp.Empty = true
		tmp.Element.Key = NeverUsed
		htb.t[i] = tmp
		//htb.t[i] = new(HashNode)
		//htb.t[i].Empty = true
		//htb.t[i].Element.Key = NeverUsed
	}
}

// HSearch 线性探查散列表的搜索
func (htb *HashTable) HSearch(k KeyType, pos *int) int {
	var i, j int
	*pos = int(k) % htb.M
	//计算基地址
	if *pos < 0 {
		*pos = htb.M + *pos
	}
	//记录初始位置
	i = *pos
	//表示未找到空值的位置
	j = -1
	for {
		//首次遇到空值的位置
		if htb.t[*pos].Element.Key == NeverUsed && j == -1 {
			j = *pos
		}
		//表中没有关键字值为K的元素
		if htb.t[*pos].Empty {
			break
		}
		//搜索成功
		if htb.t[*pos].Element.Key == k {
			return Success
		}
		//如果找不到就向下一个位置寻找
		*pos = (*pos + 1) % htb.M

		if *pos == i {
			break
		}
	}
	//散列表满了
	if j == -1 {
		return OverFlow
	}
	//设置首次遇到的空值位置，并返回
	*pos = j
	return NotPresent
}

// Search  线性探查散列表的搜索
func (htb *HashTable) Search(k KeyType, x *Entry) bool {
	var pos = new(int)
	result := htb.HSearch(k, pos)
	if result == Success {
		*x = htb.t[*pos].Element
		return true
	}
	return false
}

// Insert 线性探查散列表的插入
func (htb *HashTable) Insert(x Entry) bool {
	var pos = new(int)
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

// Delete 线性探查散列表的删除
func (htb *HashTable) Delete(k KeyType, x *Entry) bool {
	var pos = new(int)
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
