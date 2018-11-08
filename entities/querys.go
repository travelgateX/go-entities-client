package entities

import "github.com/travelgateX/go-entities-client/model"

// Accesses Entities API query function
func (c *Client) Accesses(id int) (model.AdminQuery, error) {
	return c.NewQuery(accessesRQ(id))
}
