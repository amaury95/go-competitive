package sorting_test

import (
	. "github.com/amaury95/go-competitive/sorting"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("QuickSort", func() {

	It("should return the sorted array by using quicksort", func() {
		Expect(QuickSort(Array)).To(Equal(Sorted))
	})
})
