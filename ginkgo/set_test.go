package adt_test

import (
	adt "go-tdd/ginkgo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var s *adt.Set
var _ = Describe("Set", func() {

	BeforeEach(func ()  {
		s = adt.NewSet()
	})

	Describe("Emptiness", func() {
		Context("When the set does not contain items", func() {
			It("Should be empty", func() {
				Expect(s.IsEmpty()).To(BeTrue())
			})
		})

		Context("When the set contains items", func() {
			It("Should not be empty", func ()  {
				s.Add("Red")

				Expect(s.IsEmpty()).To(BeFalse())
			})
		})
	})

	Describe("Size", func() {
		Context("As items are added", func() {
			It("Should return an increasing size", func ()  {
				By("Set size being 0", func() {
					Expect(s.Size()).To(BeZero())
				})

				By("Adding the first item", func() {
					s.Add("red")

					Expect(s.Size()).To(Equal(1))
				})

				By("Adding the second item", func() {
					s.Add("blue")

					Expect(s.Size()).To(Equal(2))
				})
			})
		})
	})

	Describe("Contains", func() {
		Context("When red has not been added", func() {
			It("Should not contain red", func ()  {
				Expect(s.Contains("red")).To(BeFalse())
			})
		})

		Context("When red has been added", func() {
			It("Should contain red", func ()  {
				s.Add("red")
				Expect(s.Contains("red")).To(BeTrue())
			})
		})

		Context("When blue has been added", func() {
			It("Should contain blue", func ()  {
				s.Add("blue")
				Expect(s.Contains("blue")).To(BeTrue())
			})
		})
	})

})
