package chaospeddler_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/chaospeddler/service_broker"
	"github.com/xchapter7x/chaospeddler/service_broker/fake"
)

var _ = Describe("Given a Maestro", func() {

	var maestro1 = new(Maestro)
	testPlanPolling("CrazyChaosPlan", maestro1, maestro1.PollCrazyPlans, KillPercentCrazy)

	var maestro2 = new(Maestro)
	testPlanPolling("AnnoyingChaosPlan", maestro2, maestro2.PollAnnoyingPlans, KillPercentAnnoying)

	var maestro3 = new(Maestro)
	testPlanPolling("MickeyMouseChaosPlan", maestro3, maestro3.PollMickeyMousePlans, KillPercentMickeyMouse)
})

func testPlanPolling(name string, maestro *Maestro, maestroFunc func(), killPercent int) {

	Describe(fmt.Sprintf("Given a %s method", name), func() {
		Context(fmt.Sprintf("When called on a valid set of %s plans", name), func() {

			var origNewServiceBinding func(bool) BindingProvisioner
			var f *fake.BaseService
			var appKiller *fake.AppKill
			BeforeEach(func() {
				appKiller = fake.NewAppKill(nil)
				maestro.AppKiller = appKiller
				f = new(fake.BaseService)
				f.FakeQueryResponse = fake.GenerateQueryResponse()
				origNewServiceBinding = NewServiceBinding
				NewServiceBinding = f.NewServiceBinding
			})

			AfterEach(func() {
				NewServiceBinding = origNewServiceBinding
			})
			for i := 1; i < 200; i++ {
				It(fmt.Sprintf("Then it should select a kill group size no greaater than %v of the result set", KillGroupSize), func() {
					maestroFunc()
					min := 1
					max := int(float64(len(f.FakeQueryResponse)) * float64(KillGroupSize))
					Ω(*appKiller.KillCounter).Should(BeNumerically(">=", min))
					Ω(*appKiller.KillCounter).Should(BeNumerically("<", max))
				})
			}
		})
	})
}
