package main

import (
	"bytes"
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

// testHFMCode 测试哈夫曼编码
func testHFMCode() {
	var w = make([]tree.BnElement[int], 0, 6)
	w = append(w, tree.BnElement[int]{W: 9}, tree.BnElement[int]{W: 11}, tree.BnElement[int]{W: 13}, tree.BnElement[int]{W: 3}, tree.BnElement[int]{W: 5}, tree.BnElement[int]{W: 12})
	var ht = tree.CreateHFMTree[int](w, len(w))
	fmt.Println(ht)
}

// testUFset 测试并查集和等价关系
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

// testBFSearch 测试斐波那契搜索
func testBFSearch() {
	var lst = new(listarr.List[int])
	lst.CreateList(10)
	lst.Push(listarr.ElementT[int]{Key: 1, Data: 1111})
	lst.Push(listarr.ElementT[int]{Key: 2, Data: 2222})
	lst.Push(listarr.ElementT[int]{Key: 3, Data: 3333})
	lst.Push(listarr.ElementT[int]{Key: 4, Data: 4444})
	lst.Push(listarr.ElementT[int]{Key: 5, Data: 5555})
	lst.Push(listarr.ElementT[int]{Key: 6, Data: 6666})
	lst.Push(listarr.ElementT[int]{Key: 7, Data: 7777})
	fmt.Println(lst)
	var x = new(listarr.ElementT[int])
	var res bool = tree.BFSearch(*lst, 7, x)
	if res == false {
		log.Printf("斐波那契搜索返回false")
		return
	}
	fmt.Println("斐波那契搜索函数返回：", x)
}
