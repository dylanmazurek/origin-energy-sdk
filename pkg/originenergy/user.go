package originenergy

import (
	"context"

	"github.com/dylanmazurek/origin-energy-sdk/pkg/originenergy/models"
	"github.com/hasura/go-graphql-client"
)

func (c *Client) GetUserAccount(ctx context.Context) (*models.UserAccountsRequest, error) {
	var userAccounts models.UserAccountsRequest

	operationName := graphql.OperationName("UserAccounts")
	err := c.authClient.internalClient.Query(ctx, &userAccounts, nil, operationName)
	if err != nil {
		return nil, err
	}

	return &userAccounts, nil
}
