package main

import (
	"net/http"
	"os"

	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-golang/lager"

	"github.com/xchapter7x/chaospeddler/service_broker"
)

func main() {
	serviceBroker := &chaospeddler.ServiceBroker{}
	logger := lager.NewLogger("chaos-peddler-servicebroker")
	credentials := brokerapi.BrokerCredentials{
		Username: "username",
		Password: "password",
	}

	brokerAPI := brokerapi.New(serviceBroker, logger, credentials)
	http.Handle("/", brokerAPI)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
