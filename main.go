package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-golang/lager"

	"github.com/xchapter7x/chaospeddler/service_broker"
	"github.com/xchapter7x/lo"
)

func main() {
	sqlConn := ExtractDBSQL()
	basicAuthInfo := ExtractBasicAuthInfo()
	cloudControllerInfo := ExtractCloudControllerInfo()
	lo.G.Debug("cloud controller", cloudControllerInfo)
	chaos := chaospeddler.NewServiceBroker(
		chaospeddler.NewMaestro(
			cloudControllerInfo.Username,
			cloudControllerInfo.Password,
			cloudControllerInfo.LoginURL,
			cloudControllerInfo.CCURL,
			sqlConn,
		),
	)
	chaos.Start()
	logger := lager.NewLogger("chaos-peddler-servicebroker")
	credentials := brokerapi.BrokerCredentials{
		Username: basicAuthInfo.Username,
		Password: basicAuthInfo.Password,
	}
	brokerAPI := brokerapi.New(chaos, logger, credentials)
	http.Handle("/", brokerAPI)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

//ExtractDBSQL ---
func ExtractDBSQL() (gormdb chaospeddler.GormDB) {
	var err error
	var db gorm.DB
	appEnv, _ := cfenv.Current()
	service, _ := appEnv.Services.WithName("sql-info")
	host := fmt.Sprintf("%v", service.Credentials["hostname"])
	port := fmt.Sprintf("%v", service.Credentials["port"])
	dbname := fmt.Sprintf("%v", service.Credentials["name"])
	user := fmt.Sprintf("%v", service.Credentials["username"])
	pass := fmt.Sprintf("%v", service.Credentials["password"])
	connectionString := user + ":" + pass + "@tcp(" + host + ":" + port + ")" + "/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	lo.G.Error("connection string: ", connectionString)

	if db, err = gorm.Open("mysql", connectionString); err == nil {
		db.DB()
		db.DB().Ping()
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		db.SingularTable(true)
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
			new(chaospeddler.ServiceInstance),
			new(chaospeddler.ServiceBinding),
		)

		gormdb = &chaospeddler.GormDBWrapper{
			DBWrapper: chaospeddler.DBWrapper{&db},
		}

	} else {
		lo.G.Error("there was an error connecting to mysql: ", err)
		panic(err)
	}
	return
}

//ExtractDBInfoMongo ----
func ExtractDBInfoMongo() (dbInfo DBInfo) {
	appEnv, _ := cfenv.Current()
	service, _ := appEnv.Services.WithName("db-info")
	dbInfo.ConnectionURL = fmt.Sprintf("%v", service.Credentials["uri"])
	dbInfo.DBName = fmt.Sprintf("%v", service.Credentials["database"])
	return
}

//ExtractBasicAuthInfo ----
func ExtractBasicAuthInfo() (basicAuthInfo BasicAuthInfo) {
	appEnv, _ := cfenv.Current()
	service, _ := appEnv.Services.WithName("basic-auth-info")
	basicAuthInfo.Username = fmt.Sprintf("%v", service.Credentials["username"])
	basicAuthInfo.Password = fmt.Sprintf("%v", service.Credentials["password"])
	return
}

//ExtractCloudControllerInfo ----
func ExtractCloudControllerInfo() (cloudControllerInfo CloudControllerInfo) {
	appEnv, _ := cfenv.Current()
	service, _ := appEnv.Services.WithName("cloud-controller-info")
	cloudControllerInfo.Username = fmt.Sprintf("%v", service.Credentials["username"])
	cloudControllerInfo.Password = fmt.Sprintf("%v", service.Credentials["password"])
	cloudControllerInfo.LoginURL = fmt.Sprintf("%v", service.Credentials["login-url"])
	cloudControllerInfo.CCURL = fmt.Sprintf("%v", service.Credentials["cc-url"])
	return
}

//CloudControllerInfo ----
type CloudControllerInfo struct {
	Username string
	Password string
	LoginURL string
	CCURL    string
}

//BasicAuthInfo ----
type BasicAuthInfo struct {
	Username string
	Password string
}

//DBInfo ----
type DBInfo struct {
	ConnectionURL string
	DBName        string
}
