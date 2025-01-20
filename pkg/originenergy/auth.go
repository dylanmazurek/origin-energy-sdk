package originenergy

import (
	"context"
	"fmt"

	"github.com/dylanmazurek/origin-energy-sdk/pkg/originenergy/constants"
	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
)

type AuthClient struct {
	internalClient *graphql.Client

	token *oauth2.Token
}

func NewAuthClient(ctx context.Context) (*AuthClient, error) {
	authClient := &AuthClient{}

	err := authClient.LoadToken()
	if err != nil {
		fmt.Printf("error loading token: %s\n", err)
		err := authClient.NewToken(ctx)
		if err != nil {
			return nil, err
		}

		err = authClient.SaveToken()
		if err != nil {
			return nil, err
		}
	}

	authClientTransport := oauth2Config.Client(ctx, authClient.token)

	graphQlUrl := fmt.Sprintf("%s%s", constants.API_BASE_URL, constants.GRAPHQL_PATH)
	newGraphQLClient := graphql.NewClient(graphQlUrl, authClientTransport)
	authClient.internalClient = newGraphQLClient

	return authClient, nil
}
