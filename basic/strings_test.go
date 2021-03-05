package basic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/amaury95/go-competitive/basic"
)

var _ = Describe("String Process", func() {
	Describe("SplitIntegers", func() {
		It("should split the given string and return an array of integers", func() {
			values := "1 2 3 4 5"
			result := []int{1, 2, 3, 4, 5}
			Expect(SplitIntegers(values)).To(Equal(result))
		})
	})

	Describe("Round4", func() {
		It("should round the number by 4 floating points by exceding", func() {
			var value float64 = 10.123456
			var result string = "10.1235"
			Expect(Round4(value)).To(Equal(result))
		})
		It("should round the number by 4 floating points by deceding", func() {
			var value float64 = 10.123432
			var result string = "10.1234"
			Expect(Round4(value)).To(Equal(result))
		})
	})
})
