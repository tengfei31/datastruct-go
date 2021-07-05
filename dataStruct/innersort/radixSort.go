package innersort

import (
	"math/rand"
	"time"
)

//基数排序

//d 每个节点存储d个key
const d = 10

//radix 需要分配radix个链表结构
const radix = 10

//RNode 节点存储类型
type RNode struct {
	Key  []T
	Link int
}

//SList 静态链表类型
type SList struct {
	Size, Head int
	Nodes      []RNode
}

//NewSList 创建链表
func NewSList(size int) *SList {
	var sList = &SList{
		Size:  0,
		Head:  0,
		Nodes: make([]RNode, size),
	}
	return sList
}

//MakeSList 生成链表
func MakeSList() *SList {
	var lst = NewSList(10)
	for i := 0; i < len(lst.Nodes); i++ {
		var tmpNode = RNode{
			Key:  make([]T, d),
			Link: -1,
		}
		for j := 0; j < d; j++ {
			rand.Seed(time.Now().UnixNano() + int64(j) + int64(i))
			tmpNode.Key[j] = T{K: rand.Intn(9)}
		}
		lst.Nodes[i] = tmpNode
		lst.Size++
	}
	return lst
}

//RadixSort 基数排序
func (lst *SList) RadixSort() {
	var (
		i, j, k, tail int
	)
	var f = make([]int, radix)
	var r = make([]int, radix)
	for i = d - 1; i >= 0; i-- {
		//分配
		for j = 0; j < radix; j++ {
			f[j] = -1
		}

		for k = lst.Head; k > -1; k = lst.Nodes[k].Link {
			j = lst.Nodes[k].Key[i].GetKey()
			if f[j] == -1 {
				f[j] = k
			} else {
				lst.Nodes[r[j]].Link = k
			}
			r[j] = k
		}
		//收集
		for j = 0; f[j] == -1; j++ {
		}

		lst.Head = f[j]
		tail = r[j]
		for j < radix {
			for j++; j < radix; j++ {
				if f[j] > -1 {
					lst.Nodes[tail].Link = f[j]
					tail = r[j]
				}
			}
		}
		lst.Nodes[tail].Link = -1
	}
}
