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
	dbInfo := ExtractDBInfo()
	basicAuthInfo := ExtractBasicAuthInfo()
	cloudControllerInfo := ExtractCloudControllerInfo()
	chaos := chaospeddler.NewServiceBroker(chaospeddler.NewMaestro(cloudControllerInfo.Username, cloudControllerInfo.Password, cloudControllerInfo.LoginURL, cloudControllerInfo.CCURL))
	chaos.Start()
	logger := lager.NewLogger("chaos-peddler-servicebroker")
	credentials := brokerapi.BrokerCredentials{
		Username: basicAuthInfo.Username,
		Password: basicAuthInfo.Password,
	}
	var dbm *mgodb.Dbm
	_ = dbm.Init(dbInfo.ConnectionURL, dbInfo.DBName, 10)
	brokerAPI := brokerapi.New(chaos, logger, credentials)
	http.Handle("/", brokerAPI)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func ExtractDBInfo() (dbInfo DBInfo) {
	return
}

func ExtractBasicAuthInfo() (basicAuthInfo BasicAuthInfo) {
	return
}

func ExtractCloudControllerInfo() (cloudControllerInfo CloudControllerInfo) {
	return
}

type CloudControllerInfo struct {
	Username string
	Password string
	LoginURL string
	CCURL    string
}

type BasicAuthInfo struct {
	Username string
	Password string
}

type DBInfo struct {
	ConnectionURL string
	DBName        string
}
