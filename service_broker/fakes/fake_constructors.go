package fakes

import (
	"fmt"

	"github.com/xchapter7x/chaospeddler/service_broker"
)

//GenerateQueryResponse ---
func GenerateQueryResponse() []chaospeddler.ServiceBinding {
	var res []chaospeddler.ServiceBinding

	for i := 0; i < 100; i++ {

		res = append(res, chaospeddler.ServiceBinding{
			InstanceID: fmt.Sprintf("randominstance-%v", i),
			BindingID:  fmt.Sprintf("something-binding-%v", i),
		})
	}
	return res
}
