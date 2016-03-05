package chaospeddler

import (
	"encoding/json"
	"time"

	"github.com/pivotal-cf/brokerapi"
)

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
	b, _ := json.Marshal(dt.Parameters)
	s.Parameters = string(b)
}

//SetActive - setter
func (s *ServiceBinding) SetActive(active bool) {
	s.Active = active

	if s.Active {
		s.CreatedAt = time.Now()
	} else {
		s.DeletedAt = time.Now()
	}
}

//FindAllMatches - find all bindings meeting the given binding signature
func (s *ServiceBinding) FindAllMatches() (results []ServiceBinding, err error) {
	return
}

//GetCName - gives collection name to object when saved
func (*ServiceBinding) GetCName() string {
	return "servicebinding"
}
