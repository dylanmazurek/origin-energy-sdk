package models

type ElectricityService struct {
	Product ElectricityProduct `graphql:"product"`
	Usage   Usage              `graphql:"usage(timeUnit: $timeUnit, startDate: $startDate, endDate: $endDate, orderBy: START_DATE_ASC)"`
}

type EmbeddedElectricityService struct {
	Usage Usage `graphql:"usage(timeUnit: $timeUnit, startDate: $startDate, endDate: $endDate, orderBy: START_DATE_ASC)"`
}

type ElectricityProduct struct {
	Rates struct {
		ElectricityRatesFilter struct {
			StepStart int `graphql:"stepStart"`
			StepEnd   int `graphql:"stepEnd"`
		} `graphql:"... on ElectricityRate"`
	} `graphql:"rates(filter: $electricityRateFilter)"`
}

type ElectricityRateFilter struct {
	ActiveEqual bool `json:"active_equal"`
}
