package entities

import (
	"errors"

	"github.com/travelgateX/go-entities-client/model"
)

// Accesses Entities API query function
func (c *Client) Accesses() (model.AdminQuery, error) {
	return c.NewQuery(accessesRQ())
}

// AccessesByGroupCode Entities API query function
func (c *Client) AccessesByGroupCode(grCode string) (model.AdminQuery, error) {
	return c.NewQuery(accessesByGroupCodeRQ(grCode))
}

// Suppliers Entities API query function
func (c *Client) Suppliers() (model.AdminQuery, error) {
	return c.NewQuery(suppliersRQ())
}

// SuppliersByGroupCode Entities API query function
func (c *Client) SuppliersByGroupCode(grCode string) (model.AdminQuery, error) {
	return c.NewQuery(suppliersByGroupCodeRQ(grCode))
}

// ClientByGroupCode Entities API query function
func (c *Client) ClientByGroupCode(grCode string) (model.AdminQuery, error) {
	return c.NewQuery(clientByGroupCodeRQ(grCode))
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
