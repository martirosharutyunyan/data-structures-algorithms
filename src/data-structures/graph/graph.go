package graph

import (
	"fmt"
)

type Graph struct {
	Vertexes map[int][]int
}

func InsertMatrix(matrix *[][]int, x, y int) {
	if len(*matrix) < x+1 {
		temp := make([][]int, x+1)
		for i := 0; i < len(*matrix); i++ {
			subMatrix := []int{}
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

	queue := [][]int{}

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

func (graph *Graph) PossiblePaths(start, end int) [][]int {
	results := [][]int{}
	stack := []int{start}
	temp := []int{}
	for len(stack) != 0 {
		front := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		temp = append(temp, front)
		if temp[len(temp)-1] == end {
			results = append(results, temp)
			temp = []int{start}
		}
		stack = append(stack, graph.Vertexes[front]...)
	}

	return results
}

func (graph *Graph) AddVertex(source int, dst []int) {
	graph.Vertexes[source] = append(graph.Vertexes[source], dst...)
}
