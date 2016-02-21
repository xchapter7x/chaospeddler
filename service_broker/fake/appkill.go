package fake

import (
	"sync/atomic"

	"github.com/xchapter7x/chaospeddler/service_broker"
)

//NewAppKill ----
func NewAppKill(e error) *AppKill {
	var kc uint64
	return &AppKill{
		KillCounter:    &kc,
		ErrKillPercent: e,
	}
}

//AppKill ---
type AppKill struct {
	KillCounter    *uint64
	ErrKillPercent error
}

//KillPercent ---
func (s *AppKill) KillPercent(sb chaospeddler.ServiceBinding, p int) (err error) {
	atomic.AddUint64(s.KillCounter, 1)
	return s.ErrKillPercent
}
