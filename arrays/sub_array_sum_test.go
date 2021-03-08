package arrays_test

import (
	"fmt"

	. "github.com/amaury95/go-competitive/arrays"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sub Array With Given Sum", func() {
	Context("given the cases", func() {
		var cases = []struct {
			inputArray     []int
			inputSum       int
			expectedOutput []int
		}{
			{
				inputArray:     []int{1, 2, 3, 4},
				inputSum:       5,
				expectedOutput: []int{2, 3},
			},
			{
				inputArray:     []int{1, 2, 3, 4},
				inputSum:       10,
				expectedOutput: []int{1, 2, 3, 4},
			},
			{
				inputArray:     []int{1, 2, 3, 1, 4, 5, 5, 3, 2, 1, 2, 3, 2, 6, 3, 4, 5},
				inputSum:       14,
				expectedOutput: []int{4, 5, 5},
			},
		}

		Describe("Iterative solution", func() {
			for i, c := range cases {
				It(fmt.Sprintf("it should pass the %d test case", i), func() {
					Expect(IterativeSumSubArray(c.inputArray, c.inputSum)).To(Equal(c.expectedOutput))
				})
			}
		})
	})
})
