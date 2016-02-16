package main

import (
	"net/http"
	"os"

	"github.com/gronpipmaster/mgodb"
	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-golang/lager"

	"github.com/xchapter7x/chaospeddler/service_broker"
)

func main() {
	chaos := chaospeddler.NewServiceBroker(new(chaospeddler.Maestro))
	chaos.Start()
	logger := lager.NewLogger("chaos-peddler-servicebroker")
	credentials := brokerapi.BrokerCredentials{
		Username: "username",
		Password: "password",
	}
	var dbm *mgodb.Dbm
	_ = dbm.Init("connectUrl", "dbName", 10)
	brokerAPI := brokerapi.New(chaos, logger, credentials)
	http.Handle("/", brokerAPI)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
