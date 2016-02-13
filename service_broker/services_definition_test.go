package chaospeddler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/chaospeddler/service_broker"
)

var _ = Describe("Given a Services() Method", func() {
	Context("When called", func() {
		It("Then we should get a service object with the proper number of chaos plans", func() {
			isSingleService := 0
			properPlanCount := 4
			Ω(len(new(ServiceBroker).Services()[isSingleService].Plans)).Should(Equal(properPlanCount))
		})
	})
})
