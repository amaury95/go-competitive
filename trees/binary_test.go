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
				Value: 1,
				LeftChild: &BinaryTree{
					Value: 2,
					LeftChild: &BinaryTree{
						Value: 4,
					},
					RightChild: &BinaryTree{
						Value: 5,
					},
				},
				RightChild: &BinaryTree{
					Value: 3,
					RightChild: &BinaryTree{
						Value: 6,
					},
				},
			}
		})

		Describe("PreOrder", func() {
			result := []int{1, 2, 4, 5, 3, 6}
			It("should return the elements in the pre order using recursive method", func() {
				Expect(tree.RecursivePreOrder()).To(Equal(result))
			})
			It("should return the elements in the pre order using iterative method", func() {
				Expect(tree.IterativePreOrder()).To(Equal(result))
			})
		})

		Describe("InOrder", func() {
			result := []int{4, 2, 5, 1, 3, 6}
			It("should return the elements in the in order using recursive method", func() {
				Expect(tree.RecursiveInOrder()).To(Equal(result))
			})
			It("should return the elements in the in order using iterative method", func() {
				Expect(tree.IterativeInOrder()).To(Equal(result))
			})
		})

	})
})
