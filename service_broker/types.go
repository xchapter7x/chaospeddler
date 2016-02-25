package chaospeddler

import (
	"net/http"
	"time"

	"github.com/gronpipmaster/mgodb"
	"github.com/pivotal-cf/brokerapi"
	"github.com/xchapter7x/cloudcontroller-client"
	"gopkg.in/mgo.v2/bson"
)

//ServiceBroker - this is the struct containing chaos peddler logic
type ServiceBroker struct {
	Orchestrator
}

//Orchestrator - it's chaos but it's organized. implement this interface to
//run the chaos show.
type Orchestrator interface {
	Start()
}

//Maestro - implements the chaosOrchestrater interface for real use in the
//service
type Maestro struct {
	AppKiller AppInstanceKiller
}

//AppInstanceKiller - an interface which kills app instances
type AppInstanceKiller interface {
	KillPercent(ServiceBinding, int) (map[string]int, error)
}

//AppKill - implements AppInstanceKiller to kill apps
type AppKill struct {
	CloudController       *ccclient.Client
	HTTPClient            *http.Client
	CloudControllerAPIURL string
}

//BaseBrokerModel - base struct describing a model to extend
type BaseBrokerModel struct {
	mgodb.Model `,inline`
	ID          bson.ObjectId `bson:"_id,omitempty" 	json:"Id,omitempty"`
	InstanceID  string        `bson:"instance_id,omitempty"		json:"InstanceId,omitempty"`
	Active      bool          `bson:"active,omitempty"		json:"Active,omitempty"`
	Created     time.Time     `bson:"created,omitempty"		json:"Created,omitempty"`
	Deleted     time.Time     `bson:"deleted,omitempty"		json:"Deleted,omitempty"`
}

//ServiceInstance - model to persist service instance information
type ServiceInstance struct {
	ServiceID        string                 `json:"service_id"`
	PlanID           string                 `json:"plan_id"`
	OrganizationGUID string                 `json:"organization_guid"`
	SpaceGUID        string                 `json:"space_guid"`
	Parameters       map[string]interface{} `json:"parameters,omitempty"`
	BaseBrokerModel
}

//ServiceBinding - model to persist service binding information
type ServiceBinding struct {
	BaseBrokerModel
	AppGUID    string                 `json:"app_guid"`
	PlanID     string                 `json:"plan_id"`
	ServiceID  string                 `json:"service_id"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	BindingID  string                 `bson:"instance_id,omitempty"		json:"InstanceId,omitempty"`
}

type mongoModeler interface {
	Save() error
	FindOne(queryDoc interface{}, doc interface{}) (err error)
	ReloadDoc(doc interface{})
}

//BindingProvisioner - interface defining a object which can provision and
//unprovision a service binding
type BindingProvisioner interface {
	mongoModeler
	SetInstanceID(string)
	SetBindingID(string)
	SetBindDetails(brokerapi.BindDetails)
	SetActive(bool)
	FindAllMatches() ([]ServiceBinding, error)
}

//InstanceProvisioner - interface defining a object which can provision and
//unprovision a service instance
type InstanceProvisioner interface {
	mongoModeler
	SetInstanceID(string)
	SetProvisionDetails(brokerapi.ProvisionDetails)
	SetActive(bool)
}
