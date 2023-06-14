package tree

import (
	"fmt"
	"testing"
)

// testAVLBTree 测试二叉平衡树
func TestAVLBTree(t *testing.T) {
	var bt = new(AVLBTree[int])
	var arr = []int{45, 28, 12, 14, 23}
	for i := 0; i < len(arr); i++ {
		var x = Elements[int]{
			W: arr[i],
		}
		AVLInsert(bt, x)
	}
	fmt.Printf("%+v\n", bt.Root)
}
