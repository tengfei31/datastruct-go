package tree

//并查集与等价关系

type UFset struct {
	Parent []int
	Size   int
}

//CreateUFset 创建每个子集中仅包含一个元素的并查集s,长度为n
func (ufset *UFset) CreateUFset(n int) {
	var i int
	ufset.Size = n
	ufset.Parent = make([]int, ufset.Size)
	for i = 0; i < ufset.Size; i++ {
		ufset.Parent[i] = -1
	}
}

//Find 返回i所在的子集合的标识
func (ufset *UFset) Find(i int) int {
	for ; ufset.Parent[i] >= 0; i = ufset.Parent[i] {
	}
	return i
}

//Find2 改进find方法
func (ufset *UFset) Find2(i int) int {
	var r, t, l int
	for ; ufset.Parent[r] >= 0; r = ufset.Parent[r] {
	}
	if i != r {
		for t = i; ufset.Parent[t] != r; t = l {
			l = ufset.Parent[t]
			ufset.Parent[t] = r
		}
	}
	return r
}

//Union 合并x,y两个子集
func (ufset *UFset) Union(x int, y int) {
	ufset.Parent[x] = y
}

//Union2 改进的union方法
func (ufset *UFset) Union2(x int, y int) {
	var temp = ufset.Parent[x] + ufset.Parent[y]
	if ufset.Parent[x] > ufset.Parent[y] {
		ufset.Parent[x] = y
		ufset.Parent[y] = temp
	} else {
		ufset.Parent[y] = x
		ufset.Parent[x] = temp
	}
}
