package hashtable

import (
	"log"
	"math"
	"math/rand"
)

type SkipListInterface interface {
	CreateSkipList(maxNum KeyType, maxLev int)
	Search(k KeyType, x *Entry) bool
	saveSearch(k KeyType) *SkipNode
	Insert(x Entry) bool
	Delete(k KeyType, x *Entry) bool
	CountLevel() int
}

type DataTypeInterface interface {
	GetData() Entry
	SetData(data DataType)
}

// KeyType key的数据类型
type KeyType int

// DataType 数据的类型
type DataType int

// Entry 数据源
type Entry struct {
	Key  KeyType
	Data DataType
}

// SkipNode 跳表结点
type SkipNode struct {
	Element Entry
	Link    []*SkipNode
}

// NewSkipNode 初始化node
func NewSkipNode(lev int) *SkipNode {
	var p = new(SkipNode)
	p.Link = make([]*SkipNode, lev, lev)
	return p
}

// SkipList 跳表结构
// MaxLevel 最大层级
// Level 当前层级数
// TailKey 尾部的key,最大的key
type SkipList struct {
	MaxLevel, Level int
	TailKey         KeyType
	Head, Tail      *SkipNode
	Last            []*SkipNode
}

// NewSkipList 创建跳表
func NewSkipList(maxNum KeyType, maxLev int) *SkipList {
	var skipList = new(SkipList)
	skipList.CreateSkipList(maxNum, maxLev)
	return skipList
}

// CreateSkipList 初始化跳表
// maxNum 最大的key的值
// maxLev 最大的层级
func (sl *SkipList) CreateSkipList(maxNum KeyType, maxLev int) {
	var i int
	sl.MaxLevel = maxLev
	sl.Level = 0
	sl.TailKey = maxNum
	sl.Head = NewSkipNode(maxLev)
	sl.Tail = NewSkipNode(0)
	sl.Last = make([]*SkipNode, maxLev+1, maxLev+1)
	sl.Tail.Element.Key = maxNum
	for i = 0; i < maxLev; i++ {
		sl.Head.Link[i] = sl.Tail
		//sl.Head.Link = append(sl.Head.Link, sl.Tail)
	}
}

// CountLevel 计算跳表的层级
func (sl *SkipList) CountLevel() int {
	var lev int
	//随机数
	rand.NewSource(1)
	for rand.Intn(math.MaxInt16) <= (math.MaxInt16 / 2) {
		lev++
	}
	if lev <= sl.MaxLevel {
		return lev
	}
	return sl.MaxLevel
}

// Search 搜索某一个key
func (sl *SkipList) Search(k KeyType, x *Entry) bool {
	if k >= sl.TailKey {
		return false
	}
	var p = sl.Head
	for i := sl.Level; i >= 0; i-- {
		for p.Link[i].Element.Key < k {
			p = p.Link[i]
		}
	}
	*x = p.Link[0].Element
	return x.Key == k
}

// SaveSearch 查找某一个key的位置，为了实现插入和删除设计的私有函数
func (sl *SkipList) saveSearch(k KeyType) *SkipNode {
	if k >= sl.TailKey {
		return nil
	}
	var p = sl.Head
	for i := sl.Level; i >= 0; i-- {
		for p.Link[i].Element.Key < k {
			p = p.Link[i]
		}
		sl.Last[i] = p
		//if len(sl.Last) <= 0 {
		//	sl.Last = append(sl.Last, p)
		//} else {
		//	sl.Last[i] = p
		//}
	}
	return p.Link[0]
}

// Insert 将key插入到跳表中
func (sl *SkipList) Insert(x Entry) bool {
	var p, y *SkipNode
	var lev, i int
	if x.Key >= sl.TailKey {
		log.Printf("bad input")
		return false
	}
	p = sl.saveSearch(x.Key)
	if p == nil {
		return false
	}
	if p.Element.Key == x.Key {
		log.Printf("duplicate")
		return false
	}
	lev = sl.CountLevel()
	if lev > sl.Level {
		sl.Level += 1
		lev = sl.Level
		sl.Last[lev] = sl.Head
		//sl.Last[lev] = sl.Head
	}
	y = NewSkipNode(lev)
	y.Element = x
	for i = 0; i <= lev; i++ {
		if len(y.Link) > i {
			y.Link[i] = sl.Last[i].Link[i]
		} else {
			y.Link = append(y.Link, sl.Last[i].Link[i])
		}
		sl.Last[i].Link[i] = y
	}
	return true
}

// Delete 删除跳表中某一个key
func (sl *SkipList) Delete(k KeyType, x *Entry) bool {
	var p *SkipNode // = new(SkipNode)
	var i int
	if k >= sl.TailKey {
		log.Println("bad input")
		return false
	}
	p = sl.saveSearch(k)
	if p.Element.Key != k {
		log.Println("p.Element.Key != k")
		return false
	}
	for i = 0; i <= sl.Level && sl.Last[i].Link[i] == p; i++ {
		sl.Last[i].Link[i] = p.Link[i]
	}
	for sl.Level > 0 && sl.Head.Link[sl.Level] == sl.Tail {
		sl.Level--
	}
	x = &(p.Element)
	return true
}
