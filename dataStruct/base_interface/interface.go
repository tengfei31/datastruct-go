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
