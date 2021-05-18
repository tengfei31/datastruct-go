package main

import (
	"bytes"
	"datastruct-go/dataStruct/hashtable"
	"datastruct-go/dataStruct/listarr"
	"datastruct-go/dataStruct/tree"
	"fmt"
	"io"
	"log"
	"os/exec"
)

type s1 struct {
	i1 int8
	i2 int16
	i3 int32
}

type s2 struct {
	i1 int8
	i2 *int16
	i3 int32
	//_ [3]int8
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	deferCall()
	//[0,1,2,3,4]
	//[0,1,2,2,1]
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	target := createTargetArray(nums, index)
	fmt.Println(target)
}

func deferCall() {
	fmt.Println("1")
	defer func() { fmt.Println("A") }()
	fmt.Println("2")
	defer func() { fmt.Println("B") }()
	fmt.Println("3")
	panic("触发异常")
	defer func() { fmt.Println("C") }()
	fmt.Println("4")
}

func createTargetArray(nums []int, index []int) []int {
	var target = make([]int, len(index))
	for i := 0; i < len(target); i++ {
		target[i] = -1
	}
	for i := 0; i < len(index); i++ {
		key := index[i]
		val := nums[i]
		//处理key的位置存在值
		if target[key] > -1 {
			copy(target[key+1:], target[key:])
		}
		target[key] = val
	}
	return target
}

func handleArr(arr []int) {
	arr[1] = 11
	//arr = append(arr, 1)
	//arr = append(arr, 2)
	//arr = append(arr, 3)
	//arr = append(arr, 4)
	//arr = append(arr, 5)
	//var i int
	//for i = 0; i < len(arr); i++ {
	//	arr[i] = arr[i] + i
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

//testBtSearch 测试二叉搜索树
func testBtSearch() {
	var bt *tree.Btree = new(tree.Btree)
	tree.BtInsert(bt, 28)
	tree.BtInsert(bt, 21)
	tree.BtInsert(bt, 25)
	tree.BtInsert(bt, 36)
	tree.BtInsert(bt, 33)
	tree.BtInsert(bt, 43)
	//中序遍历bt
	bt.InOrder()

	//二叉搜索树删除节点
	var x *tree.T = new(tree.T)
	var res bool = tree.BtRemove(bt, 28, x)
	if !res {
		log.Printf("btremove is false")
		return
	}
	//中序遍历bt
	bt.InOrder()
	fmt.Printf("%+v\n", bt.Root)
	fmt.Println(*x)
}

//testAVLBTree 测试二叉平衡树
func testAVLBTree() {
	var bt *tree.AVLBTree = new(tree.AVLBTree)
	tree.AVLInsert(bt, 45)
	tree.AVLInsert(bt, 28)
	tree.AVLInsert(bt, 15)
	tree.AVLInsert(bt, 12)
	tree.AVLInsert(bt, 14)
	tree.AVLInsert(bt, 23)
	fmt.Printf("%+v\n", bt.Root)
}

//testSkipTable 测试跳表
func testSkipTable() {
	var skip = hashtable.NewSkipList(100000, 10)
	var insertArr = []hashtable.T{
		{Key: 3, Data: 1},
		{Key: 7, Data: 20},
		{Key: 19, Data: 30},
		{Key: 22, Data: 40},
		{Key: 43, Data: 40},
		{Key: 48, Data: 40},
		{Key: 70, Data: 40},
	}
	for _, v := range insertArr {
		insert := skip.Insert(v)
		if !insert {
			fmt.Println("insert失败:", v)
		}
	}
	fmt.Println("skip的层级:", skip.Level)
	var x *hashtable.T = new(hashtable.T)
	var res = skip.Search(70, x)
	if !res {
		log.Fatalf("没有找到k")
	}

	fmt.Println(x, skip.Level)
}

//testHashTable 测试散列表
func testHashTable() {
	var hashTB *hashtable.HashTable = new(hashtable.HashTable)
	var divitor = 13
	//初始化
	hashtable.CreateHashTable(hashTB, divitor)
	var insertArr = []hashtable.T{
		{Key: 1, Data: 1},
		{Key: 2, Data: 2},
		{Key: 3, Data: 3},
		{Key: 4, Data: 4},
		{Key: 5, Data: 5},
		{Key: 15, Data: 15},
	}
	for _, v := range insertArr {
		var insert = hashTB.Insert(v)
		if !insert {
			log.Println("")
			continue
		}
	}
	var x = new(hashtable.T)
	var searchKey hashtable.KeyType = 15
	var search = hashTB.Search(searchKey, x)
	if !search {
		log.Fatalf("没有找到:%d", searchKey)
	}

	fmt.Println("散列表:", hashTB)
	fmt.Println("散列表搜索结果:", x)

	var del = hashTB.Delete(searchKey, x)
	if !del {
		log.Fatalf("删除失败:%d", searchKey)
	}
	fmt.Println("散列表:", hashTB)
	fmt.Println("散列表删除结果:", x)

}
