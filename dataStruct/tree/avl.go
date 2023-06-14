package tree

import (
	"log"
)

//二叉平衡树

type K int

type Elements[T int] struct {
	W   T
	Val T
}

func (e Elements[T]) GetWeight() T {
	return e.W
}

// AVLNode AVL二叉平衡树结点
type AVLNode[T int] struct {
	Element        Elements[T]
	Bf             int
	LChild, RChild *AVLNode[T]
}

// AVLBTree 二叉平衡树
type AVLBTree[T int] struct {
	Root *AVLNode[T]
}

// NewNode3 创建结点
func NewNode3[T int](x Elements[T]) *AVLNode[T] {
	return &AVLNode[T]{
		Element: x,
		Bf:      0,
		LChild:  nil,
		RChild:  nil,
	}
}

// LeftRotation 二叉平衡树左旋转函数
func LeftRotation[T int](s **AVLNode[T], unbalanced *bool) {
	var u, r *AVLNode[T]
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

// RightRotation 二叉平衡树右旋转函数
func RightRotation[T int](s **AVLNode[T], unbalanced *bool) {
	var u, r *AVLNode[T]
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

// AVLIst 二叉平衡树的插入
func AVLIst[T int](p **AVLNode[T], x Elements[T], unbalanced *bool) bool {
	var result = true
	if *p == nil {
		*unbalanced = true
		*p = NewNode3[T](x)
	} else if x.GetWeight() < (*p).Element.GetWeight() {
		result = AVLIst[T](&(*p).LChild, x, unbalanced)
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
		result = AVLIst[T](&(*p).RChild, x, unbalanced)
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

// AVLInsert 二叉平衡树的插入
func AVLInsert[T int](bt *AVLBTree[T], x Elements[T]) bool {
	var unbalanced = new(bool)
	return AVLIst(&bt.Root, x, unbalanced)
}
