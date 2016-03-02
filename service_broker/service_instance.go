package chaospeddler

import (
	"encoding/json"
	"time"

	"github.com/pivotal-cf/brokerapi"
)

//SetInstanceID - setter
func (s *ServiceInstance) SetInstanceID(id string) {
	s.InstanceID = id
}

//SetProvisionDetails - setter
func (s *ServiceInstance) SetProvisionDetails(dt brokerapi.ProvisionDetails) {
	s.PlanID = dt.PlanID
	s.OrganizationGUID = dt.OrganizationGUID
	s.SpaceGUID = dt.SpaceGUID
	b, _ := json.Marshal(dt.Parameters)
	s.Parameters = string(b)
	s.ServiceID = dt.ServiceID
}

//SetActive - setter
func (s *ServiceInstance) SetActive(active bool) {
	s.Active = active

	if s.Active {
		s.Created = time.Now()
	} else {
		s.Deleted = time.Now()
	}
}

//GetCName - gives collection name to object when saved
func (*ServiceInstance) GetCName() string {
	return "serviceinstance"
}
