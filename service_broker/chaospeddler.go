package chaospeddler

import "github.com/pivotal-cf/brokerapi"

//Provision a new instance here. If async is allowed, the broker can still
// chose to provision the instance synchronously.
func (*ServiceBroker) Provision(instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return brokerapi.ProvisionedServiceSpec{}, nil
}

//LastOperation If the broker provisions asynchronously, the Cloud Controller will poll this endpoint
// for the status of the provisioning operation.
// This also applies to deprovisioning (work in progress).
func (*ServiceBroker) LastOperation(instanceID string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}

//Deprovision a new instance here. If async is allowed, the broker can still
// chose to deprovision the instance synchronously, hence the first return value.
func (*ServiceBroker) Deprovision(instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.IsAsync, error) {
	return true, nil
}

//Bind to instances here
// Return a binding which contains a credentials object that can be marshalled to JSON,
// and (optionally) a syslog drain URL.
func (*ServiceBroker) Bind(instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	return brokerapi.Binding{}, nil
}

//Unbind from instances here
func (*ServiceBroker) Unbind(instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	return nil
}

//Update instance here
func (*ServiceBroker) Update(instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.IsAsync, error) {
	return true, nil
}
