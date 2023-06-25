package tree

import (
	"fmt"
	"log"
	"testing"
)

// TestBtree 测试二叉树
func TestBtree(t *testing.T) {
	var a, x, y, z Btree[int]
	//var e tree.T
	a.CreateBT()
	x.CreateBT()
	y.CreateBT()
	z.CreateBT()
	y.MakeBT(BnElement[int]{W: 'E'}, &a, &a)
	z.MakeBT(BnElement[int]{W: 'F'}, &a, &a)
	x.MakeBT(BnElement[int]{W: 'C'}, &y, &z)
	y.MakeBT(BnElement[int]{W: 'D'}, &a, &a)
	z.MakeBT(BnElement[int]{W: 'B'}, &y, &x)

	fmt.Println("先序遍历")
	//z.PreOrder()
	z.IPreOrder()
	fmt.Println("中序遍历")
	z.InOrder()
	//z.IInOrder()
	//z.BreakBT(&e, &y, &x)
	//fmt.Println("中序遍历二叉线索树")
	//z.BuildThreadBT()
	//z.TInOrder()
	fmt.Println("后序遍历")
	z.PostOrder()

}

func TestBtSearch(t *testing.T) {
	var bt = new(Btree[int])
	//二叉搜索树的插入
	for _, num := range []int{28, 21, 25, 36, 33, 43} {
		var element = BnElement[int]{W: num}
		BtInsert(bt, element)
	}
	//中序遍历bt
	bt.InOrder()

	var x *BnElement[int] = new(BnElement[int])
	var res bool = BtSearch(*bt, 21, x)

	if res {
		//二叉搜索树删除节点
		res = BtRemove(bt, 21, x)
		if !res {
			log.Printf("btremove is false")
			return
		}
	}
	//中序遍历bt
	bt.InOrder()
	fmt.Printf("%+v\n", bt.Root)
	fmt.Printf("%+v\n", *x)
}
