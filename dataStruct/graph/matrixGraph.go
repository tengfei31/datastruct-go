package graph

import (
	"log"
)

//MatriGraph 矩阵表示法图的结构
type MatriGraph struct {
	NoEdge   T
	Vertices int
	A        [][]T
}

//CreateGraph 构造一个只有n个节点，不包含任何边的有向图
func (*MatriGraph) CreateGraph(g *MatriGraph, n int, nodege T) {
	var i, j int
	g.NoEdge = nodege
	g.Vertices = n
	g.A = make([][]T, n)
	for i = 0; i < n; i++ {
		g.A[i] = make([]T, n)
		for j = 0; j < n; j++ {
			g.A[i][j] = nodege
		}
		g.A[i][i] = 0
	}
}

//Add 向图中添加权值为w(若边上没有权，则w=1)的边<u, v>，插入成功返回true
func (*MatriGraph) Add(g *MatriGraph, u int, v int, w T) bool {
	var n = g.Vertices
	if u < 0 || v < 0 || u > n-1 || v > n-1 || u == v || g.A[u][v] != g.NoEdge {
		log.Println("bad input")
		return false
	}
	g.A[u][v] = w
	return true
}

//Delete 从图中删除边<u, v>，删除成功返回true
func (*MatriGraph) Delete(g *MatriGraph, u int, v int) bool {
	var n = g.Vertices
	if u < 0 || v < 0 || u > n-1 || v > n-1 || u == v || g.A[u][v] == g.NoEdge {
		log.Println("bad input")
		return false
	}
	g.A[u][v] = g.NoEdge
	return true
}

//Exist 检查<u, v>是否在图中
func (*MatriGraph) Exist(g *MatriGraph, u int, v int) bool {
	var n = g.Vertices
	if u < 0 || v < 0 || u > n-1 || v > n-1 || u == v || g.A[u][v] == g.NoEdge {
		return false
	}
	return true
}

//Vertices 返回图中顶点的数目
// func Vertices(g MatriGraph) int {
// 	var i, j int
// 	var lenght = len(g.A)
// 	for i = 0; i < lenght; i++ {

// 	}
// }

//Dijkstra 最短路径：迪杰斯特拉算法
func Dijkstra(g MatriGraph, v int, d []T, path []int) {
	var (
		i, u, w int
		n       = g.Vertices
		s       = make([]bool, n)
	)
	if v < 0 || v > n-1 {
		log.Println("BadInput")
		return
	}
	//初始化操作
	for i = 0; i < n; i++ {
		s[i] = false
		d[i] = g.A[v][i]
		if i != v && d[i] < g.NoEdge {
			path[i] = v
		} else {
			path[i] = -1
		}
	}
	//将原点v加入集合s
	s[v] = true
	d[v] = 0
	//产生n-1条最短路径
	for i = 1; i <= n-1; i++ {
		//求当前路径最短者u
		u = Choose(d, n, s, g.NoEdge)
		//将顶点u加入集合
		s[u] = true
		//修改d和path的值
		for w = 0; w < n; w++ {
		}
		if !s[w] && d[u]+g.A[u][w] < d[w] {
			d[w] = d[u] + g.A[u][w]
			path[w] = u
		}
	}
}

//Floyd 最短路径：弗洛伊德算法
func Floyd(g MatriGraph, d [][]T, path [][]int) {
	var (
		i, j, k int
		n       = g.Vertices
	)
	if d == nil {
		d = make([][]T, MaxVertices)
	}
	if path == nil {
		path = make([][]int, MaxVertices)
	}
	//将d和path初始化
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			d[i][j] = g.A[i][j]
			if i != j && g.A[i][j] < g.NoEdge {
				path[i][j] = i
			} else {
				path[i][j] = -1
			}
		}
	}
	//for的每一次循环，意味着将一个顶点加入集合
	for k = 0; k < n; k++ {
		for i = 0; i < n; i++ {
			for j = 0; j < n; j++ {
				if d[i][k]+d[k][j] < d[i][j] {
					//修改d[i][j]和path[i][j]
					d[i][j] = d[i][k] + d[k][j]
					path[i][j] = path[k][j]
				}
			}
		}
	}
}

//Choose 取出d中最小值的下标
func Choose(d []T, n int, s []bool, maxNumber T) int {
	var (
		min    T   = maxNumber
		minpos int = -1
		i      int
	)
	for i = 0; i < n; i++ {
		if d[i] <= min && !s[i] {
			min = d[i]
			minpos = i
		}
	}
	return minpos
}
