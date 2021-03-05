package basic_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/amaury95/go-competitive/basic"
)

var _ = Describe("InputReadLine", func() {
	Context("when input is given", func() {
		It("should read and return the given input", func() {
			var input = "5\n"

			var stdin bytes.Buffer
			stdin.Write([]byte(input))

			Expect(InputReadLine(&stdin)).To(Equal(input))
		})
	})

	Context("when no input is given", func() {
		It("should return empty string", func() {
			var stdin bytes.Buffer
			Expect(InputReadLine(&stdin)).To(Equal(""))
		})
	})
})
