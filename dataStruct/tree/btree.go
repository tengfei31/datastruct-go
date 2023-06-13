package tree

// Btree 树
type Btree[T int] struct {
	Root *BTNode[T]
}

// CreateBT 创建空的二叉树
func (bt *Btree[T]) CreateBT() {
	bt.Root = nil
}

// MakeBT 构建二叉树
func (bt *Btree[T]) MakeBT(x BnElement[T], lt *Btree[T], rt *Btree[T]) {
	var p *BTNode[T] = NewNode[T]()
	p.Element = x
	p.LChild = lt.Root
	p.RChild = rt.Root
	lt.Root = nil
	rt.Root = nil
	bt.Root = p
}

// BreakBT 置空二叉树,并返回二叉树的左右节点和跟元素
func (bt *Btree[T]) BreakBT(x *BnElement[T], lt *Btree[T], rt *Btree[T]) {
	var p *BTNode[T] = bt.Root
	if p != nil {
		*x = p.Element
		lt.Root = p.LChild
		rt.Root = p.RChild
		bt.Root = nil
		p = nil
	}
}

// PreOrder 前序遍历
func (bt *Btree[T]) PreOrder() {
	bt.Root.PreOrd()
}

// IPreOrder 前序便利
func (bt *Btree[T]) IPreOrder() {
	var s []*BTNode[T]
	var p *BTNode[T] = bt.Root
	//s = make([]*BTNode, 10)
	for p != nil {
		Visit(p)
		if p.RChild != nil {
			s = append(s, p.RChild)
		}
		if p.LChild != nil {
			p = p.LChild
		} else {
			//将栈顶的赋给指针p
			if s != nil {
				p = s[len(s)-1]
			}
			//弹出栈顶元素
			s = append(s, nil)
		}
	}
}

// InOrder 中序遍历
func (bt *Btree[T]) InOrder() {
	bt.Root.InOrd()
}

// IInOrder 中序遍历 TODO: 还有问题，需要修改
func (bt *Btree[T]) IInOrder() {
	var s []*BTNode[T]
	var p *BTNode[T]
	p = bt.Root
	for p != nil || s != nil {
		if p != nil {
			if s != nil {
				p = s[len(s)-1]
			}
			s = append(s, nil)
			Visit(p)
			p = p.RChild
		} else {
			s = append(s, p)
			p = p.LChild
		}
	}
}

// PostOrder 后序遍历
func (bt *Btree[T]) PostOrder() {
	bt.Root.PostOrd()
}

// SizeBT 二叉树的节点数
func (bt *Btree[T]) SizeofBT() int {
	return bt.Root.Size()
}

// DepthofBT 计算二叉树的高度
func (bt *Btree[T]) DepthofBT() int {
	return bt.Root.Depth()
}

// CopyofBT 复制二叉树
func (bt *Btree[T]) CopyofBT() Btree[T] {
	var tree Btree[T]
	tree.Root = bt.Root.CopyBT()
	return tree
}

//线索二叉树

// BuildThreadBT 构建中序线索树
func (bt *Btree[T]) BuildThreadBT() {
	var pr *BTNode[T] = new(BTNode[T])
	if bt.Root != nil {
		//pr = nil
		bt.Root.MakeThread(&pr)
		pr.RTag = 1
	}
}

// GoFirst 第一个叶子节点
func (bt *Btree[T]) GoFirst() *BTNode[T] {
	var p *BTNode[T] = bt.Root
	if p != nil {
		for p.LChild != nil {
			p = p.LChild
		}
	}
	return p
}

// TInOrder 中序遍历二叉线索树
func (bt *Btree[T]) TInOrder() {
	var p *BTNode[T] = bt.GoFirst()
	for p != nil {
		Visit(p)
		p = p.Next()
	}
}

//var ht [MaxSize]Btree

// CreateHFMTree 返回构造成功的哈夫曼树
func CreateHFMTree[T int](w []BnElement[T], n int) Btree[T] {
	var zero Btree[T]
	var ht [MaxSize]Btree[T]
	var i, k, k1, k2 int
	zero.CreateBT()
	for i = 0; i < n; i++ {
		ht[i].MakeBT(w[i], &zero, &zero)
	}
	for k = n - 1; k > 0; k-- {
		Fmin(ht, &k1, &k2, k+1)
		x := BnElement[T]{}
		x.SetWeight(ht[k1].Root.Element.GetWeight() + ht[k2].Root.Element.GetWeight())
		ht[k1].MakeBT(x, &ht[k1], &ht[k2])
		ht[k2] = ht[k]
	}
	return ht[0]
}

// Fmin 返回k棵二叉树的最小和次最小的二叉树在数组ht中的下标
func Fmin[T int](ht [MaxSize]Btree[T], k1 *int, k2 *int, k int) {
	var min1 = ht[0].Root.Element.GetWeight()
	*k1 = 0
	var min2 = ht[1].Root.Element.GetWeight()
	*k2 = 1
	var i int
	for i = 2; i < k; i++ {
		var tmpElement = ht[i].Root.Element.GetWeight()
		if tmpElement < min2 {
			if tmpElement > min1 {
				min2 = tmpElement
				*k2 = i
			} else {
				min2 = min1
				*k2 = *k1
				min1 = tmpElement
				*k1 = i
			}
		}
	}
	//最后比较min1和min2，检查是否交换位置
	if min1 > min2 {
		min1, min2 = min2, min1
		*k1, *k2 = *k2, *k1
	}
}
