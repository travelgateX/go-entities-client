package entities

import "github.com/travelgateX/go-entities-client/model"

// GrantAccessToGroup Entities API mutation function
func (c *Client) GrantAccessToGroup(id int, groups []string) (model.AdminMutation, error) {
	return c.NewMutation(grantAccessToGroupRQ(id, groups))
}
