package basic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/amaury95/go-competitive/basic"
)

var _ = Describe("Slices", func() {
	Context("given an array", func() {
		s := []string{"a", "b", "c"}

		Describe("Push", func() {
			It("should append an element to the end of the array", func() {
				Expect(Push(s, "d")).To(Equal([]string{"a", "b", "c", "d"}))
			})
		})
		Describe("Pop", func() {
			It("should remove the last index of the array", func() {
				Expect(Pop(s)).To(Equal([]string{"a", "b"}))
			})
		})
		Describe("RemoveAt", func() {
			It("Should remove the ith element of the array", func() {
				Expect(RemoveAt(s, 1)).To(Equal([]string{"a", "c"}))
			})
		})
	})
})
