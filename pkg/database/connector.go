package database

import "github.com/MrDienns/bike-commerce/pkg/api/model"

// Connector is a generic database interface.
type Connector interface {
	Connect() error
	Close() error

	UserFromCredentials(email, password string) *model.User

	GetCustomer(id int) (*model.Customer, error)
	CreateCustomer(customer *model.Customer) error
	UpdateCustomer(customer *model.Customer) error
	DeleteCustomer(id int) error
}
