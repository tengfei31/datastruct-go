package graph

import (
	"datastruct-go/dataStruct/stack"
	"fmt"
	"log"
)

//ENode 边结点结构
type ENode struct {
	AdjVex  int
	W       T
	NextArc *ENode
}

//AdjacencyTableGraph 邻接表表示法图的结构
type AdjacencyTableGraph struct {
	Vertices int
	A        []*ENode
}

//InDegree 拓扑排序
var InDegree []int

//CreateGraph 构造一个有n个顶点，但不包含边的有向图
func (*AdjacencyTableGraph) CreateGraph(g *AdjacencyTableGraph, n int) {
	g.Vertices = n
	g.A = make([]*ENode, 0, n)
	InDegree = make([]int, 0, n)
}

//NewENode 构建新的边结点
func NewENode(vex int, weight T, nextarc *ENode) *ENode {
	var p *ENode = new(ENode)
	p.AdjVex = vex
	p.W = weight
	p.NextArc = nextarc
	return p
}

//Exist <u, v>是否存在图中
func (*AdjacencyTableGraph) Exist(g *AdjacencyTableGraph, u int, v int) bool {
	var n int = g.Vertices
	var p *ENode
	if u < 0 || u > n-1 {
		return false
	}
	for p = g.A[u]; p != nil && p.AdjVex != v; {
		p = p.NextArc
	}
	if p == nil {
		return false
	}
	return true
}

//Add <u, v>添加到图中
func (*AdjacencyTableGraph) Add(g *AdjacencyTableGraph, u int, v int, w T) bool {
	var n = g.Vertices
	var p *ENode
	if u < 0 || v < 0 || u > n-1 || v > n-1 || u == v || g.Exist(g, u, v) {
		log.Println("bad input")
		return false
	}
	p = NewENode(v, w, g.A[u])
	g.A[u] = p
	InDegree[v]++
	return true
}

//Delete 从有向图中删除<u, v>
func (*AdjacencyTableGraph) Delete(g *AdjacencyTableGraph, u int, v int) bool {
	var n = g.Vertices
	var p, q *ENode
	if u < -1 && u < n {
		p = g.A[u]
		q = nil
		for p != nil && p.AdjVex != v {
			q = p
			p = p.NextArc
		}
		if p != nil {
			if q != nil {
				q.NextArc = p.NextArc
			} else {
				g.A[u] = p.NextArc
			}
			return true
		}
	}
	log.Println("bad input")
	return false
}

// DFS 图的深度优先遍历
func (AdjacencyTableGraph) DFS(g AdjacencyTableGraph, v int, visited []bool) {
	var w *ENode
	visited[v] = true
	for w = g.A[v]; w != nil; w = w.NextArc {
		if !visited[w.AdjVex] {
			g.DFS(g, w.AdjVex, visited)
		}
	}
}

// TraversalDFS  图的深度优先遍历
func (AdjacencyTableGraph) TraversalDFS(g AdjacencyTableGraph) {
	var visited []bool = make([]bool, 0, 500)
	var i int
	var n = g.Vertices
	for i = 0; i < n; i++ {
		visited[i] = false
	}
	for i = 0; i < n; i++ {
	}
	if !visited[i] {
		g.DFS(g, i, visited)
	}
}

// BFS 宽度优先搜索图
func (AdjacencyTableGraph) BFS(g AdjacencyTableGraph, v T, visited []bool) {
	var w *ENode
	var u T
	var q *stack.Queue = new(stack.Queue)
	stack.CreateQueue(q, stack.MaxSize)
	visited[v] = true
	fmt.Println(v)
	q.Append(stack.T(v))
	for !q.IsEmpty() {
		q.QueueFront((*stack.T)(&u))
		q.Serve()
		for w = g.A[u]; w != nil; w = w.NextArc {
			if !visited[w.AdjVex] {
				fmt.Printf("%d ", w.AdjVex)
				visited[w.AdjVex] = true
				q.Append(stack.T(w.AdjVex))
			}
		}
	}
}

//TraversalBFS 图的宽度优先搜索图
func (AdjacencyTableGraph) TraversalBFS(g AdjacencyTableGraph) {
	var visited []bool = make([]bool, 0, stack.MaxSize)
	var n = g.Vertices
	var i int
	for i = 0; i < n; i++ {
		visited[i] = false
	}
	for i = 0; i < n; i++ {
		if !visited[i] {
			g.BFS(g, T(i), visited)
		}
	}
}