package chaospeddler

import "github.com/pivotal-cf/brokerapi"

//Services - returns out chaos peddler service meta data.
func (*ServiceBroker) Services() (services []brokerapi.Service) {
	services = append(services, brokerapi.Service{
		ID:            "e6fcd6aa-047e-42be-bcfc-c8a2b426d707",
		Name:          "Chaos-Peddler",
		Description:   "Use the chaos peddler and have AIs randomly killed",
		Bindable:      true,
		Tags:          []string{"chaos", "app", "testing"},
		PlanUpdatable: true,
		Plans: []brokerapi.ServicePlan{
			getCrazyChaosPlan(),
			getAnnoyingChaosPlan(),
			getMickeyMouseChaosPlan(),
			getNoChaosPlan(),
		},
		Metadata:        &brokerapi.ServiceMetadata{},
		DashboardClient: &brokerapi.ServiceDashboardClient{},
	})
	return
}

func getCrazyChaosPlan() brokerapi.ServicePlan {
	free := true
	return brokerapi.ServicePlan{
		ID:          "f5b828db-dd86-4223-a190-628626c46585",
		Name:        "Crazy-Chaos",
		Description: "high frequency of AI killing, w/ high percentage of total AI count",
		Free:        &free,
		Metadata:    &brokerapi.ServicePlanMetadata{},
	}
}

func getAnnoyingChaosPlan() brokerapi.ServicePlan {
	free := true
	return brokerapi.ServicePlan{
		ID:          "f3c31f5d-e14f-4c69-a67a-9894a58fa417",
		Name:        "Just-Annoying-Chaos",
		Description: "high frequency of AI killing, w/ low percentage of total AI count",
		Free:        &free,
		Metadata:    &brokerapi.ServicePlanMetadata{},
	}
}

func getMickeyMouseChaosPlan() brokerapi.ServicePlan {
	free := true
	return brokerapi.ServicePlan{
		ID:          "0bd1dee2-ea35-4088-bf11-019cd0947fff",
		Name:        "Mickey-Mouse-Chaos",
		Description: "lower frequency of AI killing, w/ lo percentage of total AI count ",
		Free:        &free,
		Metadata:    &brokerapi.ServicePlanMetadata{},
	}
}

func getNoChaosPlan() brokerapi.ServicePlan {
	free := false
	return brokerapi.ServicePlan{
		ID:          "10eeaa82-2193-4a59-84b2-0340cfdd9e43",
		Name:        "Where-Is-The-Chaos?",
		Description: "No chaos, no killing of AIs",
		Free:        &free,
		Metadata: &brokerapi.ServicePlanMetadata{
			Costs: []brokerapi.ServiceCost{
				brokerapi.ServiceCost{
					Unit: "future uncontrolled chaos costs you are not ready for",
					Amount: map[string]float64{
						"downtime/consulting/debugging": 10000000.00,
					},
				},
			},
		},
	}
}
