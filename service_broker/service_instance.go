package chaospeddler

import "github.com/pivotal-cf/brokerapi"

//SetInstanceID - setter
func (s *ServiceInstance) SetInstanceID(id string) {
	s.InstanceID = id
}

//SetDeprovisionDetails - setter
func (s *ServiceInstance) SetDeprovisionDetails(dt brokerapi.DeprovisionDetails) {
	s.DeprovisionDetails = dt
}

//SetProvisionDetails - setter
func (s *ServiceInstance) SetProvisionDetails(dt brokerapi.ProvisionDetails) {
	s.ProvisionDetails = dt
}

//SetActive - setter
func (s *ServiceInstance) SetActive(active bool) {
	s.Active = active
}
