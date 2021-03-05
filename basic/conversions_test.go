package basic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/amaury95/go-competitive/basic"
)

var _ = Describe("Type Conversions", func() {
	Describe("ConvertString2Int", func() {
		It("should parse the value to int", func() {
			value := "31"
			Expect(ConvertString2Int(value)).To(Equal(31))
		})
	})

	Describe("ConvertString2Float", func() {
		It("should parse the value to float32", func() {
			value := "3.14"
			Expect(ConvertString2Float(value)).To(Equal(3.14))
		})
	})

	Describe("ConvertString2Boolean", func() {
		It("should parse the value to boolean", func() {
			value := "true"
			Expect(ConvertString2Boolean(value)).To(BeTrue())
		})
	})
	
	Describe("ConvertNumber2String", func() {
		It("should convert an integer to string", func() {
			value := 42
			Expect(ConvertNumber2String(value)).To(Equal("42"))
		})
	})
})
