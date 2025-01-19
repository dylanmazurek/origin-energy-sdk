package originenergy

import (
	"context"
	"errors"
	"fmt"
	"net/http"

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
		err := authClient.NewToken(ctx)
		if err != nil {
			return nil, err
		}

		err = authClient.SaveToken()
		if err != nil {
			return nil, err
		}
	}

	authClientTransport, err := authClient.createAuthTransport()
	if err != nil {
		return nil, err
	}

	authClient.internalClient = authClientTransport

	return authClient, nil
}

func (c *AuthClient) createAuthTransport() (*graphql.Client, error) {
	authClient := &http.Client{
		Transport: &addAuthHeaderTransport{
			T:     http.DefaultTransport,
			Token: c.token,
		},
	}

	graphQlUrl := fmt.Sprintf("%s%s", constants.API_BASE_URL, constants.GRAPHQL_PATH)
	newGraphQLClient := graphql.NewClient(graphQlUrl, authClient)

	return newGraphQLClient, nil
}

type addAuthHeaderTransport struct {
	T     http.RoundTripper
	Token *oauth2.Token
}

func (adt *addAuthHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if adt.Token == nil {
		return nil, errors.New("no token")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", adt.Token.AccessToken))

	return adt.T.RoundTrip(req)
}
