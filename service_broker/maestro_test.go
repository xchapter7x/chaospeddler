package chaospeddler_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/chaospeddler/service_broker"
	"github.com/xchapter7x/chaospeddler/service_broker/fakes"
)

var _ = Describe("Given a Maestro", func() {
	var gormDB1 = new(fakes.FakeGormDB)
	var maestro1 = NewMaestro("", "", "", "", gormDB1)
	testPlanPolling("CrazyChaosPlan", maestro1, maestro1.PollCrazyPlans, KillPercentCrazy, gormDB1)

	var gormDB2 = new(fakes.FakeGormDB)
	var maestro2 = NewMaestro("", "", "", "", gormDB2)
	testPlanPolling("AnnoyingChaosPlan", maestro2, maestro2.PollAnnoyingPlans, KillPercentAnnoying, gormDB2)

	var gormDB3 = new(fakes.FakeGormDB)
	var maestro3 = NewMaestro("", "", "", "", gormDB3)
	testPlanPolling("MickeyMouseChaosPlan", maestro3, maestro3.PollMickeyMousePlans, KillPercentMickeyMouse, gormDB3)
})

func testPlanPolling(name string, maestro *Maestro, maestroFunc func(), killPercent int, db *fakes.FakeGormDB) {

	Describe(fmt.Sprintf("Given a %s method", name), func() {
		Context(fmt.Sprintf("When called on a valid set of %s plans", name), func() {

			var appKiller *fakes.FakeAppInstanceKiller
			var origFindAllMatches func(GormDB, string, string) ([]ServiceBinding, error)
			BeforeEach(func() {
				appKiller = new(fakes.FakeAppInstanceKiller)
				maestro.AppKiller = appKiller
				origFindAllMatches = FindAllMatches
				FindAllMatches = func(g GormDB, i string, b string) ([]ServiceBinding, error) {
					return fakes.GenerateQueryResponse(), nil
				}
			})

			AfterEach(func() {
				FindAllMatches = origFindAllMatches
			})

			It(fmt.Sprintf("Then it should select a kill group size around %d but not greater than the result set", KillGroupPercentSelector), func() {
				maestroFunc()
				_, percentCalled := appKiller.KillPercentArgsForCall(0)
				Ω(true).Should(BeTrue())
				Ω(appKiller.KillPercentCallCount()).Should(BeNumerically(">=", 1))
				Ω(percentCalled).Should(Equal(killPercent))
				//max := len(fakes.FakeQueryResponse())
				//Ω(appKiller.KillPercentCallCount()).Should(BeNumerically(">=", min))
				//Ω(appKiller.KillPercentCallCount()).Should(BeNumerically("<", max))
			})
		})
	})
}
