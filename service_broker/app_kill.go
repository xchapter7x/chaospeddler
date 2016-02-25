package chaospeddler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/xchapter7x/lo"
)

//KillPercent - kills a given app by the given percentage
func (s *AppKill) KillPercent(binding ServiceBinding, percentKill int) (killRatio map[string]int, err error) {
	var aiList map[string]map[string]interface{}
	var killList []int
	aiList, err = s.getAIInfo(binding.AppGUID)
	aiKillCount := int(math.Ceil(float64(len(aiList)) / float64(10)))

	for i := 1; i <= aiKillCount; i++ {
		killList = append(killList, random(0, len(aiList)))
	}
	killRatio, err = s.killAIIndexes(binding.AppGUID, killList...)
	killRatio["of"] = len(aiList)
	return
}

func (s *AppKill) getAIInfo(appGUID string) (aiList map[string]map[string]interface{}, err error) {
	var req *http.Request
	var res *http.Response
	req, err = http.NewRequest("GET", s.CloudControllerAPIURL+"/v2/apps/"+appGUID+"/instances", nil)
	s.CloudController.AccessTokenDecorate(req)
	res, err = s.HttpClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &aiList)
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
