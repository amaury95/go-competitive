package sorting_test

import (
	. "github.com/amaury95/go-competitive/sorting"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BubbleSort", func() {

	It("should return the sorted array by using bubble sort", func() {
		Expect(BubbleSort(Array)).To(Equal(Sorted))
	})
})
