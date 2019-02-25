package entities

import (
	"errors"

	"github.com/travelgateX/go-entities-client/model"
)

// CreateAccess Entities API query function
func (c *Client) CreateAccess() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// UpdateAccess Entities API query function
func (c *Client) UpdateAccess() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// GrantAccessToGroup Entities API mutation function
func (c *Client) GrantAccessToGroup(id int, groups []string) (model.AdminMutation, error) {
	return c.NewMutation(grantAccessToGroupRQ(id, groups))
}

// DeleteAccessFromGroup Entities API mutation function
func (c *Client) DeleteAccessFromGroup(id int, groups []string) (model.AdminMutation, error) {
	return c.NewMutation(deleteAccessFromGroupRQ(id, groups))
}

// GrantSupplierToGroup Entities API mutation function
func (c *Client) GrantSupplierToGroup(id string, groups []string) (model.AdminMutation, error) {
	if id == "" {
		return model.AdminMutation{}, errors.New("supplier id can't be blank")
	}
	return c.NewMutation(grantSupplierToGroupRQ(id, groups))
}

// DeleteSupplierFromGroup Entities API mutation function
func (c *Client) DeleteSupplierFromGroup(id string, groups []string) (model.AdminMutation, error) {
	return c.NewMutation(deleteSupplierFromGroupRQ(id, groups))
}

// GrantClientToGroup Entities API mutation function
func (c *Client) GrantClientToGroup() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// DeleteClientFromGroup Entities API mutation function
func (c *Client) DeleteClientFromGroup() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// CreateDefaultClient Entities API mutation function
func (c *Client) CreateDefaultClient(input model.CreateClientInput) (model.AdminMutation, error) {
	return c.NewMutation(createDefaultClient(input))
}

// CreateClient Entities API mutation function
func (c *Client) CreateClient() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// UpdateClient Entities API mutation function
func (c *Client) UpdateClient() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// CreateProfile Entities API mutation function
func (c *Client) CreateProfile() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// UpdateProfile Entities API mutation function
func (c *Client) UpdateProfile() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// AddEntitiesToProfile Entities API mutation function
func (c *Client) AddEntitiesToProfile() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}

// RemoveEntitiesFromProfile Entities API mutation function
func (c *Client) RemoveEntitiesFromProfile() (model.AdminMutation, error) {
	return model.AdminMutation{}, errors.New("not implemented")
}
