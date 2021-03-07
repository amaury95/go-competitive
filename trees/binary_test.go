package trees_test

import (
	. "github.com/amaury95/go-competitive/trees"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinaryTree", func() {
	Context("given the following tree", func() {
		var tree *BinaryTree

		BeforeEach(func() {
			tree = &BinaryTree{
				Value:      1,
				LeftChild:  &BinaryTree{Value: 2},
				RightChild: &BinaryTree{Value: 3},
			}
		})

		Describe("PreOrder", func() {
			It("should return the elements in the pre order", func() {
				result := []int{1, 2, 3}
				Expect(tree.PreOrder()).To(Equal(result))
				Expect(tree.IterativePreOrder()).To(Equal(result))
			})
		})
		
	})
})
