package trees_test

import (
	. "github.com/amaury95/go-competitive/trees"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tree", func() {
	var tree Tree

	BeforeEach(func() {
		tree = Node(3)
		tree.AddChild(Node(2), Node(6))
	})

	Describe("Value", func() {
		It("should match the value", func() {
			Expect(tree.Value()).To(Equal(3))
		})
	})

	Describe("IsBinary", func() {
		It("should be binary", func() {
			Expect(tree.IsBinary()).To(BeTrue())
		})

		It("should not be binary now", func() {
			tree.AddChild(Node(5))
			Expect(tree.IsBinary()).NotTo(BeTrue())
		})
	})

	Describe("PreOrder", func() {
		It("should print the right preorder", func() {
			Expect(tree.PreOrder()).To(Equal([]int{3, 2, 6}))
		})
	})
})
