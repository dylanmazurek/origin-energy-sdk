package originenergy

import (
	"context"
	"os"
	"time"

	"github.com/dylanmazurek/origin-energy-sdk/pkg/originenergy/constants"
	"github.com/dylanmazurek/origin-energy-sdk/pkg/originenergy/models"
	"github.com/hasura/go-graphql-client"
)

// AccountUsageFilter is a filter for the GetAccountUsage method.
// [AgreementId] is required. StartDate defaults to 1 day ago. EndDate defaults to now. TimeUnit defaults to Hourly.
type AccountUsageFilter struct {
	StartDate time.Time
	EndDate   time.Time
	Type      constants.ServiceType
	TimeUnit  constants.UsageTimeUnit
}

func (f *AccountUsageFilter) GetVariables() (map[string]any, error) {
	now := time.Now()
	defaultVariables := map[string]any{
		"timeUnit":  constants.Hourly,
		"startDate": constants.DateTime(now.AddDate(0, 0, -1).Format(time.RFC3339)),
		"endDate":   constants.DateTime(now.Format(time.RFC3339)),
	}

	agreementId, agreementIdSet := os.LookupEnv("AGREEMENT_ID")
	if !agreementIdSet {
		return nil, ErrAgreementIDNotSet
	}

	variables := map[string]any{
		"agreementId": agreementId,
		"type":        f.Type,
		"electricityRateFilter": models.ElectricityRateFilter{
			ActiveEqual: true,
		},
	}

	for key, value := range defaultVariables {
		if _, ok := variables[key]; !ok {
			variables[key] = value
		}
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
