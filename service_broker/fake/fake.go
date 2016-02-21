package fake

import (
	"fmt"
	"sync/atomic"

	"github.com/gronpipmaster/mgodb"
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
	model.DocsAssignment = s.DocsAssignment
	model.FakeQueryResponse = s.FakeQueryResponse
	return model
}

//BaseService ---
type BaseService struct {
	ErrFake           error
	SaveCount         *uint64
	FakeQueryResponse interface{}
	DocsAssignment    func(interface{}, interface{})
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

//FindAll ---
func (s *BaseService) FindAll(query mgodb.Query, docs interface{}) (err error) {
	lo.G.Debug("calling FindAll fake", docs, s.FakeQueryResponse)
	s.DocsAssignment(docs, s.FakeQueryResponse)
	return
}

//ServiceBindingDocsAssignment ----
func ServiceBindingDocsAssignment(dst interface{}, src interface{}) {
	var dstTmp []chaospeddler.ServiceBinding
	for _, v := range src.([]chaospeddler.ServiceBinding) {
		dstTmp = append(dstTmp, v)
	}
	*(dst.(*[]chaospeddler.ServiceBinding)) = dstTmp
}

//GenerateQueryResponse ---
func GenerateQueryResponse() []chaospeddler.ServiceBinding {
	var res []chaospeddler.ServiceBinding

	for i := 0; i < 100; i++ {

		res = append(res, chaospeddler.ServiceBinding{
			BaseBrokerModel: chaospeddler.BaseBrokerModel{
				InstanceID: fmt.Sprintf("randominstance-%v", i),
			},
			BindingID: fmt.Sprintf("something-binding-%v", i),
		})
	}
	return res
}
