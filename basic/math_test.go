package basic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/amaury95/go-competitive/basic"
)

var _ = Describe("Math", func() {
	Describe("Asb", func() {
		It("should give us the expected asbolute value for the given input", func() {
			var value int = -10
			var result float64 = 10
			Expect(Abs(value)).To(Equal(result))
		})
	})
})