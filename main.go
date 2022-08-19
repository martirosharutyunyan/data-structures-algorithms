package main

import (
	"DSA/src/data-structures/graph"
)

func main() {
	g := graph.Graph{}
	g.Vertexes = map[int][]int{}
	g.AddVertex(0, 1)
	g.AddVertex(0, 2)
	g.AddVertex(1, 2)
	g.BFS(0)
	// adj := [][]int{}
	// graph.AddEdge(&adj, 0, 1)
	// graph.AddEdge(&adj, 2, 3)
	// graph.AddEdge(&adj, 3, 2)
	// graph.AddEdge(&adj, 3, 0)
	// fmt.Println(graph.IsBipartite(adj))
}
