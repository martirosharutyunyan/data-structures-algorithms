package main

import (
	"DSA/src/data-structures/graph"
	"fmt"
)

func main() {
	// g := graph.Graph{Directed: true}
	// g.Vertexes = map[int][]int{}
	// g.AddVertex(0, []int{1, 2, 3})
	// g.AddVertex(1, []int{2, 4})
	// g.AddVertex(3, []int{2})
	// g.AddVertex(4, []int{3, 5})
	// fmt.Println(g.CountPaths(0, 2))
	// g2 := graph.Graph{Directed: true}
	// g2.Vertexes = map[int][]int{}
	// g2.AddVertex(0, []int{1, 2})
	// g2.AddVertex(1, []int{})
	// fmt.Println(g2.HaveCycle(0))

	adj := [][]int{
		{0, 5, graph.Inf, 10},
		{graph.Inf, 0, 3, graph.Inf},
		{graph.Inf, graph.Inf, 0, 1},
		{graph.Inf, graph.Inf, graph.Inf, 0},
	}
	fmt.Println(graph.FloydMarshal(&adj))
}
