package chaospeddler

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

//Package constants
const (
	ChaosPeddlerServiceID    = "e6fcd6aa-047e-42be-bcfc-c8a2b426d707"
	CrazyChaosPlanID         = "f5b828db-dd86-4223-a190-628626c46585"
	AnnoyingChaosPlanID      = "f3c31f5d-e14f-4c69-a67a-9894a58fa417"
	MickeyMouseChaosPlanID   = "0bd1dee2-ea35-4088-bf11-019cd0947fff"
	NoChaosPlanID            = "10eeaa82-2193-4a59-84b2-0340cfdd9e43"
	KillPercentCrazy         = 20
	KillPercentAnnoying      = 10
	KillPercentMickeyMouse   = 5
	KillGroupPercentSelector = 5
	KillAIMaxPercentage      = 25
)

var random = func(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

var percentChanceOfTrue = func(i int) bool {
	mod := float64(100 / i)
	return math.Mod(float64(random(0, 100)), mod) == 0
}

//FindAllMatches - finds all matches for the given arguments
var FindAllMatches = func(db GormDB, planID, serviceID string) (serviceBindings []ServiceBinding, err error) {
	db.Raw("SELECT * FROM service_binding WHERE plan_id = ? and service_id = ? and deleted_at = '0000-00-00 00:00:00'", planID, serviceID).Scan(&serviceBindings)

	if len(serviceBindings) == 0 {
		err = errors.New(fmt.Sprintf("so plan/serviceid matches found for: PlanID: %s and ServiceID: %s ", planID, serviceID))
	}
	return
}
