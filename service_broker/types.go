package chaospeddler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pivotal-cf/brokerapi"
	"github.com/xchapter7x/cloudcontroller-client"
)

//ServiceBroker - this is the struct containing chaos peddler logic
type ServiceBroker struct {
	Orchestrator
}

//Maestro - implements the chaosOrchestrater interface for real use in the
//service
type Maestro struct {
	newDBConnection func() GormDB
	db              GormDB
	AppKiller       AppInstanceKiller
}

//AppKill - implements AppInstanceKiller to kill apps
type AppKill struct {
	CloudController       *ccclient.Client
	HTTPClient            *http.Client
	CloudControllerAPIURL string
}

//Model - base model fields
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//ServiceInstance - model to persist service instance information
type ServiceInstance struct {
	Model
	ServiceID        string `json:"service_id"`
	PlanID           string `json:"plan_id"`
	OrganizationGUID string `json:"organization_guid"`
	SpaceGUID        string `json:"space_guid"`
	Parameters       string `json:"parameters,omitempty"`
	InstanceID       string `json:"InstanceId,omitempty"`
	Active           bool   `json:"Active,omitempty"`
}

//ServiceBinding - model to persist service binding information
type ServiceBinding struct {
	Model
	InstanceID string `json:"InstanceId,omitempty"`
	Active     bool   `json:"Active,omitempty"`
	AppGUID    string `json:"app_guid"`
	PlanID     string `json:"plan_id"`
	ServiceID  string `json:"service_id"`
	Parameters string `json:"parameters,omitempty"`
	BindingID  string `json:"InstanceId,omitempty"`
}

//BindingProvisioner - interface defining a object which can provision and
//unprovision a service binding
type BindingProvisioner interface {
	SetInstanceID(string)
	SetBindingID(string)
	SetBindDetails(brokerapi.BindDetails)
	SetActive(bool)
}

//InstanceProvisioner - interface defining a object which can provision and
//unprovision a service instance
type InstanceProvisioner interface {
	SetInstanceID(string)
	SetProvisionDetails(brokerapi.ProvisionDetails)
	SetActive(bool)
}

//Orchestrator - it's chaos but it's organized. implement this interface to
//run the chaos show.
type Orchestrator interface {
	DB() GormDB
	Start()
}

//AppInstanceKiller - an interface which kills app instances
type AppInstanceKiller interface {
	KillPercent(ServiceBinding, int) (map[string]int, error)
}

//GormDB the public interface of a gorm.DB
type GormDB interface {
	Close() error
	Create(value interface{}) *gorm.DB
	DB() *sql.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Raw(sql string, values ...interface{}) *gorm.DB
	Scan(dest interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
}
