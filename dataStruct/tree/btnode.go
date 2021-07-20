package tree

import (
	"fmt"
	"math"
)

//T 节点元素值
//TODO:这里需要重新设计，因为在tree里面引用graph，在graph里引用了tree，编译会报错，这里只是示例意思
type T Point

// BTNode 节点
//	LTag RTag 链接上下级节点
//	LChild RChild 左右子树
//	Element 元素值
type BTNode struct {
	Element        T
	LChild, RChild *BTNode
	LTag, RTag     int
}

//NewNode 创建新的节点
func NewNode() *BTNode {
	return new(BTNode)
}

//NewNode2 构造新节点
func NewNode2(x T) *BTNode {
	var p = new(BTNode)
	p.Element = x
	p.LChild = nil
	p.RChild = nil
	return p
}

//Visit 打印每个节点元素
func Visit(p *BTNode) {
	fmt.Printf("%d\n", p.Element.GetWeight())
}

//PreOrd 前序遍历
func (node *BTNode) PreOrd() {
	if node != nil {
		Visit(node)
		node.LChild.PreOrd()
		node.RChild.PreOrd()
	}
}

//InOrd 中序遍历
func (node *BTNode) InOrd() {
	if node != nil {
		node.LChild.InOrd()
		Visit(node)
		node.RChild.InOrd()
	}
}

//PostOrd 后序遍历
func (node *BTNode) PostOrd() {
	if node != nil {
		node.LChild.PostOrd()
		node.RChild.PostOrd()
		Visit(node)
	}
}

//Size 二叉树的节点数
func (node *BTNode) Size() int {
	var s, s1, s2 int
	if node == nil {
		return 0
	} else {
		s1 = node.LChild.Size()
		s2 = node.RChild.Size()
		s = s1 + s2 + 1
		return s
	}
}

//Size1 二叉树的节点数
func (node *BTNode) Size1() int {
	if node == nil {
		return 0
	} else {
		return node.LChild.Size1() + node.RChild.Size1() + 1
	}
}

//Depth 计算二叉树的高度
func (node *BTNode) Depth() int {
	if node == nil {
		return 0
	} else {
		return int(1 + math.Max(float64(node.LChild.Depth()), float64(node.RChild.Depth())))
	}
}

//CopyBT 复制二叉树
func (node *BTNode) CopyBT() *BTNode {
	var q *BTNode
	if node == nil {
		return nil
	}
	q = NewNode()
	q.Element = node.Element
	q.LChild = node.LChild.CopyBT()
	q.RChild = node.RChild.CopyBT()
	return q
}

//二叉线索树

//MakeThread TODO:还有问题，需要修改 构造中序线索树
func (node *BTNode) MakeThread(ppr **BTNode) {
	var t = node
	if t != nil {
		node.LChild.MakeThread(ppr)
		if *ppr != nil {
			if (*ppr).RChild == nil {
				(*ppr).RTag = 1
				(*ppr).RChild = t
			} else {
				(*ppr).RTag = 0
			}
		}
		if t.LChild == nil {
			t.LTag = 1
			t.LChild = *ppr
		} else {
			t.LTag = 0
		}
		ppr = &t
		node.RChild.MakeThread(ppr)
	}
}

//Next 找到node指定节点的后继结点
func (node *BTNode) Next() *BTNode {
	var q = node.RChild
	if node.RTag == 0 {
		for q.LTag != 0 {
			q = q.LChild
		}
	}
	return q
}
