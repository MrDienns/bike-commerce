package api

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MrDienns/bike-commerce/pkg/util"

	"github.com/MrDienns/bike-commerce/pkg/api/response"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
)

// Client is the rest API client.
type Client struct {
	Token    string
	User     *model.User
	key      *rsa.PublicKey
	endpoint string
}

// NewClient accepts a token and returns a new *Client
func NewClient(key *rsa.PublicKey, endpoint string) *Client {
	return &Client{key: key, endpoint: endpoint}
}

func (c *Client) Authenticate(email, password string) (*model.User, string, error) {

	authRequest := &model.AuthRequest{Email: email, Password: password}
	var authResponse response.AuthResponse
	err := c.invoke("/api/authenticate", http.MethodPost, authRequest, &authResponse)
	if err != nil {
		return nil, "", err
	}

	user, err := util.UserFromToken(c.key, authResponse.Token)
	if err != nil {
		return nil, "", err
	}

	c.Token = authResponse.Token
	c.User = user

	return user, authResponse.Token, nil
}

// CreateUser accepts a user and invokes the rest API to create it.
func (c *Client) CreateUser(user *model.User) error {
	return c.invokeEmpty("/api/user", http.MethodPost, user)
}

// GetUsers invokes the rest API and returns all users.
func (c *Client) GetUsers() ([]*model.User, error) {

	var resp []*model.User
	err := c.invoke("/api/user", http.MethodGet, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetUser invokes the rest API and loads a user based on the provided ID.
func (c *Client) GetUser(id string) (*model.User, error) {
	return &model.User{}, nil
}

// UpdateUser takes a user as parameter, invokes the rest API with it and patches the provided user with the new data.
func (c *Client) UpdateUser(user *model.User) error {
	return c.invokeEmpty(fmt.Sprintf("/api/user/%v", user.Id), http.MethodPut, user)
}

// DeleteUser invokes the rest API to delete the passed user.
func (c *Client) DeleteUser(user *model.User) error {
	return c.invokeEmpty(fmt.Sprintf("/api/user/%v", user.Id), http.MethodDelete, nil)
}

// CreateBike accepts a user and invokes the rest API to create it.
func (c *Client) CreateBike(bike *model.Bike) error {
	return c.invokeEmpty("/api/bike", http.MethodPost, bike)
}

// GetBikes invokes the rest API and returns all bikes.
func (c *Client) GetBikes() ([]*model.Bike, error) {
	var resp []*model.Bike
	err := c.invoke("/api/bike", http.MethodGet, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetBike invokes the rest API with the passed bike ID and returns the bike.
func (c *Client) GetBike(id string) (*model.Bike, error) {
	return &model.Bike{}, nil
}

// UpdateBike takes a bike as argument and updates it by invoking the rest API.
func (c *Client) UpdateBike(bike *model.Bike) error {
	return c.invokeEmpty(fmt.Sprintf("/api/bike/%v", bike.ID), http.MethodPut, bike)
}

// DeleteBike takes a bike as argument and deletes it by invoking the rest API.
func (c *Client) DeleteBike(bike *model.Bike) error {
	return c.invokeEmpty(fmt.Sprintf("/api/bike/%v", bike.ID), http.MethodDelete, nil)
}

// CreateCustomer accepts a customer and invokes the rest API to create it.
func (c *Client) CreateCustomer(customer *model.Customer) error {
	return c.invokeEmpty("/api/customer", http.MethodPost, customer)
}

// GetCustomers invokes the rest API and returns all customers.
func (c *Client) GetCustomers() ([]*model.Customer, error) {

	var resp []*model.Customer
	err := c.invoke("/api/customer", http.MethodGet, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetCustomer invokes the rest API and loads a customer based on the provided ID.
func (c *Client) GetCustomer(id string) (*model.Customer, error) {
	return &model.Customer{}, nil
}

// UpdateCustomer takes a customer as parameter, invokes the rest API with it and patches the provided customer with
// the new data.
func (c *Client) UpdateCustomer(customer *model.Customer) error {
	return c.invokeEmpty(fmt.Sprintf("/api/customer/%v", customer.ID), http.MethodPut, customer)
}

// DeleteCustomer invokes the rest API to delete the passed customer.
func (c *Client) DeleteCustomer(customer *model.Customer) error {
	return c.invokeEmpty(fmt.Sprintf("/api/customer/%v", customer.ID), http.MethodDelete, nil)
}

// CreateAccessory accepts a customer and invokes the rest API to create it.
func (c *Client) CreateAccessory(accessory *model.Accessory) error {
	return c.invokeEmpty("/api/accessory", http.MethodPost, accessory)
}

// GetAccessories invokes the rest API and returns all accessories.
func (c *Client) GetAccessories() ([]*model.Accessory, error) {

	var resp []*model.Accessory
	err := c.invoke("/api/accessory", http.MethodGet, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAccessory invokes the rest API and loads a accessory based on the provided ID.
func (c *Client) GetAccessory(id string) (*model.Accessory, error) {
	return &model.Accessory{}, nil
}

// UpdateAccessory takes a customer as parameter, invokes the rest API with it and patches the provided accessory with
// the new data.
func (c *Client) UpdateAccessory(accessory *model.Accessory) error {
	return c.invokeEmpty(fmt.Sprintf("/api/accessory/%v", accessory.ID), http.MethodPut, accessory)
}

// DeleteAccessory invokes the rest API to delete the passed accessory.
func (c *Client) DeleteAccessory(accessory *model.Accessory) error {
	return c.invokeEmpty(fmt.Sprintf("/api/accessory/%v", accessory.ID), http.MethodDelete, nil)
}

// CreateRental accepts a rental and invokes the rest API to create it.
func (c *Client) CreateRental(rental *model.Rental) error {
	return c.invokeEmpty("/api/rental", http.MethodPost, rental)
}

// GetRentals invokes the rest API and returns all rentals.
func (c *Client) GetRentals() ([]*model.Rental, error) {

	var resp []*model.Rental
	err := c.invoke("/api/rental", http.MethodGet, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetRental invokes the rest API and loads a rental based on the provided ID.
func (c *Client) GetRental(id string) (*model.Rental, error) {
	return &model.Rental{}, nil
}

// UpdateRental takes a rental as parameter, invokes the rest API with it and patches the provided rental with
// the new data.
func (c *Client) UpdateRental(rental *model.Rental) error {
	return c.invokeEmpty(fmt.Sprintf("/api/rental/%v", rental.ID), http.MethodPut, rental)
}

// DeleteRental invokes the rest API to delete the passed rental.
func (c *Client) DeleteRental(rental *model.Rental) error {
	return c.invokeEmpty(fmt.Sprintf("/api/rental/%v", rental.ID), http.MethodDelete, nil)
}

// invoke accepts a URL, method, body and response object pointer to send and handle the request.
func (c *Client) invoke(url, method string, body, responseObj interface{}) error {
	client := &http.Client{}
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(b)
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.endpoint, url), buffer)
	request.Header.Add("Content-Type", "application/json")
	if c.User != nil {
		request.Header.Add("Authorization", "Bearer "+c.Token)
	}
	if err != nil {
		return err
	}
	httpResponse, err := client.Do(request)
	if err != nil {
		return err
	}
	if httpResponse.StatusCode != 200 {
		var errorResponse response.Error
		err = unmarshalResponse(httpResponse, &errorResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf(errorResponse.Message)
	}

	return unmarshalResponse(httpResponse, responseObj)
}

// invokeEmpty takes a URL, method and body to send and handle an HTTP request.
func (c *Client) invokeEmpty(url, method string, body interface{}) error {
	client := &http.Client{}
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(b)
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.endpoint, url), buffer)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	if c.User != nil {
		request.Header.Add("Authorization", "Bearer "+c.Token)
	}
	httpResponse, err := client.Do(request)
	if err != nil {
		return err
	}
	if httpResponse.StatusCode != 200 {
		var errorResponse response.Error
		err = unmarshalResponse(httpResponse, &errorResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf(errorResponse.Message)
	}
	return nil
}

// unmarshalResponse takes an http response, and a pointer to a response object.
func unmarshalResponse(httpResp *http.Response, response interface{}) error {
	b, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, response)
}
