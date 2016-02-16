package chaospeddler

import (
	"errors"

	"github.com/pivotal-cf/brokerapi"
)

//NewServiceBroker - service broker constructor
func NewServiceBroker(orch Orchestrator) *ServiceBroker {
	return &ServiceBroker{
		Orchestrator: orch,
	}
}

//Provision a new instance here. If async is allowed, the broker can still
// chose to provision the instance synchronously.
func (*ServiceBroker) Provision(instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	serviceInstance := NewServiceInstance(true)
	serviceInstance.SetInstanceID(instanceID)
	serviceInstance.SetProvisionDetails(details)
	serviceInstance.SetActive(true)
	err := serviceInstance.Save()
	return brokerapi.ProvisionedServiceSpec{
		IsAsync: false,
	}, err
}

//LastOperation If the broker provisions asynchronously, the Cloud Controller will poll this endpoint
// for the status of the provisioning operation.
// This also applies to deprovisioning (work in progress).
func (*ServiceBroker) LastOperation(instanceID string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, errors.New("something is wrong. this service is not currently asynchronous")
}

//Deprovision a new instance here. If async is allowed, the broker can still
// chose to deprovision the instance synchronously, hence the first return value.
func (*ServiceBroker) Deprovision(instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.IsAsync, error) {
	serviceInstance := NewServiceInstance(false)
	serviceInstance.SetInstanceID(instanceID)
	serviceInstance.FindOne(serviceInstance, &serviceInstance)
	serviceInstance.ReloadDoc(serviceInstance)
	serviceInstance.SetDeprovisionDetails(details)
	serviceInstance.SetActive(false)
	err := serviceInstance.Save()
	return false, err
}

//Bind to instances here
// Return a binding which contains a credentials object that can be marshalled to JSON,
// and (optionally) a syslog drain URL.
func (*ServiceBroker) Bind(instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	serviceBinding := NewServiceBinding(true)
	serviceBinding.SetInstanceID(instanceID)
	serviceBinding.SetBindingID(bindingID)
	serviceBinding.SetBindDetails(details)
	serviceBinding.SetActive(true)
	err := serviceBinding.Save()
	return brokerapi.Binding{}, err
}

//Unbind from instances here
func (*ServiceBroker) Unbind(instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	serviceBinding := NewServiceBinding(false)
	serviceBinding.SetInstanceID(instanceID)
	serviceBinding.SetBindingID(bindingID)
	serviceBinding.FindOne(serviceBinding, &serviceBinding)
	serviceBinding.ReloadDoc(serviceBinding)
	serviceBinding.SetUnbindDetails(details)
	serviceBinding.SetActive(false)
	err := serviceBinding.Save()
	return err
}

//Update instance here
func (*ServiceBroker) Update(instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.IsAsync, error) {
	return false, errors.New("this functionality is not yet implemented")
}
