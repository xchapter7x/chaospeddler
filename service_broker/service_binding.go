package chaospeddler

import "github.com/pivotal-cf/brokerapi"

//SetInstanceID - setter
func (s *ServiceBinding) SetInstanceID(id string) {
	s.InstanceID = id
}

//SetBindingID - setter
func (s *ServiceBinding) SetBindingID(id string) {
	s.BindingID = id
}

//SetUnbindDetails - setter
func (s *ServiceBinding) SetUnbindDetails(dt brokerapi.UnbindDetails) {
	s.UnbindDetails = dt
}

//SetBindDetails - setter
func (s *ServiceBinding) SetBindDetails(dt brokerapi.BindDetails) {
	s.BindDetails = dt
}

//SetActive - setter
func (s *ServiceBinding) SetActive(active bool) {
	s.Active = active
}
