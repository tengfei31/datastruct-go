package tree

import (
	"fmt"
	"log"
	"testing"
)

func TestBtSearch(t *testing.T) {
	var bt *Btree = new(Btree)
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
