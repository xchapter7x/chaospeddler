package chaospeddler_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/chaospeddler/service_broker"
	"github.com/xchapter7x/chaospeddler/service_broker/fake"
)

var _ = Describe("Given a Maestro", func() {

	var maestro = new(Maestro)
	Describe("Given a PollCrazyPlans method", func() {
		Context("When called on a valid set of crazy plans", func() {

			var origNewServiceBinding func(bool) BindingProvisioner
			var f *fake.BaseService
			var appKiller *fake.AppKill
			BeforeEach(func() {
				maestro = new(Maestro)
				appKiller = fake.NewAppKill(nil)
				maestro.AppKiller = appKiller
				f = new(fake.BaseService)
				f.FakeQueryResponse = fake.GenerateQueryResponse()
				f.DocsAssignment = fake.ServiceBindingDocsAssignment
				origNewServiceBinding = NewServiceBinding
				NewServiceBinding = f.NewServiceBinding
			})

			AfterEach(func() {
				NewServiceBinding = origNewServiceBinding
			})
			It(fmt.Sprintf("Then it should select a subset randomly and kill % of the selected apps' AIs", KillPercentCrazy), func() {
				maestro.PollCrazyPlans()
				min := 1
				max := int(float64(len(f.FakeQueryResponse.([]ServiceBinding))) * float64(KillGroupSize))
				Ω(*appKiller.KillCounter).Should(BeNumerically(">=", min))
				Ω(*appKiller.KillCounter).Should(BeNumerically("<", max))
			})
		})
	})
})
