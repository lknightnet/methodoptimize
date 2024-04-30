package main

import (
	"fmt"
	"math"
)

const infinity = math.MaxInt32

type Graph struct {
	countEdges int
	graph      [][]int
}

func MakeGraph(n int) *Graph {
	return &Graph{
		countEdges: n,
		graph:      make([][]int, 0),
	}
}

func (g *Graph) AddEdge(side []int) {
	g.graph = append(g.graph, side)
}

type Path struct {
	Nodes []int
}

func (g *Graph) floydWarshall() ([][]int, [][]Path) {
	dist := make([][]int, g.countEdges)
	next := make([][]int, g.countEdges)
	paths := make([][]Path, g.countEdges)

	for i := range dist {
		dist[i] = make([]int, g.countEdges)
		next[i] = make([]int, g.countEdges)
		paths[i] = make([]Path, g.countEdges)
		for j := range dist[i] {
			dist[i][j] = g.graph[i][j]
			if g.graph[i][j] != infinity {
				next[i][j] = j
				paths[i][j].Nodes = []int{i, j}
			} else {
				next[i][j] = -1
			}
		}
	}

	for k := 0; k < g.countEdges; k++ {
		for i := 0; i < g.countEdges; i++ {
			for j := 0; j < g.countEdges; j++ {
				if dist[i][k] != infinity && dist[k][j] != infinity && dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					next[i][j] = next[i][k]
					paths[i][j].Nodes = make([]int, len(paths[i][k].Nodes)+len(paths[k][j].Nodes))
					copy(paths[i][j].Nodes, paths[i][k].Nodes)
					copy(paths[i][j].Nodes[len(paths[i][k].Nodes):], paths[k][j].Nodes)
				}
			}
		}
	}
	return dist, paths
}

func (g *Graph) printShortestPaths(paths [][]Path, dist [][]int) {
	fmt.Println("Короткие пути:")
	for i := 0; i < g.countEdges; i++ {
		for j := 0; j < g.countEdges; j++ {
			if i != j {
				var nodes = make([]int, 0)
				for _, v := range paths[i][j].Nodes {
					q := v + 1
					nodes = append(nodes, q)
				}
				fmt.Printf("От %d до %d. Дистанция: %d, Путь: %v\n", i+1, j+1, dist[i][j], nodes)
				nodes = nil
			}
		}
	}
}

func main() {
	graph := [][]int{
		{0, 8, infinity, infinity, infinity, 2, infinity, infinity, infinity},
		{8, 0, 12, 10, 6, 12, infinity, infinity, infinity},
		{infinity, 12, 0, 6, 4, infinity, infinity, infinity, infinity, 3},
		{infinity, 10, 6, 0, 14, infinity, infinity, 10, infinity},
		{infinity, 6, infinity, 14, 0, 6, 16, 2, infinity},
		{2, 12, infinity, infinity, 6, 0, 4, infinity, 7},
		{infinity, infinity, infinity, infinity, 16, 4, 0, 29, 9},
		{infinity, infinity, infinity, 10, 2, infinity, 20, 0, infinity},
		{infinity, infinity, 3, infinity, infinity, 7, 9, infinity, 0},
	}

	g := MakeGraph(9)

	for _, side := range graph {
		g.AddEdge(side)
	}

	dist, paths := g.floydWarshall()

	g.printShortestPaths(paths, dist)

	for _, v := range dist {
		fmt.Println(v)
	}
}
