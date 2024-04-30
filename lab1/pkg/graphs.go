package pkg

import "fmt"

const (
	inf = int(^uint(0) >> 1)
)

type Edge struct {
	Weight int
	Start  string
	End    string
}

func (e *Edge) String() string {
	return fmt.Sprintf("weight %d,start %s,end %s", e.Weight, e.Start, e.End)
}

type FordFulkerson struct {
	Graph       map[string]map[string]*Edge
	NumVertices int
}

func (ff *FordFulkerson) AddEdge(start, end string, weight int) {
	if _, ok := ff.Graph[start]; !ok {
		ff.Graph[start] = make(map[string]*Edge)
	}
	ff.Graph[start][end] = &Edge{Weight: weight, Start: start, End: end}

	if _, ok := ff.Graph[end]; !ok {
		ff.Graph[end] = make(map[string]*Edge)
	}
	ff.Graph[end][start] = &Edge{Weight: 0, Start: end, End: start}
}

func (ff *FordFulkerson) bfs(s, t string, parent map[string]string) bool {
	visited := make(map[string]bool)
	queue := []string{s}
	visited[s] = true
	parent[s] = ""

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v, e := range ff.Graph[u] {
			if !visited[v] && e.Weight > 0 {
				queue = append(queue, v)
				parent[v] = u
				visited[v] = true
			}
		}
	}

	return visited[t]
}

func (ff *FordFulkerson) FordFulkerson(s, t string) (int, [][]string, []int) {
	parent := make(map[string]string)
	maxFlow := 0
	var allPaths [][]string
	var pathFlows []int

	for ff.bfs(s, t, parent) {
		pathFlow := inf
		var path []string

		for v := t; v != s; v = parent[v] {
			u := parent[v]
			path = append([]string{v}, path...)
			pathFlow = min(pathFlow, ff.Graph[u][v].Weight)
		}
		path = append([]string{s}, path...)

		allPaths = append(allPaths, path)
		pathFlows = append(pathFlows, pathFlow)

		for v := t; v != s; v = parent[v] {
			u := parent[v]
			ff.Graph[u][v].Weight -= pathFlow
			ff.Graph[v][u].Weight += pathFlow
		}

		maxFlow += pathFlow

	}

	//for _, v := range ff.graph {
	//	for i2, v2 := range v {
	//		fmt.Println(i2, v2)
	//	}
	//}

	return maxFlow, allPaths, pathFlows
}
