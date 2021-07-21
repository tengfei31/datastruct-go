package base_interface

//Weight 设置权重、获取权重，各种自定义类型调用优先权队列的时候，需要实现这两个方法
type Weight interface {
	//GetWeight 获取权重
	GetWeight() int
	//SetWeight 设置权重
	SetWeight(w int)
}

//插入排序

//Order 获取排序元素的接口
type Order interface {
	SetKey(int)
	GetKey() int
}

//二叉树

//KeyType 关键字类型
type KeyType int

type T Weight

//Set 集合
type Set interface {
	//CreateList 创建空集合
	CreateList(maxSize int)
	//IsEmpty 集合是否为空
	IsEmpty() bool
	//IsFull 集合是否满了
	IsFull() bool
	//Search 在集合中搜索关键字值为k的元素，并将该元素放入x中，并返回true
	Search(k KeyType, x *T) bool
	//Insert 在集合中搜索关键字值为k的元素，如果不存在就插入到集合中，并返回true,否则返回false
	Insert(x T) bool
	//Remove 在集合中搜索关键字值为k的元素，如果存在，就将该元素赋值给*x, 并从集合中删除该元素，返回true，否则返回false
	Remove(k KeyType, x *T) bool
}
