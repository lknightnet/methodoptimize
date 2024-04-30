package main

//LAB 5

import (
	"fmt"
	"math"
)

const INF = math.MaxInt64

// Граф представлен в виде матрицы расстояний
var graph = [][]int{
	{0, 6, 7, 8, 8, 5},
	{6, 0, 5, 6, 4, 8},
	{7, 5, 0, 6, 9, 4},
	{8, 6, 6, 0, 6, 5},
	{8, 4, 9, 6, 0, 6},
	{5, 8, 4, 5, 6, 0},
}

func main() {
	n := len(graph)
	start := 0

	fmt.Println("Самая близкая вставка:")
	tour := nearestInsertion(graph, start, n)
	length := pathLength(tour)
	for i, _ := range tour {
		tour[i]++
	}
	fmt.Println("Минимальный путь:", tour, "Длина:", length)
	//fmt.Printf("Оценка: %d\n\n", evaluateSolution(tour, graph))

	fmt.Println("Ближайший сосед:")
	tour = nearestNeighbor(graph, start, n)
	length = pathLength(tour)
	for i, _ := range tour {
		tour[i]++
	}
	fmt.Println("Минимальный путь:", tour, "Длина:", length)
	//fmt.Printf("Оценка: %d\n\n", evaluateSolution(tour, graph))

	fmt.Println("Двойной обход минимального остовного дерева:")
	tour, _ = doubleMinimumSpanningTree(graph)
	length = pathLength(tour)
	for i, _ := range tour {
		tour[i]++
	}
	fmt.Println("Минимальный путь:", tour, "Длина:", length)
	//fmt.Printf("Оценка: %d\n", evaluateSolution(tour, graph))

	//fmt.Println(calculateTotalWeight(graph, treeEdges))
	//fmt.Println(treeEdges)

}

func calculateTotalWeight(distances [][]int, edges [][2]int) int {
	totalWeight := 0
	for _, edge := range edges {
		totalWeight += distances[edge[0]][edge[1]]
	}
	return totalWeight
}

func nearestInsertion(graph [][]int, start, n int) []int {
	visited := make([]bool, n)
	tour := make([]int, 0, n)
	tour = append(tour, start)
	visited[start] = true

	for len(tour) < n {
		minDist := math.MaxInt64
		insertIndex := -1

		for _, city := range tour {
			for i := 0; i < n; i++ {
				if !visited[i] && graph[city][i] < minDist {
					minDist = graph[city][i]
					insertIndex = i
				}
			}
		}

		minDist = math.MaxInt64
		minCityIndex := -1

		for i, city := range tour {
			for j := 0; j < len(tour)-1; j++ {
				dist := graph[city][tour[j]] + graph[city][tour[j+1]] - graph[tour[j]][tour[j+1]]
				if dist < minDist {
					minDist = dist
					minCityIndex = i
				}
			}
		}

		tour = append(tour[:minCityIndex+1], append([]int{insertIndex}, tour[minCityIndex+1:]...)...)
		visited[insertIndex] = true
	}

	return tour
}

func nearestNeighbor(graph [][]int, start, n int) []int {
	visited := make([]bool, n)
	tour := make([]int, 0, n)
	tour = append(tour, start)
	visited[start] = true

	for len(tour) < n {
		minDist := math.MaxInt64
		next := -1

		for i := 0; i < n; i++ {
			if !visited[i] && graph[tour[len(tour)-1]][i] < minDist {
				minDist = graph[tour[len(tour)-1]][i]
				next = i
			}
		}

		tour = append(tour, next)
		visited[next] = true
	}

	return tour
}

func minKey(key []int, mstSet []bool) int {
	min := INF
	minIndex := -1

	for i, k := range key {
		if !mstSet[i] && k < min {
			min = k
			minIndex = i
		}
	}

	return minIndex
}

func doubleMinimumSpanningTree(distances [][]int) ([]int, [][2]int) {
	n := len(distances)
	unvisited := make(map[int]bool)
	for i := 1; i < n; i++ {
		unvisited[i] = true
	}
	currentCity := 0
	treeEdges := make([][2]int, 0)

	for len(unvisited) > 0 {
		var nearestCity int
		minDistance := -1
		for city := range unvisited {
			if minDistance == -1 || distances[currentCity][city] < minDistance {
				minDistance = distances[currentCity][city]
				nearestCity = city
			}
		}
		treeEdges = append(treeEdges, [2]int{currentCity, nearestCity})
		delete(unvisited, nearestCity)
		currentCity = nearestCity
	}

	circuit := make([]int, 0)

	var dfs func(int)
	dfs = func(node int) {
		for _, edge := range treeEdges {
			if edge[0] == node || edge[1] == node {
				nextNode := edge[1]
				if edge[1] == node {
					nextNode = edge[0]
				}
				if !contains(circuit, nextNode) {
					circuit = append(circuit, nextNode)
					dfs(nextNode)
				}
			}
		}
	}

	circuit = append(circuit, 0)
	dfs(0)

	return circuit, treeEdges
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func pathLength(path []int) int {
	length := 0
	for i := 1; i < len(path); i++ {
		length += graph[path[i-1]][path[i]]
	}
	length += graph[path[len(path)-1]][path[0]] // возврат к начальной точке
	return length
}

func evaluateSolution(solution []int, distances [][]int) int {
	totalDistance := 0
	for i := 0; i < len(solution)-1; i++ {
		totalDistance += distances[solution[i]][solution[i+1]]
	}
	return totalDistance
}
