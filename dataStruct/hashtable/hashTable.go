package hashtable

// HashTable 散列表

// NeverUsed 表示这个位置没有被使用
const NeverUsed KeyType = 0

// 搜索返回结果
const (
	UnderFlow  int = iota // UnderFlow 溢出
	OverFlow              // OverFlow hash表满了
	Success               // Success 搜索成功
	Duplicate             // Duplicate hash表有重复
	NotPresent            // NotPresent hash表里不存在
)

// HashNode 节点
type HashNode struct {
	Empty   bool  //该节点是否为空
	Element Entry //该节点具体值
}

// HashTable 散列表
type HashTable struct {
	M        int         // 散列表大小
	PrimeNum int         //固定一个素数,最好与hash表大小接近
	t        []*HashNode // 散列表集合
}

//构建线性探查散列表

// NewArray 创建空节点数组
func NewArray(len int) []*HashNode {
	return make([]*HashNode, len)
}

// CreateHashTable 创建空的线性探查散列表
func CreateHashTable(htb *HashTable, divitor int) {
	var i int
	htb.M = divitor
	htb.PrimeNum = divitor - 2
	if divitor%2 == 0 {
		htb.PrimeNum = divitor - 1
	}
	htb.t = make([]*HashNode, htb.M)
	for i = 0; i < htb.M; i++ {
		//构造每个节点，并标记该节点为空
		var tmpNode = new(HashNode)
		tmpNode.Empty = true
		//将该节点的关键字值标记为未使用
		tmpNode.Element.Key = NeverUsed
		htb.t[i] = tmpNode
	}
}

// hSearch 线性探查散列表的搜索
func (htb *HashTable) hSearch(k KeyType, pos *int) int {
	var i int  //记录初始位置
	var j = -1 //表示未找到空值的位置
	*pos = int(k) % htb.M
	//计算基地址
	if *pos < 0 {
		*pos = htb.M + *pos
	}
	i = *pos
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
		//方法1、如果找不到就向下一个位置寻找
		*pos = (*pos + 1) % htb.M
		//方法2、伪随机探查法
		//*pos = (*pos + htb.PrimeNum) % htb.M

		//搜索完整个表了
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
	result := htb.hSearch(k, pos)
	if result == Success {
		*x = htb.t[*pos].Element
		return true
	}
	return false
}

// Insert 线性探查散列表的插入
func (htb *HashTable) Insert(x Entry) bool {
	var pos = new(int)
	result := htb.hSearch(x.Key, pos)
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
	result := htb.hSearch(k, pos)
	//搜索到，删除成功
	if result == Success {
		*x = htb.t[*pos].Element
		htb.t[*pos].Element.Key = NeverUsed
		return true
	}
	//删除失败
	return false
}
