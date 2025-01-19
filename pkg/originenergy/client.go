package originenergy

import (
	"context"
)

type Client struct {
	authClient *AuthClient
}

func New(ctx context.Context) (*Client, error) {
	authClient, err := NewAuthClient(ctx)
	if err != nil {
		return nil, err
	}

	newServiceClient := &Client{
		authClient: authClient,
	}

	return newServiceClient, nil
}
