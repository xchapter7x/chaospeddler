package chaospeddler

import (
	"fmt"

	"github.com/robfig/cron"
	"github.com/xchapter7x/lo"
)

//Start - starts the maestro orchestrating its chaotic ways.
func (s *Maestro) Start() {
	c := cron.New()
	c.AddFunc("@every 1m", s.PollCrazyPlans)
	c.AddFunc("@every 10m", s.PollAnnoyingPlans)
	c.AddFunc("@every 30m", s.PollMickeyMousePlans)
	c.Start()
}

//Db - allows maestro to implement the orachestrator interface, returns the db
//connection
func (s *Maestro) DB() GormDB {
	if err := s.db.Ping(); err != nil {
		lo.G.Error("dropped db connection, reconnecting now: ", err)

		if s.newDBConnection != nil {
			s.db = s.newDBConnection()
		} else {
			lo.G.Error("sorry no reconnect function defined")
		}
	}

	return s.db
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

func (s *Maestro) poll(planid, serviceid string, percent int) {

	if serviceBindings, err := FindAllMatches(s.DB(), planid, serviceid); err == nil {
		killSet := s.extractKillSet(serviceBindings)
		s.kill(killSet, percent)

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

	for i, v := range serviceBindings {

		if percentChanceOfTrue(KillGroupPercentSelector) {
			lo.G.Debug("selected index: ", i)
			killSet = append(killSet, v)
		}
	}
	return
}
