package main

import (
	"DSA/src/data-structures/graph"
	"fmt"
)

func main() {
	g := graph.Graph{}
	g.Vertexes = map[int][]int{}
	g.AddVertex(0, []int{1, 2, 3})
	g.AddVertex(1, []int{2, 4})
	g.AddVertex(2, []int{})
	g.AddVertex(3, []int{2})
	g.AddVertex(4, []int{3})
	fmt.Println(g.PossiblePaths(0, 2))
	// adj := [][]int{}
	// graph.AddEdge(&adj, 0, 1)
	// graph.AddEdge(&adj, 2, 3)
	// graph.AddEdge(&adj, 3, 2)
	// graph.AddEdge(&adj, 3, 0)
	// fmt.Println(graph.IsBipartite(adj))
}
