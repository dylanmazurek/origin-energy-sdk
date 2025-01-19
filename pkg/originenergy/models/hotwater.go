package models

type EmbeddedHotWaterService struct {
	Usage Usage `graphql:"usage(timeUnit: $timeUnit, startDate: $startDate, endDate: $endDate, orderBy: START_DATE_ASC)"`
}
