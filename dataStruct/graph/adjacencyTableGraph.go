package graph

import (
	"datastruct-go/dataStruct/stack"
	"datastruct-go/dataStruct/tree"
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
func (g *AdjacencyTableGraph) CreateGraph(n int) {
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
func (g *AdjacencyTableGraph) Exist(u int, v int) bool {
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
func (g *AdjacencyTableGraph) Add(u int, v int, w T) bool {
	var n = g.Vertices
	var p *ENode
	if u < 0 || v < 0 || u > n-1 || v > n-1 || u == v || g.Exist(u, v) {
		log.Println("bad input")
		return false
	}
	p = NewENode(v, w, g.A[u])
	g.A[u] = p
	InDegree[v]++
	return true
}

//Delete 从有向图中删除<u, v>
func (g *AdjacencyTableGraph) Delete(u int, v int) bool {
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
func (g AdjacencyTableGraph) DFS(v int, visited []bool) {
	var w *ENode
	visited[v] = true
	for w = g.A[v]; w != nil; w = w.NextArc {
		if !visited[w.AdjVex] {
			g.DFS(w.AdjVex, visited)
		}
	}
}

// TraversalDFS  图的深度优先遍历
func (g AdjacencyTableGraph) TraversalDFS() {
	var visited []bool = make([]bool, 0, 500)
	var i int
	var n = g.Vertices
	for i = 0; i < n; i++ {
		visited[i] = false
	}
	for i = 0; i < n; i++ {
	}
	if !visited[i] {
		g.DFS(i, visited)
	}
}

// BFS 宽度优先搜索图
func (g AdjacencyTableGraph) BFS(v T, visited []bool) {
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
func (g AdjacencyTableGraph) TraversalBFS() {
	var visited []bool = make([]bool, 0, stack.MaxSize)
	var n = g.Vertices
	var i int
	for i = 0; i < n; i++ {
		visited[i] = false
	}
	for i = 0; i < n; i++ {
		if !visited[i] {
			g.BFS(T(i), visited)
		}
	}
}

// Prim 最小代价生成树：普里姆算法
func (g AdjacencyTableGraph) Prim(k int, nearest []int, lowcost []T) {
	var i, j, n int
	n = g.Vertices
	var (
		min  T
		mark = make([]bool, MaxVertices)
		p    *ENode
	)
	if k < 0 || k > n-1 {
		fmt.Println("BadInput")
		return
	}
	// 初始化
	for i = 0; i < n; i++ {
		nearest[i] = -1
		mark[i] = false
		lowcost[i] = MaxNum
	}
	// 源点k加入到生成树
	lowcost[k] = 0
	nearest[k] = k
	mark[k] = true
	for i = 1; i < n; i++ {
		// 修改lowcost和nearest的值
		for p = g.A[i]; p != nil; p = p.NextArc {
			j = p.AdjVex
			if mark[j] == false && lowcost[j] > p.W {
				lowcost[j] = p.W
				nearest[j] = k
			}
		}
		// 求下一条最小权值的边
		min = MaxNum
		for j = 0; j < n; j++ {
			if mark[j] == false && lowcost[j] < min {
				min = lowcost[j]
				k = j
			}
		}
		// 将顶点k加到生成树上
		mark[k] = true
	}
}

//EdgeNode 最小代价生成树：克鲁斯卡尔算法结构
type EdgeNode struct {
	tree.Weight
	U, V int
	W    T
}

//GetWeight 获取权重
func (node EdgeNode) GetWeight() int {
	return int(node.W)
}

//SetWeight 设置权重
func (node EdgeNode) SetWeight(w int) {
	node.W = T(w)
}

//Kruskal 最小代价生成树：克鲁斯卡尔算法
// 优先权队列pq中保存无向图边的集合，n是无向图的顶点个数
func Kruskal(pq *tree.PQueue, n int) {
	var (
		s       = new(tree.UFset)
		x       EdgeNode
		u, v, k int
	)
	k = 0
	//建立并查集
	s.CreateUFset(n)
	//求最小代价生成树的n-1条边
	for k < n-1 && pq.IsEmpty() == false {
		//从优先权队列pq中取出最小代价的边x=(u,v,w)
		x = EdgeNode(pq.Serve())
		//分别查找边的x.u和x.v所在的子集
		u = s.Find2(x.U)
		v = s.Find2(x.V)
		//如果边的两个端点不在同一个子集
		if u != v {
			//合并两个子集u和v
			s.Union2(u, v)
			k++
			//输出边x
			fmt.Printf("%d, %d, %d", x.U, x.V, x.W)
		}
	}
	// 不足n-1条边，原图不是连通图
	if k < n-2 {
		fmt.Println("The graph is not connected!")
	}
}

//最短路径：迪杰斯特拉算法

//最短路径：弗洛伊德算法
