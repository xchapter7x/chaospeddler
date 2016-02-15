package chaospeddler

import "gopkg.in/mgo.v2/bson"

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
