package basic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/amaury95/go-competitive/basic"
)

var _ = Describe("Slices", func() {
	Context("given an array", func() {
		s := []int{1, 2, 3}

		Describe("Push", func() {
			It("should append an element to the end of the array", func() {
				Expect(Push(s, 4)).To(Equal([]int{1, 2, 3, 4}))
			})
		})
		Describe("Pop", func() {
			It("should remove the last index of the array", func() {
				Expect(Pop(s)).To(Equal([]int{1, 2}))
			})
		})
		Describe("RemoveAt", func() {
			It("Should remove the ith element of the array", func() {
				Expect(RemoveAt(s, 1)).To(Equal([]int{1, 3}))
			})
		})
	})
})
