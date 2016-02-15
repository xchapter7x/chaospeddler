package fake

import (
	"sync/atomic"

	"github.com/xchapter7x/chaospeddler/service_broker"
	"github.com/xchapter7x/lo"
)

//NewServiceInstance - constructor for a serviceinstance object
func (s *BaseService) NewServiceInstance(setID bool) chaospeddler.InstanceProvisioner {
	i := uint64(0)
	s.SaveCount = &i
	model := new(serviceInstance)
	model.ErrFake = s.ErrFake
	model.SaveCount = s.SaveCount
	return model
}

//NewServiceBinding - constructor for a servicebinding object
func (s *BaseService) NewServiceBinding(setID bool) chaospeddler.BindingProvisioner {
	i := uint64(0)
	s.SaveCount = &i
	model := new(serviceBinding)
	model.ErrFake = s.ErrFake
	model.SaveCount = s.SaveCount
	return model
}

//BaseService ---
type BaseService struct {
	ErrFake   error
	SaveCount *uint64
}

type serviceBinding struct {
	BaseService
	chaospeddler.ServiceBinding
}
type serviceInstance struct {
	BaseService
	chaospeddler.ServiceInstance
}

//Save - spy on save calls and fake error responses
func (s *BaseService) Save() error {
	lo.G.Debug("calling save on fake", s.SaveCount)
	atomic.AddUint64(s.SaveCount, 1)
	return s.ErrFake
}
