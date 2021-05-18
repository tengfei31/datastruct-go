package graph

//graphInterface 基础接口
// type graphInterface interface {
// 	//CreateGraph 构造一个只有n个节点，不包含任何边的有向图
// 	CreateGraph(g *Graph, n int, nodege T)
// 	//Add 向图中添加权值为w(若边上没有权，则w=1)的边<u, v>，插入成功返回true
// 	Add(g *Graph, u int, v int, w T) bool
// 	//Delete 从图中删除边<u, v>，删除成功返回true
// 	Delete(g *Graph, u int, v int) bool
// 	//Exist 检查<u, v>是否在图中
// 	Exist(g *Graph, u int, v int) bool
// 	//Vertices 返回图中顶点的数目
// 	Vertices(g Graph) int

// 	//DFS 深度优先搜索图
// 	DFS(g Graph)
// 	//BFS 宽度优先搜索图
// 	BFS(g Graph)
// 	//TopoSort 拓扑排序
// 	TopoSort(g Graph)
// 	//CriticalPath 关键路径
// 	CriticalPath(g Graph)
// 	//Prim 普里姆算法求最小代价生成树
// 	Prim(g Graph, k int)
// 	//Kruskal 克鲁斯卡尔算法求最小代价生成树
// 	Kruskal(g Graph, edges int)
// 	//Dijkstra 迪杰斯特拉算法求单源最短路径
// 	Dijkstra(g Graph, k int, d []T, p []int)
// 	//Floyd 佛洛依德算法求所有丁点之间的最短路径
// 	Floyd(g Graph, d **T, path **int)
// }

//KeyType 基础类型
type KeyType int

//DataType 基础类型
type DataType int

//T 数据的类型
type T int

const MaxNum T = 100
const MaxVertices T = 100
