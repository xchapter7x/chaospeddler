package chaospeddler

import (
	"errors"

	"github.com/pivotal-cf/brokerapi"
)

//Provision a new instance here. If async is allowed, the broker can still
// chose to provision the instance synchronously.
func (s *ServiceBroker) Provision(instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	serviceInstance := NewServiceInstance()
	serviceInstance.SetInstanceID(instanceID)
	serviceInstance.SetProvisionDetails(details)
	serviceInstance.SetActive(true)
	err := s.save(serviceInstance)
	return brokerapi.ProvisionedServiceSpec{
		IsAsync: false,
	}, err
}

//LastOperation If the broker provisions asynchronously, the Cloud Controller will poll this endpoint
// for the status of the provisioning operation.
// This also applies to deprovisioning (work in progress).
func (s *ServiceBroker) LastOperation(instanceID string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, errors.New("something is wrong. this service is not currently asynchronous")
}

//Deprovision a new instance here. If async is allowed, the broker can still
// chose to deprovision the instance synchronously, hence the first return value.
func (s *ServiceBroker) Deprovision(instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.IsAsync, error) {
	serviceInstance := s.findOneInstance(instanceID)
	serviceInstance.SetActive(false)
	err := s.save(serviceInstance)
	return false, err
}

//Bind to instances here
// Return a binding which contains a credentials object that can be marshalled to JSON,
// and (optionally) a syslog drain URL.
func (s *ServiceBroker) Bind(instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	serviceBinding := NewServiceBinding()
	serviceBinding.SetInstanceID(instanceID)
	serviceBinding.SetBindingID(bindingID)
	serviceBinding.SetBindDetails(details)
	serviceBinding.SetActive(true)
	err := s.save(serviceBinding)
	return brokerapi.Binding{}, err
}

//Unbind from instances here
func (s *ServiceBroker) Unbind(instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	serviceBinding := s.findOneBinding(instanceID, bindingID)
	serviceBinding.SetActive(false)
	err := s.save(serviceBinding)
	return err
}

func (s *ServiceBroker) Update(instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.IsAsync, error) {
	return false, errors.New("this functionality is not yet implemented")
}

func (s *ServiceBroker) save(model interface{}) (err error) {
	s.Orchestrator.DB().Create(model)
	return
}

func (s *ServiceBroker) findOneBinding(instanceID, bindingID string) (serviceBinding ServiceBinding) {
	s.Orchestrator.DB().Where("instanceid = ? and bindingid = ?", instanceID, bindingID).First(&serviceBinding)
	return
}

func (s *ServiceBroker) findOneInstance(instanceID string) (serviceInstance ServiceInstance) {
	s.Orchestrator.DB().Where("instanceid = ?", instanceID).First(&serviceInstance)
	return
}
