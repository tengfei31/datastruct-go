package tree

import (
	"fmt"
	"log"
	"testing"
)

func TestBtSearch(t *testing.T) {
	var bt *Btree = new(Btree)
	BtInsert(bt, BnElement[int]{W: 28})
	BtInsert(bt, BnElement[int]{W: 21})
	BtInsert(bt, BnElement[int]{W: 25})
	BtInsert(bt, BnElement[int]{W: 36})
	BtInsert(bt, BnElement[int]{W: 33})
	BtInsert(bt, BnElement[int]{W: 43})
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
	fmt.Println(*x)
}
