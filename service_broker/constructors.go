package chaospeddler

import (
	"crypto/tls"
	"net/http"

	"github.com/xchapter7x/cloudcontroller-client"
	"github.com/xchapter7x/lo"

	"gopkg.in/mgo.v2/bson"
)

//NewAppKill - construct a new app kill object
func NewAppKill(username, password, loginurl, ccurl string) (a *AppKill) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}

	if client, err := ccclient.New(loginurl, username, password, httpClient).Login(); err == nil {
		a = &AppKill{
			CloudController:       client,
			HttpClient:            httpClient,
			CloudControllerAPIURL: ccurl,
		}
	} else {
		lo.G.Error("there was an error logging in", err)
	}
	return
}

//NewMaestro - constructor for a maestro object
var NewMaestro = func(username, password, loginurl, ccurl string) (m *Maestro) {
	m = new(Maestro)
	m.AppKiller = NewAppKill(username, password, loginurl, ccurl)
	return
}

//NewServiceInstance - constructor for a serviceinstance object
var NewServiceInstance = func(setID bool) InstanceProvisioner {
	model := new(ServiceInstance)

	if setID {
		model.ID = bson.NewObjectId()
	}
	model.SetDoc(model)
	return model
}

//NewServiceBinding - constructor for a servicebinding object
var NewServiceBinding = func(setID bool) BindingProvisioner {
	model := new(ServiceBinding)

	if setID {
		model.ID = bson.NewObjectId()
	}
	model.SetDoc(model)
	return model
}
