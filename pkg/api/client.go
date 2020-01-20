package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MrDienns/bike-commerce/pkg/api/response"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
)

// Client is the rest API client.
type Client struct {
	token    string
	endpoint string
}

// NewClient accepts a token and returns a new *Client
func NewClient(token, endpoint string) *Client {
	return &Client{token: token, endpoint: endpoint}
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

// GetBikes invokes the rest API and returns all bikes.
func (c *Client) GetBikes() ([]*model.Bike, error) {
	return make([]*model.Bike, 0), nil
}

// GetBike invokes the rest API with the passed bike ID and returns the bike.
func (c *Client) GetBike(id string) (*model.Bike, error) {
	return &model.Bike{}, nil
}

// UpdateBike takes a bike as argument and updates it by invoking the rest API.
func (c *Client) UpdateBike(bike *model.Bike) error {
	return nil
}

// DeleteBike takes a bike as argument and deletes it by invoking the rest API.
func (c *Client) DeleteBike(bike *model.Bike) error {
	return nil
}

func (c *Client) invoke(url, method string, responseObj *interface{}) error {
	client := &http.Client{}
	var reader io.Reader
	request, err := http.NewRequest(fmt.Sprintf("%s%s", c.endpoint, url), method, reader)
	if err != nil {
		return err
	}
	httpResponse, err := client.Do(request)
	if err != nil {
		return err
	}
	if httpResponse.StatusCode != 200 {
		var errorResponse response.Error
		err = unmarshalResponse(httpResponse.Body, &errorResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf(errorResponse.Message)
	}

	return unmarshalResponse(httpResponse.Body, responseObj)
}

func (c *Client) invokeEmpty(url, method string) error {
	client := &http.Client{}
	var reader io.Reader
	request, err := http.NewRequest(fmt.Sprintf("%s%s", c.endpoint, url), method, reader)
	if err != nil {
		return err
	}
	httpResponse, err := client.Do(request)
	if err != nil {
		return err
	}
	if httpResponse.StatusCode != 200 {
		var errorResponse response.Error
		err = unmarshalResponse(httpResponse.Body, &errorResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf(errorResponse.Message)
	}
	return nil
}

func unmarshalResponse(reader io.Reader, response interface{}) error {
	var bytes []byte
	_, err := reader.Read(bytes)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, response)
}
