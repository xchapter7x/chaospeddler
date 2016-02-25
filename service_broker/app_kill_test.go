package chaospeddler_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	. "github.com/xchapter7x/chaospeddler/service_broker"
)

var _ = Describe("Given a AppKill", func() {
	Describe("Given a KillPercent Method", func() {
		var (
			server  *ghttp.Server
			appKill *AppKill
		)

		BeforeEach(func() {
			server = NewTestServer(ghttp.NewTLSServer())
			appKill = NewAppKill("", "", server.URL(), server.URL())
		})

		AfterEach(func() {
			server.Close()
		})

		Context("when called with a service binding object representing a valid object", func() {
			var (
				err       error
				killRatio map[string]int
			)

			for _, control := range []map[string]int{{"killed": 1, "of": 4, "percent": 10}} {
				BeforeEach(func() {
					killRatio, err = appKill.KillPercent(ServiceBinding{AppGUID: "someguid"}, control["percent"])
				})
				It("then it should kill the given percentage of AIs from the overall AI count", func() {
					Ω(killRatio["killed"]).Should(Equal(control["killed"]))
					Ω(killRatio["of"]).Should(Equal(control["of"]))
				})
			}
		})
	})
})

func NewTestServer(server *ghttp.Server) *ghttp.Server {
	sampleSuccessTokenStringFormat := `{"access_token":"%s","token_type":"bearer","refresh_token":"%s","expires_in":599,"scope":"password.write cloud_controller.write openid cloud_controller.read","jti":"%s"}`
	loginTokenResponse := fmt.Sprintf(sampleSuccessTokenStringFormat, "access-token", "refresh-token", "jti")
	aiBasicResponse, _ := ioutil.ReadFile("fixtures/ai_basic_response.json")
	aiInfoHandler := ghttp.RespondWith(http.StatusOK, aiBasicResponse)
	aiDeleteHandler := ghttp.RespondWith(http.StatusNoContent, "")
	aiInfoPath, _ := regexp.Compile("/v2/apps/.*/instances")
	aiDeletePath, _ := regexp.Compile("/v2/apps/.*/instances/.*")
	server.RouteToHandler("GET", aiInfoPath, aiInfoHandler)
	server.RouteToHandler("DELETE", aiDeletePath, aiDeleteHandler)
	server.AppendHandlers(
		ghttp.RespondWith(http.StatusOK, loginTokenResponse),
	)
	return server
}
