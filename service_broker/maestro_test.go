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

			It(fmt.Sprintf("Then it should select a kill group size around %d but not greater than the result set", KillGroupPercentSelector), func() {
				maestroFunc()
				min := 0
				max := len(f.FakeQueryResponse)
				Ω(*appKiller.KillCounter).Should(BeNumerically(">=", min))
				Ω(*appKiller.KillCounter).Should(BeNumerically("<", max))
			})
		})
	})
}
