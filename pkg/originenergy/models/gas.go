package models

type GasService struct {
	Usage Usage `graphql:"usage(timeUnit: $timeUnit, startDate: $startDate, endDate: $endDate, orderBy: START_DATE_ASC)"`
}

type EmbeddedGasService struct {
	Usage Usage `graphql:"usage(timeUnit: $timeUnit, startDate: $startDate, endDate: $endDate, orderBy: START_DATE_ASC)"`
}
