package graphs_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGraphs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graph Suite")
}

var graph map[int][]int

func init() {
	graph = make(map[int][]int)
	graph[0] = []int{1, 2, 3, 4}
	graph[1] = []int{}
	graph[2] = []int{8, 5, 6}
	graph[3] = []int{6}
	graph[4] = []int{7}
	graph[5] = []int{}
	graph[6] = []int{7}
	graph[7] = []int{2}
	graph[8] = []int{9}
	graph[9] = []int{0}
}
