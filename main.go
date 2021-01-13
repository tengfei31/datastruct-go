package main

import (
	"bytes"
	"datastruct-go/dataStruct/listarr"
	"datastruct-go/dataStruct/tree"
	"fmt"
	"io"
	"log"
	"os/exec"
	"sort"
)

func main() {
	//testBtree()
	//testHeap()
	//testHFMCode()
	//testUFset()
	testBFSearch()

	//sortArr()
	//handleArr()

	//nginxPipe()
	//createPipe()

}

var arr []int

func InitArr() {
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 2)
	arr = append(arr, 2)
	arr = append(arr, 2)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
}

func handleArr() {
	var tmp []int
	var i int
	var arrLen = len(arr) / 2
	if len(arr)%2 > 0 {
		arrLen += 1
	}
	for i = 0; i < arrLen; i++ {
		tmp = append(tmp, arr[i])
		if i >= len(arr)-(1+i) {
			break
		}
		tmp = append(tmp, arr[len(arr)-(1+i)])
	}
	fmt.Println(arr)
	fmt.Println(tmp)
}

//sortArr 排序
func sortArr() {
	InitArr()
	sort.Ints(arr)
	//var i, j int
	//var tmpArr []int
	//for i = 0; i < len(arr); i++ {
	//	for j = i; j < len(arr); j++ {
	//
	//	}
	//}
}

func createPipe() {
	read, write := io.Pipe()

	n, err := write.Write([]byte("xxxxx"))
	if err != nil {
		log.Printf("write.Write error:%s\n", err)
		return
	}
	fmt.Printf("write.Write %d bytes [file-based pipe]\n", n)

	//read.ReadFrom(write)
	rn, err := read.Read([]byte("wtf\r"))
	if err != nil {
		log.Printf("read.Read error:%s\n", err)
		return
	}
	fmt.Printf("read.Read %d bytes [file-based pipe]\n", rn)
}

func nginxPipe() {
	cmd1 := exec.Command("ps", "aux")
	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1
	if err := cmd1.Start(); err != nil {
		log.Printf("cmd1 start error:%s\n", err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		log.Printf("cmd1 Wait error:%s\n", err)
		return
	}

	cmd2 := exec.Command("grep", "nginx")
	cmd2.Stdin = &outputBuf1
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		log.Printf("cmd2 start error:%s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		log.Printf("cmd2 Wait error:%s\n", err)
		return
	}
	fmt.Printf("%s\n", outputBuf2.String())
}

//testBtree 测试二叉树
func testBtree() {
	var a, x, y, z tree.Btree
	//var e tree.T
	a.CreateBT()
	x.CreateBT()
	y.CreateBT()
	z.CreateBT()
	y.MakeBT('E', &a, &a)
	z.MakeBT('F', &a, &a)
	x.MakeBT('C', &y, &z)
	y.MakeBT('D', &a, &a)
	z.MakeBT('B', &y, &x)

	fmt.Println("先序遍历")
	//z.PreOrder()
	z.IPreOrder()
	fmt.Println("中序遍历")
	z.InOrder()
	//z.IInOrder()
	//z.BreakBT(&e, &y, &x)
	fmt.Println("中序遍历二叉线索树")
	z.BuildThreadBT()
	z.TInOrder()
	fmt.Println("后序遍历")
	z.PostOrder()

}

//testHeap 测试最小堆
func testHeap() {
	var heapQueue *tree.PQueue = new(tree.PQueue)

	heapQueue.CreatePQ(10)
	heapQueue.Append(71)
	heapQueue.Append(74)
	heapQueue.Append(2)
	heapQueue.Append(72)
	heapQueue.Append(54)
	heapQueue.Append(93)
	heapQueue.Append(52)
	heapQueue.Append(28)
	fmt.Println(heapQueue)
	fmt.Println("1:", heapQueue.Serve())
	fmt.Println("2:", heapQueue.Serve())
}

//testHFMCode 测试哈夫曼编码
func testHFMCode() {
	var w []tree.T = make([]tree.T, 0, 6)
	w = append(w, 9, 11, 13, 3, 5, 12)
	var ht tree.Btree = tree.CreateHFMTree(w, len(w))
	fmt.Println(ht)
}

//testUFset 测试并查集和等价关系
func testUFset() {
	var ufset *tree.UFset = new(tree.UFset)
	ufset.CreateUFset(tree.MaxSize)
	ufset.Union(0, -1)
	ufset.Union(1, 3)
	ufset.Union(2, 4)
	ufset.Union(3, -1)
	ufset.Union(4, 6)
	ufset.Union(5, 7)
	ufset.Union(6, 8)
	var j int = ufset.Find(2)
	var i int = ufset.Find(2)
	fmt.Println(j, i, ufset)
}

//testBFSearch 测试斐波那契搜索
func testBFSearch() {
	var lst *listarr.List = new(listarr.List)
	lst.CreateList(10)
	lst.Push(listarr.T{Key: 1, Data: 1111})
	lst.Push(listarr.T{Key: 2, Data: 2222})
	lst.Push(listarr.T{Key: 3, Data: 3333})
	lst.Push(listarr.T{Key: 4, Data: 4444})
	lst.Push(listarr.T{Key: 5, Data: 5555})
	lst.Push(listarr.T{Key: 6, Data: 6666})
	lst.Push(listarr.T{Key: 7, Data: 7777})
	fmt.Println(lst)
	var x *listarr.T = new(listarr.T)
	var res bool = tree.BFSearch(*lst, 7, x)
	if res == false {
		log.Printf("斐波那契搜索返回false")
		return
	}
	fmt.Println("斐波那契搜索函数返回：", x)

}
