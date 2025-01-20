package originenergy

import (
	"context"
	"time"

	"github.com/dylanmazurek/origin-energy-sdk/pkg/originenergy/constants"
	"github.com/dylanmazurek/origin-energy-sdk/pkg/originenergy/models"
	"github.com/hasura/go-graphql-client"
)

// AccountUsageFilter is a filter for the GetAccountUsage method.
type AccountUsageFilter struct {
	// AgreementID is the agreement ID to query. (Required)
	AgreementID string
	// StartDate is the start date of the usage data. (Optional - defaults to 1 day ago)
	StartDate time.Time
	// EndDate is the end date of the usage data. (Optional - defaults to now)
	EndDate time.Time
	// Type is the service type to query. (Required)
	Type constants.ServiceType
	// TimeUnit is the time unit to query. (Optional - defaults to hourly)
	TimeUnit constants.UsageTimeUnit
}

func (f *AccountUsageFilter) GetVariables() (map[string]any, error) {
	var startDateStr, endDateStr constants.DateTime
	startDateStr.FromTime(f.StartDate)
	endDateStr.FromTime(f.EndDate)

	variables := map[string]any{
		"agreementId": f.AgreementID,
		"startDate":   startDateStr,
		"endDate":     endDateStr,
		"type":        f.Type,
		"timeUnit":    f.TimeUnit,
		"electricityRateFilter": models.ElectricityRateFilter{
			ActiveEqual: true,
		},
	}
	return variables, nil
}

func (c *Client) GetAccountUsage(ctx context.Context, filter AccountUsageFilter) (*models.AccountUsageQueryRequest, error) {
	variables, err := filter.GetVariables()
	if err != nil {
		return nil, err
	}

	var accountUsage models.AccountUsageQueryRequest
	operationName := graphql.OperationName("AccountUsageQuery")
	err = c.authClient.internalClient.Query(ctx, &accountUsage, variables, operationName)
	if err != nil {
		return nil, err
	}

	return &accountUsage, nil
}
