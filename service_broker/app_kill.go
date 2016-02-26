package chaospeddler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/xchapter7x/lo"
)

//KillPercent - kills a given app by the given percentage
func (s *AppKill) KillPercent(binding ServiceBinding, percentKill int) (killRatio map[string]int, err error) {
	var aiList map[string]map[string]interface{}
	var killList []int
	aiList, err = s.getAIInfo(binding.AppGUID)

	for i, v := range aiList {

		if percentChanceOfTrue(percentKill) {
			lo.G.Debug("selected index: ", i, v)
			idx, _ := strconv.Atoi(i)
			killList = append(killList, idx)
		}

		if len(killList)/len(aiList) >= KillAIMaxPercentage/100 {
			break
		}
	}
	killRatio, err = s.killAIIndexes(binding.AppGUID, killList...)
	killRatio["of"] = len(aiList)
	return
}

func (s *AppKill) getAIInfo(appGUID string) (aiList map[string]map[string]interface{}, err error) {
	var req *http.Request
	var res *http.Response
	req, err = http.NewRequest("GET", s.CloudControllerAPIURL+"/v2/apps/"+appGUID+"/instances", nil)

	if _, err = s.CloudController.Login(); err == nil {
		s.CloudController.AccessTokenDecorate(req)

		if res, err = s.HTTPClient.Do(req); err == nil {
			var body []byte

			if body, err = ioutil.ReadAll(res.Body); err == nil {
				err = json.Unmarshal(body, &aiList)
			}
		}
	}
	return
}

func (s *AppKill) killAIIndexes(guid string, indexes ...int) (killRatio map[string]int, err error) {
	killRatio = make(map[string]int)
	killRatio["killed"] = 0

	for _, idx := range indexes {
		if _, err = http.NewRequest("DELETE", fmt.Sprintf("%s/v2/apps/%s/instances/%d", s.CloudControllerAPIURL, guid, idx), nil); err != nil {
			lo.G.Error("there was an error while deleting guid:"+guid+" idx:"+string(idx)+" url:"+s.CloudControllerAPIURL, err)
		} else {
			killRatio["killed"]++
		}
	}
	return
}
