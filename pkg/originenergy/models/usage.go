package models

type AccountUsageQueryVariables struct {
	AgreementID           string `json:"agreementId"`
	ElectricityRateFilter struct {
		ActiveEqual bool `json:"active_equal"`
	} `json:"electricityRateFilter"`
	EndDate   string `json:"endDate"`
	StartDate string `json:"startDate"`
	TimeUnit  string `json:"timeUnit"`
	Type      string `json:"type"`
}

type AccountUsageQueryRequest struct {
	Viewer struct {
		Kraken struct {
			Service struct {
				ElectricityService         ElectricityService         `graphql:"... on ElectricityService"`
				GasService                 GasService                 `graphql:"... on GasService"`
				EmbeddedHotWaterService    EmbeddedHotWaterService    `graphql:"... on EmbeddedHotWaterService"`
				EmbeddedElectricityService EmbeddedElectricityService `graphql:"... on EmbeddedElectricityService"`
			} `graphql:"service(type: $type, agreementId: $agreementId)"`
		} `graphql:"kraken"`
	} `graphql:"viewer"`
}

func (u *AccountUsageQueryRequest) OperationName() string {
	return "AccountUsageQuery"
}

type Usage struct {
	HasMissingCosts bool           `graphql:"hasMissingCosts" json:"hasMissingCosts"`
	Type            string         `graphql:"type" json:"type"`
	TimeUnit        string         `graphql:"timeUnit" json:"timeUnit"`
	HasFeedIn       bool           `graphql:"hasFeedIn" json:"hasFeedIn"`
	HasNext         bool           `graphql:"hasNext" json:"hasNext"`
	HasPrevious     bool           `graphql:"hasPrevious" json:"hasPrevious"`
	LastReadingDate string         `graphql:"lastReadingDate" json:"lastReadingDate"`
	UnitsOfMeasure  UnitsOfMeasure `graphql:"unitsOfMeasure"`
	DataPoints      []DataPoint    `graphql:"dataPoints"`
}

type UnitsOfMeasure struct {
	TotalCost           string `graphql:"totalCost" json:"totalCost"`
	UsageCost           string `graphql:"usageCost" json:"usageCost"`
	SupplyCost          string `graphql:"supplyCost" json:"supplyCost"`
	TotalConsumedEnergy string `graphql:"totalConsumedEnergy" json:"totalConsumedEnergy"`
	TotalEarnings       string `graphql:"totalEarnings" json:"totalEarnings"`
	TotalFeedInEnergy   string `graphql:"totalFeedInEnergy" json:"totalFeedInEnergy"`
}
