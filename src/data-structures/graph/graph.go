package graph

import (
	"fmt"
)

type Graph struct {
	Vertexes map[int][]int
	Directed bool
}

const Inf = 99999

type Priority struct {
	Cost  int
	Value int
}

type WeightedGraph struct {
	Vertexes map[int][]Priority
}

func (graph *WeightedGraph) AddVertex(src int, dst []Priority) {
	graph.Vertexes[src] = append(graph.Vertexes[src], dst...)
}

func InsertMatrix(matrix *[][]int, x, y int) {
	if len(*matrix) < x+1 {
		temp := make([][]int, x+1)
		for i := 0; i < len(*matrix); i++ {
			var subMatrix []int
			for j := 0; j < len((*matrix)[i]); j++ {
				subMatrix = append(subMatrix, (*matrix)[i][j])
			}
			temp[i] = subMatrix
		}
		*matrix = temp
		(*matrix)[x] = append((*matrix)[x], y)

		return
	}

	(*matrix)[x] = append((*matrix)[x], y)
}

func AddEdge(adj *[][]int, src, dest int) {
	InsertMatrix(adj, src, dest)
}

func DisplayGraph(adj [][]int) {
	for i := 0; i < len(adj); i++ {
		text := fmt.Sprintf("%v -->", i)
		for _, v := range adj[i] {
			text += fmt.Sprintf(" %v", v)
		}
		fmt.Println(text)
	}
}

func TransposeGraph(adj, transpose *[][]int) {
	for i := 0; i < len(*adj); i++ {
		for j := 0; j < len((*adj)[i]); j++ {
			AddEdge(transpose, (*adj)[i][j], i)
		}
	}
}

func IsBipartite(adj [][]int) bool {
	col := make([]int, len(adj))
	Fill(&col, -1)

	var queue [][]int

	for i := 0; i < len(adj); i++ {
		if col[i] == -1 {
			queue = append(queue, []int{i, 0})
			col[i] = 0

			for len(queue) != 0 {
				p := queue[0]
				queue = queue[1:]
				v, c := p[0], p[1]

				for _, j := range adj[v] {
					if col[j] == c {
						return false
					}

					if col[j] == -1 {
						col[j] = 1 ^ c
						queue = append(queue, []int{j, col[j]})
					}
				}
			}
		}
	}

	return true
}

func Fill[T any](arr *[]T, value T) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = value
	}
}

func (graph *Graph) BFS(source int) {
	visited := map[int]bool{}

	queue := []int{source}
	visited[source] = true

	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]
		fmt.Println(front)

		for i := 0; i < len(graph.Vertexes[front]); i++ {
			if visited[graph.Vertexes[front][i]] {
				continue
			}
			visited[graph.Vertexes[front][i]] = true
			queue = append(queue, graph.Vertexes[front][i])
		}
	}
}

func (graph *Graph) DFS(source int) {
	visited := map[int]bool{}

	stack := []int{source}
	visited[source] = true

	for len(stack) != 0 {
		front := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(front)

		for i := 0; i < len(graph.Vertexes[front]); i++ {
			if visited[graph.Vertexes[front][i]] {
				continue
			}
			visited[graph.Vertexes[front][i]] = true
			stack = append(stack, graph.Vertexes[front][i])
		}
	}
}

// review and fix bugs
func (graph *Graph) PossiblePaths(start, end int) [][]int {
	var results [][]int
	stack := []int{start}
	var temp []int
	for len(stack) != 0 {
		front := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := graph.Vertexes[front]; ok || front == end {
			temp = append(temp, front)
			stack = append(stack, graph.Vertexes[front]...)
		}
		if temp[len(temp)-1] == end {
			results = append(results, temp)
			temp = []int{start}
		}
	}

	return results
}

func (graph *Graph) CountPaths(src, dst int) int {
	pathCount := 0

	CountPathsUtil(graph, src, dst, &pathCount)

	return pathCount
}

func CountPathsUtil(graph *Graph, src, dst int, pathCount *int) {
	if src == dst {
		*pathCount++
	} else {
		if _, ok := graph.Vertexes[src]; ok {
			for _, value := range graph.Vertexes[src] {
				CountPathsUtil(graph, value, dst, pathCount)
			}
		}
	}
}

func (graph *Graph) AddVertex(source int, dst []int) {
	graph.Vertexes[source] = append(graph.Vertexes[source], dst...)
	if !graph.Directed {
		for _, v := range dst {
			graph.Vertexes[v] = append(graph.Vertexes[v], source)
		}
	}
}

func (graph *Graph) HaveCycle(source int) bool {
	visited := map[int]bool{}
	queue := []int{source}
	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]
		fmt.Println(front)

		for i := 0; i < len(graph.Vertexes[front]); i++ {
			if visited[graph.Vertexes[front][i]] {
				return true
			}
			visited[graph.Vertexes[front][i]] = true
			queue = append(queue, graph.Vertexes[front][i])
		}
	}

	return false
}

func FloydMarshal(graph *[][]int) [][]int {
	var dist [][]int
	for i := 0; i < len(*graph); i++ {
		for j := 0; j < len((*graph)[i]); j++ {
			AddEdge(&dist, i, (*graph)[i][j])
		}
	}
	var i, j, k int

	for i = 0; i < len(*graph); i++ {
		for j = 0; j < len(*graph); j++ {
			for k = 0; k < len(*graph); k++ {
				if dist[i][j] > (dist[i][k]+dist[k][j]) && (dist[k][j] != Inf && dist[i][k] != Inf) {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

func MinDistance(dist []int, sptSet []bool, V int) int {
	min := Inf
	minIndex := 0

	for v := 0; v < V; v++ {
		if sptSet[v] == false && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}

	return minIndex
}

func Dijkstra(graph [][]int, src int) []int {
	var dist []int
	var sptSet []bool

	for i := 0; i < len(graph); i++ {
		dist = append(dist, Inf)
		sptSet = append(sptSet, false)
	}

	dist[src] = 0

	for count := 0; count < len(graph)-1; count++ {
		u := MinDistance(dist, sptSet, len(graph))
		sptSet[u] = true
		for v := 0; v < len(graph); v++ {
			if !sptSet[v] && graph[u][v] != 0 && dist[u] != Inf && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	return dist
}
