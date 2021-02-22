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
