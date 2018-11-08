package entities

import (
	"context"
	"log"
	"os"

	"github.com/machinebox/graphql"
	"github.com/travelgateX/go-entities-client/model"
)

// Client : Grapqhql client
type Client struct {
	graphql *graphql.Client // graphql client
	bearer  string          // authentification bearer
	debug   bool            // log graphql pkg transaction
}

// NewClient constructor
func NewClient(bearer, endpoint string) Client {
	cli := graphql.NewClient(endpoint)
	return Client{graphql: cli, bearer: bearer}
}

// NewDefaultClient default constructor
func NewDefaultClient(bearer string) Client {
	cli := NewClient(bearer, EntityEndPointDev)
	if os.Getenv("DEPLOY_MODE") == "prod" || os.Getenv("DEPLOY_MODE") == "localProd" {
		cli = NewClient(bearer, EntityEndPointProd)
	}
	return cli
}

// DebugMode set debug mode
func (c *Client) DebugMode(debug bool) {
	if debug {
		c.graphql.Log = func(s string) { log.Println(s) }
	} else {
		c.graphql.Log = func(s string) {}
	}
}

// NewQuery excute new graphql query
func (c *Client) NewQuery(rq string) (model.AdminQuery, error) {
	ctx := context.Background()
	req := graphql.NewRequest(rq)
	req.Header.Add("Authorization", c.bearer)

	var res model.AdminQuery
	if err := c.graphql.Run(ctx, req, &res); err != nil {
		return model.AdminQuery{}, err
	}
	return res, nil
}

// NewMutation execute new graphql mutation
func (c *Client) NewMutation(rq string) (model.AdminMutation, error) {
	ctx := context.Background()
	req := graphql.NewRequest(rq)
	req.Header.Add("Authorization", c.bearer)

	var res model.AdminMutation
	if err := c.graphql.Run(ctx, req, &res); err != nil {
		return model.AdminMutation{}, err
	}
	return res, nil
}
