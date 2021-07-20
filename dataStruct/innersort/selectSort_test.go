package innersort

import (
	"context"
	"testing"
	"time"
	"unsafe"
)

func TestSelectSort(t *testing.T) {
	var lst = makeList(10, t)
	t.Log("简单选择排序前", lst.Elements)
	lst.SelectSort()
	t.Log("简单选择排序后", lst.Elements)
}

func IsLittleEndian() bool {
	var val int32 = 1 // 占4byte 转换成16进制 0x00 00 00 01
	// 大端(16进制)：00 00 00 01
	// 小端(16进制)：01 00 00 00
	var pointer = unsafe.Pointer(&val)
	var p = (*byte)(pointer)
	if *p == 1 {
		return true
	}
	return false
}

func TestEndian(t *testing.T) {
	t.Log(IsLittleEndian())
}

func TestContext(t *testing.T) {
	var start = time.Now()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			//time.Sleep(time.Second * 1)
			select {
			case <-ctx.Done():
				t.Log("child exit", i)
			}
		}(i, ctx)
	}
	select {
	case <-ctx.Done():
		t.Log("main取消了")
	}
	var diff = time.Since(start)
	t.Log("执行时间：", diff)
}
