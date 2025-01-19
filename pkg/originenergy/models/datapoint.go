package models

import "time"

type DataPoint struct {
	Type                string    `graphql:"type" json:"type"`
	StartDate           time.Time `graphql:"startDate" json:"startDate"`
	EndDate             time.Time `graphql:"endDate" json:"endDate"`
	TotalConsumedEnergy int       `graphql:"totalConsumedEnergy" json:"totalConsumedEnergy"`
	TotalCost           string    `graphql:"totalCost" json:"totalCost"`
	UsageCost           string    `graphql:"usageCost" json:"usageCost"`
	SupplyCost          string    `graphql:"supplyCost" json:"supplyCost"`
	TotalEarnings       string    `graphql:"totalEarnings" json:"totalEarnings"`
	TotalFeedInEnergy   string    `graphql:"totalFeedInEnergy" json:"totalFeedInEnergy"`
}
