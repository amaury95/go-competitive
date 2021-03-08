package graphs_test

import (
	. "github.com/amaury95/go-competitive/graphs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bfs", func() {
	It("should find a path from 0 to 7", func() {
		path, err := Bfs(graph, 0, 7)
		expected := []int{0, 4, 7}
		Expect(err).NotTo(HaveOccurred())
		Expect(path).To(Equal(expected))
	})
})
