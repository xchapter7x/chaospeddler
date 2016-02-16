package chaospeddler

import (
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

}

//ExpireCrazyRunlist -
func (s *Maestro) ExpireCrazyRunlist() {
	lo.G.Info("expire crazy chaos log")
}

//PollAnnoyingPlans -
func (s *Maestro) PollAnnoyingPlans() {

}

//ExpireAnnoyingRunlist -
func (s *Maestro) ExpireAnnoyingRunlist() {
	lo.G.Info("expire annoying chaos log")
}

//PollMickeyMousePlans -
func (s *Maestro) PollMickeyMousePlans() {

}

//ExpireMickeyMouseRunlist -
func (s *Maestro) ExpireMickeyMouseRunlist() {
	lo.G.Info("expire mickeymouse chaos log")
}
