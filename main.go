package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/gronpipmaster/mgodb"
	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-golang/lager"

	"github.com/xchapter7x/chaospeddler/service_broker"
	"github.com/xchapter7x/lo"
)

func main() {
	dbInfo := ExtractDBInfo()
	basicAuthInfo := ExtractBasicAuthInfo()
	cloudControllerInfo := ExtractCloudControllerInfo()
	lo.G.Debug("cloud controller", cloudControllerInfo)
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
	appEnv, _ := cfenv.Current()
	service, _ := appEnv.Services.WithName("db-info")
	dbInfo.ConnectionURL = fmt.Sprintf("%v", service.Credentials["uri"])
	dbInfo.DBName = fmt.Sprintf("%v", service.Credentials["database"])
	return
}

func ExtractBasicAuthInfo() (basicAuthInfo BasicAuthInfo) {
	appEnv, _ := cfenv.Current()
	service, _ := appEnv.Services.WithName("basic-auth-info")
	basicAuthInfo.Username = fmt.Sprintf("%v", service.Credentials["username"])
	basicAuthInfo.Password = fmt.Sprintf("%v", service.Credentials["password"])
	return
}

func ExtractCloudControllerInfo() (cloudControllerInfo CloudControllerInfo) {
	appEnv, _ := cfenv.Current()
	service, _ := appEnv.Services.WithName("cloud-controller-info")
	cloudControllerInfo.Username = fmt.Sprintf("%v", service.Credentials["username"])
	cloudControllerInfo.Password = fmt.Sprintf("%v", service.Credentials["password"])
	cloudControllerInfo.LoginURL = fmt.Sprintf("%v", service.Credentials["login-url"])
	cloudControllerInfo.CCURL = fmt.Sprintf("%v", service.Credentials["cc-url"])
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
