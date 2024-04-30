package main

import (
	"fmt"
	"lab1/pkg"
)

func main() {
	ff := pkg.FordFulkerson{Graph: make(map[string]map[string]*pkg.Edge)}
	ff.NumVertices = 10

	edges := []pkg.Edge{
		{Weight: 15, Start: "V0", End: "V1"},
		{Weight: 14, Start: "V1", End: "V2"},
		{Weight: 25, Start: "V2", End: "V3"},
		{Weight: 18, Start: "V3", End: "Vn"},

		{Weight: 9, Start: "V0", End: "V4"},
		{Weight: 6, Start: "V4", End: "V1"},
		{Weight: 2, Start: "V4", End: "V2"},
		{Weight: 5, Start: "V4", End: "V5"},

		{Weight: 5, Start: "V5", End: "V2"},
		{Weight: 2, Start: "V5", End: "V3"},
		{Weight: 3, Start: "V5", End: "V6"},
		{Weight: 2, Start: "V6", End: "Vn"},

		{Weight: 5, Start: "V0", End: "V7"},
		{Weight: 4, Start: "V7", End: "V8"},
		{Weight: 6, Start: "V4", End: "V8"},
		{Weight: 8, Start: "V8", End: "V5"},

		{Weight: 8, Start: "V5", End: "V9"},
		{Weight: 5, Start: "V6", End: "V9"},
		{Weight: 10, Start: "V9", End: "Vn"},
	}

	ok, description := pkg.CheckConditions(edges)
	if !ok {
		fmt.Println(description)
	}
	if ok {

		for _, edge := range edges {
			ff.AddEdge(edge.Start, edge.End, edge.Weight)
		}

		source := "V0"
		sink := "Vn"

		maxFlow, paths, pathFlows := ff.FordFulkerson(source, sink)
		fmt.Println("Максимально возможный поток:", maxFlow)
		fmt.Println("Пути:")
		var flow int
		for i, path := range paths {
			//fmt.Println(i == len(paths)-1)
			flow = flow + pathFlows[i]
			if i > 0 && i != len(paths)-1 {
				fmt.Printf("Путь: %v, Минимальный поток: %d, Общий поток: %d\n", path, pathFlows[i], flow)
			} else if i == len(paths)-1 { //больший поток
				fmt.Printf("Max: Путь: %v, Минимальный поток: %d, Общий поток: %d\n", path, pathFlows[i], flow)
			} else if i != len(paths)-1 {
				fmt.Printf("Путь: %v, Минимальный поток: %d\n", path, pathFlows[i])
			}
		}
		fmt.Println(len(paths))
	}

}
