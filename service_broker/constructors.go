package chaospeddler

import (
	"net/http"

	"github.com/xchapter7x/cloudcontroller-client"

	"gopkg.in/mgo.v2/bson"
)

var NewMaestro = func(username, password, loginurl string) (m *Maestro) {
	m = new(Maestro)
	m.AppKiller = &AppKill{
		CloudController: ccclient.New(loginurl, username, password, new(http.Client)),
	}
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
