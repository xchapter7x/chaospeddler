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

//Model - base model struct for sql
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
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

//GormDBWrapper - wraps a gorm db so we can implement a cleaner interface
type GormDBWrapper struct {
	DBWrapper
}

//DBWrapper - a struct to wrap the gorm.DB
type DBWrapper struct {
	*gorm.DB
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
	AddError(err error) error
	AddForeignKey(field string, dest string, onDelete string, onUpdate string) *gorm.DB
	AddIndex(indexName string, column ...string) *gorm.DB
	AddUniqueIndex(indexName string, column ...string) *gorm.DB
	Assign(attrs ...interface{}) *gorm.DB
	Association(column string) *gorm.Association
	Attrs(attrs ...interface{}) *gorm.DB
	AutoMigrate(values ...interface{}) *gorm.DB
	Begin() *gorm.DB
	Close() error
	Commit() *gorm.DB
	Count(value interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	CreateTable(values ...interface{}) *gorm.DB
	CurrentDatabase() string
	DB() *sql.DB
	Debug() *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	DropColumn(column string) *gorm.DB
	DropTable(values ...interface{}) *gorm.DB
	DropTableIfExists(values ...interface{}) *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	FirstOrCreate(out interface{}, where ...interface{}) *gorm.DB
	FirstOrInit(out interface{}, where ...interface{}) *gorm.DB
	Get(name string) (value interface{}, ok bool)
	GetErrors() (errors []error)
	Group(query string) *gorm.DB
	HasTable(value interface{}) bool
	Having(query string, values ...interface{}) *gorm.DB
	InstantSet(name string, value interface{}) *gorm.DB
	Joins(query string) *gorm.DB
	Last(out interface{}, where ...interface{}) *gorm.DB
	Limit(value interface{}) *gorm.DB
	LogMode(enable bool) *gorm.DB
	Model(value interface{}) *gorm.DB
	ModifyColumn(column string, typ string) *gorm.DB
	New() *gorm.DB
	NewRecord(value interface{}) bool
	NewScope(value interface{}) *gorm.Scope
	Not(query interface{}, args ...interface{}) *gorm.DB
	Offset(value interface{}) *gorm.DB
	Omit(columns ...string) *gorm.DB
	Or(query interface{}, args ...interface{}) *gorm.DB
	Order(value string, reorder ...bool) *gorm.DB
	Pluck(column string, value interface{}) *gorm.DB
	Preload(column string, conditions ...interface{}) *gorm.DB
	Raw(sql string, values ...interface{}) *gorm.DB
	RecordNotFound() bool
	Related(value interface{}, foreignKeys ...string) *gorm.DB
	RemoveIndex(indexName string) *gorm.DB
	Rollback() *gorm.DB
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	Save(value interface{}) *gorm.DB
	Scan(dest interface{}) *gorm.DB
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Set(name string, value interface{}) *gorm.DB
	SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface)
	SingularTable(enable bool)
	Table(name string) *gorm.DB
	Unscoped() *gorm.DB
	Update(attrs ...interface{}) *gorm.DB
	UpdateColumn(attrs ...interface{}) *gorm.DB
	UpdateColumns(values interface{}) *gorm.DB
	Updates(values interface{}, ignoreProtectedAttrs ...bool) *gorm.DB
	Where(query interface{}, args ...interface{}) GormDB
}
