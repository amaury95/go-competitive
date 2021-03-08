package graphs

import (
	"errors"

	"github.com/amaury95/go-competitive/basic"
)

// Dfs ...
func Dfs(graph map[int][]int, from int, to int) (path []int, err error) {
	visited := make(map[int]bool, len(graph))
	elements := &basic.Stack{}

	if _, ok := graph[from]; !ok {
		return nil, errors.New("element do not exist")
	}

	elements.Push(from)

	for elements.Len() > 0 {
		current, _ := elements.Peek()
		 
		if current == to {
			break
		}

		visited[current] = true

		for _, edge := range graph[current] {
			if !visited[edge] {
				elements.Push(edge)
				break
			}
		}

		if v, _ := elements.Peek(); v == current {
			elements.Pop()
		}
	}

	return elements.Values(), nil
}
