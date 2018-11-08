package entities

import (
	"errors"

	"github.com/travelgateX/go-entities-client/model"
)

// Accesses Entities API query function
func (c *Client) Accesses(id int) (model.AdminQuery, error) {
	return c.NewQuery(accessesRQ(id))
}

// Suppliers Entities API query function
func (c *Client) Suppliers() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}

// Clients Entities API query function
func (c *Client) Clients() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}

// ServiceAPI Entities API query function
func (c *Client) ServiceAPI() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}

// PointsOfSale Entities API query function
func (c *Client) PointsOfSale() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}

// Profiles Entities API query function
func (c *Client) Profiles() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}

// Entities Entities API query function
func (c *Client) Entities() (model.AdminQuery, error) {
	return model.AdminQuery{}, errors.New("not implemented")
}
