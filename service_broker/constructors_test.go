package chaospeddler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/chaospeddler/service_broker"
)

var _ = Describe("Given some constructors", func() {
	testSetIDCases("NewServiceInstance", func(a bool) string {
		return NewServiceInstance(a).(*ServiceInstance).ID.Hex()
	})

	testSetIDCases("NewServiceBinding", func(a bool) string {
		return NewServiceBinding(a).(*ServiceBinding).ID.Hex()
	})

})

func testSetIDCases(name string, f func(bool) string) {
	Describe("Given a "+name+" constructor", func() {
		Context("When called with a true setID", func() {

			It("then it should set a new object id", func() {
				id := f(true)
				Ω(id).ShouldNot(BeEmpty())
			})
		})

		Context("When called with a false setID", func() {

			It("then it should set a new object id", func() {
				id := f(false)
				Ω(id).Should(BeEmpty())
			})
		})
	})
}
