package multierror_test

import (
	"fmt"

	. "github.com/cloudfoundry/multierror"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Multierror", func() {
	Describe("Add", func() {
		It("adds an error", func() {
			m := MultiError{}
			m.Add(fmt.Errorf("Sample Error"))
			Expect(m.HasAny()).To(BeTrue())
			Expect(m.Error()).To(ContainSubstring("Sample Error"))
		})

		Context("adding a MultiError", func() {
			It("adds all the errors", func() {
				m1 := MultiError{}
				m1.Add(fmt.Errorf("Error 1"))
				m1.Add(fmt.Errorf("Error 2"))
				m2 := MultiError{}
				m2.Add(fmt.Errorf("Error 3"))
				m2.Add(fmt.Errorf("Error 4"))

				m1.Add(m2)

				Expect(m1.Error()).To(ContainSubstring("Error 1"))
				Expect(m1.Error()).To(ContainSubstring("Error 2"))
				Expect(m1.Error()).To(ContainSubstring("Error 3"))
				Expect(m1.Error()).To(ContainSubstring("Error 4"))
			})
		})
	})

	Describe("AddWithPrefix", func() {
		It("adds an error with prefix", func() {
			m := MultiError{}
			m.AddWithPrefix(fmt.Errorf("Error 1"), "Prefix:")

			Expect(m.Error()).To(ContainSubstring("Prefix:Error 1"))
		})

		Context("adding a MultiError", func() {
			It("adds all the errors with prefix", func() {
				m1 := MultiError{}
				m1.AddWithPrefix(fmt.Errorf("Error 1"), "Prefix:")
				m1.AddWithPrefix(fmt.Errorf("Error 2"), "Prefix:")
				m2 := MultiError{}
				m2.Add(fmt.Errorf("Error 3"))
				m2.Add(fmt.Errorf("Error 4"))

				m1.AddWithPrefix(m2, "Prefix:")

				Expect(m1.Error()).To(ContainSubstring("Prefix:Error 1"))
				Expect(m1.Error()).To(ContainSubstring("Prefix:Error 2"))
				Expect(m1.Error()).To(ContainSubstring("Prefix:Error 3"))
				Expect(m1.Error()).To(ContainSubstring("Prefix:Error 4"))
			})
		})
	})

	Describe("Error", func() {
		It("prints all the errors", func() {
			m := MultiError{}
			m.Add(fmt.Errorf("Error 1"))
			m.Add(fmt.Errorf("Error 2"))

			Expect(m.Error()).To(ContainSubstring("encountered 2 errors during validation"))
			Expect(m.Error()).To(ContainSubstring("Error 1"))
			Expect(m.Error()).To(ContainSubstring("Error 2"))
		})

		It("prints all the errors", func() {
			m := MultiError{}
			m.Add(fmt.Errorf("Error 1"))
			m.Add(fmt.Errorf("Error 2"))

			Expect(m.Error()).To(ContainSubstring("encountered 2 errors during validation"))
			Expect(m.Error()).To(ContainSubstring("Error 1"))
			Expect(m.Error()).To(ContainSubstring("Error 2"))
		})
	})

	Describe("HasAny", func() {
		Context("when there are errors", func() {
			It("returns true", func() {
				m := MultiError{}
				m.Add(fmt.Errorf("Error 1"))
				Expect(m.HasAny()).To(BeTrue())
			})
		})

		Context("when there are no errors", func() {
			It("returns false", func() {
				m := MultiError{}
				Expect(m.HasAny()).To(BeFalse())
			})
		})
	})
})
