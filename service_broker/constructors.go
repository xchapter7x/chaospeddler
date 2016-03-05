package chaospeddler

import (
	"crypto/tls"
	"net/http"

	"github.com/xchapter7x/cloudcontroller-client"
	"github.com/xchapter7x/lo"
)

//NewServiceBroker - service broker constructor
func NewServiceBroker(orch Orchestrator) *ServiceBroker {
	return &ServiceBroker{
		Orchestrator: orch,
	}
}

//NewAppKill - construct a new app kill object
func NewAppKill(username, password, loginurl, ccurl string) (a *AppKill) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	client := ccclient.New(loginurl, username, password, httpClient)
	lo.G.Debug("cc client: ", client)
	a = &AppKill{
		CloudController:       client,
		HTTPClient:            httpClient,
		CloudControllerAPIURL: ccurl,
	}
	return
}

//NewMaestro - constructor for a maestro object
var NewMaestro = func(username, password, loginurl, ccurl string, db GormDB) (m *Maestro) {
	m = new(Maestro)
	m.db = db
	m.AppKiller = NewAppKill(username, password, loginurl, ccurl)
	return
}

//NewServiceInstance - constructor for a serviceinstance object
var NewServiceInstance = func() InstanceProvisioner {
	model := new(ServiceInstance)
	return model
}

//NewServiceBinding - constructor for a servicebinding object
var NewServiceBinding = func() BindingProvisioner {
	model := new(ServiceBinding)
	return model
}
