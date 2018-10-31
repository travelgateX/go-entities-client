package entities

import (
	"context"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

// Client : Grapqhql client
type Client struct {
	graphql *graphql.Client // graphql client
	bearer  string          // authentification bearer
}

// NewClient constructor
func NewClient(bearer string, debug bool) Client {
	cli := graphql.NewClient(EntityEndPointDev)
	if os.Getenv("DEPLOY_MODE") == "prod" || os.Getenv("DEPLOY_MODE") == "localProd" {
		cli = graphql.NewClient(EntityEndPointProd)
	}
	if debug {
		cli.Log = func(s string) { log.Println(s) }
	}
	return Client{graphql: cli, bearer: bearer}
}

// NewRequest creates new graphql request
func (c *Client) NewRequest(request string) struct{} {
	ctx := context.Background()
	req := graphql.NewRequest(request)
	req.Header.Add("Authorization", c.bearer)

	var res struct{}
	if err := c.graphql.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}
	return res
}
