package hashtable

import (
	"log"
	"math"
	"math/rand"
	"time"
)

type KeyType int
type DataType int

type Entry struct {
	Key  KeyType
	Data DataType
}
type T Entry

type SkipNode struct {
	Element T
	Link    []*SkipNode
}

//NewSkipNode 初始化node
func NewSkipNode(lev int) *SkipNode {
	//var p []SkipNode
	//p = make([]SkipNode, lev + 1)
	//p := &SkipNode{}
	return &SkipNode{}
}

type SkipList struct {
	MaxLevel   int
	Level      int
	TailKey    KeyType
	Head, Tail *SkipNode
	Last       []*SkipNode
}

func NewSkipList() *SkipList {
	return &SkipList{}
}

//CreateSkipList 初始化跳表
func (sl *SkipList) CreateSkipList(maxNum KeyType, maxLev int) {
	var i int
	_ = i
	sl.MaxLevel = maxLev
	sl.Level = 0
	sl.TailKey = maxNum
	sl.Head = NewSkipNode(maxLev)
	sl.Tail = NewSkipNode(0)
	sl.Last = make([]*SkipNode, 0, maxLev+1)
	sl.Tail.Element.Key = maxNum
	for i = 0; i <= maxLev; i++ {
		sl.Head.Link[i] = sl.Tail
	}
}

//CountLevel 计算跳表的层级
func (sl *SkipList) CountLevel() int {
	var lev int
	//随机数
	rand.Seed(time.Now().UnixNano())
	for rand.Intn(math.MaxInt32) < (math.MaxInt32 / 2) {
		lev++
	}
	if lev <= sl.MaxLevel {
		return lev
	}
	return sl.MaxLevel
}

//Search 搜索某一个key
func (sl *SkipList) Search(k KeyType, x *T) bool {
	return false
}

//SaveSearch 查找某一个key的位置
func (sl *SkipList) SaveSearch(k KeyType) *SkipNode {
	return nil
}

//Insert 将key插入到跳表中
func (sl *SkipList) Insert() {}

//Delete 删除跳表中某一个key
func (sl *SkipList) Delete(k KeyType, x *T) bool {
	var p *SkipNode
	p = new(SkipNode)
	var i int
	if k >= sl.TailKey {
		log.Println("bad input")
		return false
	}
	p = sl.SaveSearch(k)
	if p.Element.Key != k {
		log.Println("p.Element.Key != k")
		return false
	}
	for i = 0; i <= sl.Level && sl.Last[i].Link[i] == p; i++ {
		sl.Last[i].Link[i] = p
	}
	for sl.Level > 0 && sl.Head.Link[sl.Level] == sl.Tail {
		sl.Level--
	}
	x = &(p.Element)
	return true
}
