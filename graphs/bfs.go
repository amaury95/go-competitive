package graphs

import (
	"errors"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

// Bfs ...
func Bfs(graph map[int][]int, from int, to int) (path []int, err error) {
	visited := make(map[int]bool, len(graph))
	parent := make(map[int]int, len(graph))

	q := queue.New()

	if _, ok := graph[from]; !ok {
		return nil, errors.New("element do not exist")
	}

	q.Enqueue(from)

	for q.Len() > 0 {
		node := q.Dequeue().(int)
		for _, edge := range graph[node] {
			if !visited[edge] {
				parent[edge] = node
				if edge == to {
					q = queue.New()
					break
				}
				visited[edge] = true
				q.Enqueue(edge)
			}
		}
	}

	p, ok := parent[to]
	if !ok {
		return
	}

	result := stack.New()
	result.Push(to)
	for p != from {
		result.Push(p)
		p = parent[p]
	}
	result.Push(from)

	for result.Len() > 0 {
		path = append(path, result.Pop().(int))
	}

	return
}
