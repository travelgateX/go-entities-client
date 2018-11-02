package entities

import (
	"context"
	"log"
	"os"

	"github.com/machinebox/graphql"
	"github.com/travelgateX/go-entities-client/domain"
)

// DEBUG : log graphql pkg transaction
const DEBUG = false

// Client : Grapqhql client
type Client struct {
	graphql *graphql.Client // graphql client
	bearer  string          // authentification bearer
}

// NewClient constructor
func NewClient(bearer, endpoint string, debug bool) Client {
	cli := graphql.NewClient(endpoint)
	if debug {
		cli.Log = func(s string) { log.Println(s) }
	}
	return Client{graphql: cli, bearer: bearer}
}

// NewDefaultClient default constructor
func NewDefaultClient(bearer string) Client {
	cli := NewClient(bearer, EntityEndPointDev, DEBUG)
	if os.Getenv("DEPLOY_MODE") == "prod" || os.Getenv("DEPLOY_MODE") == "localProd" {
		cli = NewClient(bearer, EntityEndPointProd, DEBUG)
	}
	return cli
}

// NewRequest creates new graphql request
func (c *Client) NewRequest(request string) domain.Data {
	ctx := context.Background()
	req := graphql.NewRequest(request)
	req.Header.Add("Authorization", c.bearer)

	var res domain.Data
	if err := c.graphql.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}
	return res
}
