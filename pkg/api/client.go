package api

import "github.com/MrDienns/bike-commerce/pkg/api/model"

type Client struct {
	token string
}

// NewClient accepts a token and returns a new *Client
func NewClient(token string) *Client {
	return &Client{token: token}
}

// CreateUser accepts a user and invokes the rest API to create it.
func (c *Client) CreateUser(user *model.User) error {
	return nil
}

// GetUsers invokes the rest API and returns all users.
func (c *Client) GetUsers() ([]*model.User, error) {
	return make([]*model.User, 0), nil
}

// GetUser invokes the rest API and loads a user based on the provided ID.
func (c *Client) GetUser(id string) (*model.User, error) {
	return &model.User{}, nil
}

// UpdateUser takes a user as parameter, invokes the rest API with it and patches the provided user with the new data.
func (c *Client) UpdateUser(user *model.User) error {
	return nil
}

// DeleteUser invokes the rest API to delete the passed user.
func (c *Client) DeleteUser(user *model.User) error {
	return nil
}
