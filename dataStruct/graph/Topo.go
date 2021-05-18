package graph

import "fmt"

//TopoSort 拓扑排序
func TopoSort(g AdjacencyTableGraph, order []int) {
	var i, j, k int
	var (
		count = -1
		top   = -1
		n     = g.Vertices
	)
	var p *ENode
	for i = 0; i < n; i++ {
		if InDegree[i] <= 0 {
			InDegree[i] = top
			top = i
		}
	}
	for i = 0; i < n; i++ {
		if top == -1 {
			fmt.Print("Network has a cycle. TopoSort terminated")
			return
		} else {
			j = top
			top = InDegree[top]
			count++
			order[count] = j
			for p = g.A[j]; p != nil; p = p.NextArc {
				k = p.AdjVex
				InDegree[k]--
				if InDegree[k] <= 0 {
					InDegree[k] = top
					top = k
				}
			}
		}
	}
}

// Earliest 关键路径，按照拓扑次序计算earliest
func Earliest(g AdjacencyTableGraph, order []int, earliest []int) {
	var i, k int
	var p *ENode
	var n int = g.Vertices
	earliest = make([]int, n)
	for i = 0; i < n; i++ {
		k = order[i]
		for p = g.A[k]; p != nil; p = p.NextArc {
			if earliest[p.AdjVex] < earliest[k]+int(p.W) {
				earliest[p.AdjVex] = earliest[k] + int(p.W)
			}
		}
	}
}

//Latest 关键路径，按照逆拓扑次序计算latest值
func Latest(g AdjacencyTableGraph, order []int, earliest []int, latest []int) {
	var i, j, k int
	var n int = g.Vertices
	var p *ENode
	for i = 0; i < n; i++ {
		latest[i] = earliest[n-1]
	}
	for i = n - 2; i > -1; i-- {
		j = order[i]
		for p = g.A[j]; p != nil; p = p.NextArc {
			k = p.AdjVex
			if latest[j] > latest[k]-int(p.W) {
				latest[j] = latest[k] - int(p.W)
			}
		}
	}
}
