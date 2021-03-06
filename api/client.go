package api

import (
	"log"
	"os"

	"github.com/okta/okta-sdk-golang/okta"
)

// OktaClient is a wrapper around Okta Client struct
type OktaClient struct {
	*okta.Client
}

// NewOktaClient is initializing Okta config
func NewOktaClient() *OktaClient {
	myclient := &OktaClient{}

	if os.Getenv("OKTA_CLIENT_ORGURL") == "" {
		log.Fatal("Must set environment variable OKTA_CLIENT_ORGURL (ie: https://example.okta.com)")
	}
	if os.Getenv("OKTA_CLIENT_TOKEN") == "" {
		log.Fatal("Must set API token as environment variable OKTA_CLIENT_TOKEN")
	}

	oktaDomain := os.Getenv("OKTA_CLIENT_ORGURL")
	apiToken := os.Getenv("OKTA_CLIENT_TOKEN")

	config := okta.NewConfig().WithOrgUrl(oktaDomain).WithToken(apiToken)

	myclient.Client = okta.NewClient(config, nil, nil)
	return myclient
}
