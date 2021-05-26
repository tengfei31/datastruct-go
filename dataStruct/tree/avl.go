package tree

import (
	"log"
)

//二叉平衡树

type K int

//AVLNode AVL二叉平衡树结点
type AVLNode struct {
	Element        T
	Bf             int
	LChild, RChild *AVLNode
}

//AVLBTree 二叉平衡树
type AVLBTree struct {
	Root *AVLNode
}

//NewNode2 创建结点
func (avl *AVLNode) NewNode2(x T) *AVLNode {
	avl = &AVLNode{
		Element: x,
		Bf:      0,
		LChild:  nil,
		RChild:  nil,
	}
	return avl
}

//LeftRotation 二叉平衡树左旋转函数
func LeftRotation(s **AVLNode, unbalanced *bool) {
	var u, r *AVLNode
	r = (*s).LChild
	if r.Bf == 1 {
		(*s).LChild = r.RChild
		r.RChild = *s
		(*s).Bf = 0
		*s = r
	} else {
		u = r.RChild
		r.RChild = u.LChild
		u.LChild = r
		(*s).LChild = u.RChild
		u.RChild = *s
		switch u.Bf {
		case 1:
			(*s).Bf = -1
			r.Bf = 0
			break
		case 0:
			(*s).Bf = 0
			r.Bf = 0
			break
		case -1:
			(*s).Bf = 0
			r.Bf = 1
			break
		}
		*s = u
	}
	(*s).Bf = 0
	*unbalanced = false
}

//RightRotation 二叉平衡树右旋转函数
func RightRotation(s **AVLNode, unbalanced *bool) {
	var u, r *AVLNode
	r = (*s).RChild
	if r.Bf == -1 {
		(*s).RChild = r.LChild
		r.LChild = *s
		(*s).Bf = 0
		*s = r
	} else {
		u = r.LChild
		r.LChild = u.RChild
		u.RChild = r
		(*s).RChild = u.LChild
		u.LChild = *s
		switch u.Bf {
		case 1:
			(*s).Bf = 0
			r.Bf = -1
			break
		case 0:
			(*s).Bf = 0
			r.Bf = 0
			break
		case -1:
			(*s).Bf = 1
			r.Bf = 0
			break
		}
		*s = u
	}
	(*s).Bf = 0
	*unbalanced = false
}

//AVLIst 二叉平衡树的插入
func AVLIst(p **AVLNode, x T, unbalanced *bool) bool {
	var result bool = true
	if *p == nil {
		*unbalanced = true
		*p = (*p).NewNode2(x)
	} else if x.GetWeight() < (*p).Element.GetWeight() {
		result = AVLIst(&(*p).LChild, x, unbalanced)
		if *unbalanced {
			switch (*p).Bf {
			case -1:
				(*p).Bf = 0
				*unbalanced = false
				break
			case 0:
				(*p).Bf = 1
				break
			case 1:
				LeftRotation(p, unbalanced)
				break
			}
		}
	} else if x == (*p).Element {
		*unbalanced = false
		result = false
		log.Println("The key is already in the tree")
	} else {
		result = AVLIst(&(*p).RChild, x, unbalanced)
		if *unbalanced {
			switch (*p).Bf {
			case 1:
				(*p).Bf = 0
				*unbalanced = false
				break
			case 0:
				(*p).Bf = -1
				break
			case -1:
				RightRotation(p, unbalanced)
			}
		}
	}
	return result
}

//AVLInsert 二叉平衡树的插入
func AVLInsert(bt *AVLBTree, x T) bool {
	var unbalanced bool
	return AVLIst(&bt.Root, x, &unbalanced)
}
