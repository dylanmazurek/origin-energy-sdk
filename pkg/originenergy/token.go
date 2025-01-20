package originenergy

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dylanmazurek/origin-energy-sdk/pkg/originenergy/constants"
	"golang.org/x/oauth2"
)

var oauth2Config oauth2.Config = oauth2.Config{
	ClientID:     constants.AUTH_CLIENT_ID,
	ClientSecret: "",
	Endpoint: oauth2.Endpoint{
		AuthURL:   fmt.Sprintf("https://%s/authorize", constants.AUTH_DOMAIN),
		TokenURL:  fmt.Sprintf("https://%s/oauth/token", constants.AUTH_DOMAIN),
		AuthStyle: oauth2.AuthStyleInParams,
	},
	RedirectURL: constants.AUTH_CALLBACK_URL,
	Scopes:      []string{"openid", "email", "all", "offline_access"},
}

func (c *AuthClient) LoadToken() error {
	var sessionFile string = "session.json"
	sesisonFileEnv, ok := os.LookupEnv("SESSION_FILE")
	if ok {
		log.Printf("using session file: %s\n", sesisonFileEnv)
		sessionFile = sesisonFileEnv
	}

	_, err := os.Stat(sessionFile)
	if os.IsNotExist(err) {
		return ErrSessionFileNotFound
	}

	session, err := os.OpenFile(sessionFile, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}

	token := &oauth2.Token{}
	err = json.NewDecoder(session).Decode(token)
	if err != nil {
		return err
	}

	c.token = token

	return nil
}

func (c *AuthClient) SaveToken() error {
	var sessionFile string = "session.json"
	sesisonFileEnv, ok := os.LookupEnv("SESSION_FILE")
	if ok {
		sessionFile = sesisonFileEnv
	}

	session, err := os.OpenFile(sessionFile, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	err = json.NewEncoder(session).Encode(c.token)
	if err != nil {
		return err
	}

	return nil
}

func (c *AuthClient) NewToken(ctx context.Context) error {
	verifier := oauth2.GenerateVerifier()

	opts := []oauth2.AuthCodeOption{
		oauth2.AccessTypeOffline,
		oauth2.S256ChallengeOption(verifier),
		oauth2.SetAuthURLParam("audience", "https://digitalapi"),
	}

	url := oauth2Config.AuthCodeURL("state", opts...)
	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)

	fmt.Print("Enter the code: ")
	var code string
	_, err := fmt.Scan(&code)
	if err != nil {
		return err
	}

	token, err := oauth2Config.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		return err
	}

	c.token = token

	return nil
}
