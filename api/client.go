package api

import (
	"context"
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

	// config := okta.NewConfig().WithOrgUrl(oktaDomain).WithToken(apiToken)
	// d := time.Now().Add(50 * time.Millisecond)
	// ctx, _ := context.WithDeadline(context.Background(), d)
	config, _ := okta.NewClient(context.TODO(), okta.WithOrgUrl(oktaDomain), okta.WithToken(apiToken))

	// client := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

	// myclient.Client, _ = okta.NewClient(config)
	// return myclient
	myclient.Client = config
	// myclient.WithOrgUrl = oktaDomain
	// c := OktaClient{config}
	// d := new(OktaClient)
	// println(myclient)
	// println(config)
	// myclient.Client =
	// return config
	return myclient

}

// c := &Client{}
// 	c.config = config
// 	c.requestExecutor = NewRequestExecutor(&config.HttpClient, oktaCache, config)

// 	c.resource.client = c
