// Package api ...
// The 'EntitiesApi' is a tool to connect to the "Entities API Graphql" endpoint.
// There are functions that it provides basic requests for common use.
//
// WARNING: 'EntitiesAPI' is a Graphql server, requests can not be typed so if you
// need more personalized request use the client Query/Mutation main function.
package api

import (
	"github.com/travelgateX/go-entities-client/entities"
	"github.com/travelgateX/go-entities-client/model"
)

// EntitiesAPI struct groups entities functions
type EntitiesAPI struct {
	client entities.Client
}

// NewEntitiesAPI constructor
func NewEntitiesAPI(bearer, endpoint string) EntitiesAPI {
	cli := entities.NewClient(bearer, endpoint)
	return EntitiesAPI{client: cli}
}

// NewDefaultEntitiesAPI constructor
func NewDefaultEntitiesAPI(bearer string) EntitiesAPI {
	cli := entities.NewDefaultClient(bearer)
	return EntitiesAPI{client: cli}
}

// DebugMode set debug mode
func (q *EntitiesAPI) DebugMode(debug bool) {
	q.client.DebugMode(debug)
}

// NewQuery excute new graphql query
func (q *EntitiesAPI) NewQuery(rq string) (model.AdminQuery, error) {
	return q.client.NewQuery(rq)
}

// NewMutation execute new graphql mutation
func (q *EntitiesAPI) NewMutation(rq string) (model.AdminMutation, error) {
	return q.client.NewMutation(rq)
}

//////////////////////////////////////////////////////////////////////////////////////////////
///// Query Functions helpers /////

// Accesses Entities API query function
func (q *EntitiesAPI) Accesses(id int) (model.AdminQuery, error) {
	return q.client.NewQuery(accessesRQ(id))
}

//////////////////////////////////////////////////////////////////////////////////////////////
///// Mutation Functions helpers /////

// GrantAccessToGroup Entities API mutation function
func (q *EntitiesAPI) GrantAccessToGroup(id int, groups []string) (model.AdminMutation, error) {
	return q.client.NewMutation(grantAccessToGroupRQ(id, groups))
}
