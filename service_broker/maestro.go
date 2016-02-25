package chaospeddler

import (
	"fmt"

	"github.com/pivotal-cf/brokerapi"
	"github.com/robfig/cron"
	"github.com/xchapter7x/lo"
)

//Start - starts the maestro orchestrating its chaotic ways.
func (s *Maestro) Start() {
	s.cleanup()
	c := cron.New()
	c.AddFunc("@every 1m", s.PollCrazyPlans)
	c.AddFunc("@every 10m", s.PollAnnoyingPlans)
	c.AddFunc("@every 30m", s.PollMickeyMousePlans)
	c.AddFunc("@every 15m", s.ExpireCrazyRunlist)
	c.AddFunc("@every 45m", s.ExpireAnnoyingRunlist)
	c.AddFunc("@every 2h", s.ExpireMickeyMouseRunlist)
	c.Start()
}

func (s *Maestro) cleanup() {
	s.ExpireCrazyRunlist()
	s.ExpireAnnoyingRunlist()
	s.ExpireMickeyMouseRunlist()
}

//PollCrazyPlans -
func (s *Maestro) PollCrazyPlans() {
	s.poll(CrazyChaosPlanID, ChaosPeddlerServiceID, KillPercentCrazy)
}

//PollAnnoyingPlans -
func (s *Maestro) PollAnnoyingPlans() {
	s.poll(AnnoyingChaosPlanID, ChaosPeddlerServiceID, KillPercentAnnoying)
}

//PollMickeyMousePlans -
func (s *Maestro) PollMickeyMousePlans() {
	s.poll(MickeyMouseChaosPlanID, ChaosPeddlerServiceID, KillPercentMickeyMouse)
}

//ExpireCrazyRunlist -
func (s *Maestro) ExpireCrazyRunlist() {
	lo.G.Info("expire crazy chaos log")
}

//ExpireAnnoyingRunlist -
func (s *Maestro) ExpireAnnoyingRunlist() {
	lo.G.Info("expire annoying chaos log")
}

//ExpireMickeyMouseRunlist -
func (s *Maestro) ExpireMickeyMouseRunlist() {
	lo.G.Info("expire mickeymouse chaos log")
}

func (s *Maestro) poll(planid, serviceid string, percent int) {
	queryBinding := NewServiceBinding(false)
	queryBinding.SetBindDetails(brokerapi.BindDetails{
		PlanID:    CrazyChaosPlanID,
		ServiceID: ChaosPeddlerServiceID,
	})
	if serviceBindings, err := queryBinding.FindAllMatches(); err == nil {
		killSet := s.extractKillSet(serviceBindings)
		s.kill(killSet, KillPercentCrazy)
	} else {
		lo.G.Error("there was an error when looking for matching records: ", err)
	}
}

func (s *Maestro) kill(killSet []ServiceBinding, percent int) {

	for _, v := range killSet {
		lo.G.Debug("killing app: ", v)

		if killRatio, err := s.AppKiller.KillPercent(v, percent); err != nil {
			lo.G.Error("error on app kill: ", err, killRatio)
		} else {
			lo.G.Debug("killratio: ", killRatio)
		}
	}
	lo.G.Debug(fmt.Sprintf("killing %v apps", len(killSet)))
}

func (s *Maestro) extractKillSet(serviceBindings []ServiceBinding) (killSet []ServiceBinding) {
	lengthOfSet := int(float64(len(serviceBindings)) * float64(KillGroupSize))
	lengthOfKillSet := random(lengthOfSet/4, lengthOfSet)

	for i := 0; i < lengthOfKillSet; i++ {
		idx := random(0, len(serviceBindings))
		lo.G.Debug("selected index: ", idx)
		killSet = append(killSet, serviceBindings[idx])
	}
	return
}
