package kinde

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/axatol/kinde-go/pkg/kindeapi"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type ClientOptions struct {
	Domain       string
	Audience     string
	ClientID     string
	ClientSecret string
	Scopes       []string
}

func ClientOptionsFromEnv() *ClientOptions {
	return &ClientOptions{
		Domain:       os.Getenv("KINDE_DOMAIN"),
		Audience:     os.Getenv("KINDE_AUDIENCE"),
		ClientID:     os.Getenv("KINDE_CLIENT_ID"),
		ClientSecret: os.Getenv("KINDE_CLIENT_SECRET"),
		Scopes:       strings.Fields(os.Getenv("KINDE_SCOPES")),
	}
}

type Client struct {
	kindeapi.ClientWithResponsesInterface
	Client *http.Client
}

func New(ctx context.Context, options *ClientOptions) (*Client, error) {
	if options == nil {
		options = ClientOptionsFromEnv()
	}

	if options.Domain == "" ||
		options.Audience == "" ||
		options.ClientID == "" ||
		options.ClientSecret == "" {
		return nil, fmt.Errorf("missing required Kinde client options")
	}

	oauth2Config := clientcredentials.Config{
		ClientID:       options.ClientID,
		ClientSecret:   options.ClientSecret,
		TokenURL:       options.Domain + "/oauth/token",
		EndpointParams: url.Values{"audience": {options.Audience}},
		AuthStyle:      oauth2.AuthStyleInParams,
		Scopes:         options.Scopes,
	}

	authenticatedClient := oauth2Config.Client(ctx)

	apiClient, err := kindeapi.NewClientWithResponses(
		options.Domain,
		kindeapi.WithHTTPClient(oauth2Config.Client(ctx)),
	)

	if err != nil {
		return nil, err
	}

	client := &Client{
		Client:                       authenticatedClient,
		ClientWithResponsesInterface: apiClient,
	}

	return client, nil
}
