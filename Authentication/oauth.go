package authentication

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type ClientCredentials struct {
	GrantType    string
	ClientId     string
	ClientSecret string
	ApiUrl       string
}

type AuthResponse struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

func (x ClientCredentials) Login() *http.Client {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     x.ClientId,
		ClientSecret: x.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL: x.ApiUrl,
		},
	}

	verifier := oauth2.GenerateVerifier()

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	tok, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	return client
}
