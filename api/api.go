package api

import (
	"github.com/travelgateX/go-entities-client/entities"
	"github.com/travelgateX/go-entities-client/model"
)

// EntitiesAPI struct groups all functions in one place and run the request to the server.
// WARNING! this struct it's a helper, 'Entities API' is a Graphql server, requests can not
// be typed so if you need more personalized request use the client Query/Mutation main function.
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
///// Mutaion Functions helpers /////

// GrantAccessToGroup Entities API mutation function
func (q *EntitiesAPI) GrantAccessToGroup(id int, groups []string) (model.AdminMutation, error) {
	return q.client.NewMutation(grantAccessToGroupRQ(id, groups))
}
