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

//SetBindDetails - setter
func (s *ServiceBinding) SetBindDetails(dt brokerapi.BindDetails) {
	s.AppGUID = dt.AppGUID
	s.PlanID = dt.PlanID
	s.ServiceID = dt.ServiceID
	s.Parameters = dt.Parameters
}

//SetActive - setter
func (s *ServiceBinding) SetActive(active bool) {
	s.Active = active
}
