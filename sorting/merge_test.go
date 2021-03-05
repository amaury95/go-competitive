package sorting_test

import (
	. "github.com/amaury95/go-competitive/sorting"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MergeSort", func() {

	It("should return the sorted array by using merge sort", func() {
		Expect(MergeSort(Array)).To(Equal(Sorted))
	})
})
