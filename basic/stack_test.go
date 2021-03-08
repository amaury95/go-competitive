package basic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/amaury95/go-competitive/basic"
)

var _ = Describe("Stack", func() {
	Context("given an empty stack", func() {
		var stack *Stack
		BeforeEach(func() {
			stack = &Stack{}
		})

		It("should return an empty array", func() {
			Expect(stack.Values()).To(HaveLen(0))
		})
		It("should have length 0", func() {
			Expect(stack.Len()).To(Equal(0))
		})
		It("should fail poping an element", func() {
			val, err := stack.Pop()
			Expect(val).To(Equal(0))
			Expect(err).NotTo(BeNil())
		})
		Context("adding an element", func(){
			var element = 1
			BeforeEach(func(){
				stack.Push(element)
			})	

			It("should return a one element array", func() {
				Expect(stack.Values()).To(HaveLen(1))
			})
			It("should have length 1", func() {
				Expect(stack.Len()).To(Equal(1))
			})
			It("should succed poping an element", func() {
				val, err := stack.Pop()
				Expect(val).To(Equal(element))
				Expect(err).To(BeNil())
				Expect(stack.Len()).To(Equal(0))
			})
		})
	})
	Context("given a 32 elements stack", func() {
		var stack *Stack
		BeforeEach(func() {
			stack = &Stack{}
			for i := 0; i < 32; i++ {
				stack.Push(i)
			}
		})

		It("should return a 32 length array", func() {
			Expect(stack.Values()).To(HaveLen(32))
		})
		It("should have length 32", func() {
			Expect(stack.Len()).To(Equal(32))
		})
		 
		Context("adding an element", func(){
			var element = 100
			BeforeEach(func(){
				stack.Push(element)
			})	

			It("should return a 33 element array", func() {
				Expect(stack.Values()).To(HaveLen(33))
			})
			It("should have length 33", func() {
				Expect(stack.Len()).To(Equal(33))
			})
			It("should succed poping an element", func() {
				val, err := stack.Pop()
				Expect(val).To(Equal(element))
				Expect(err).To(BeNil())
				Expect(stack.Len()).To(Equal(32))
			})
		})
	})
})
