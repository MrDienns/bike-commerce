package database

import "github.com/MrDienns/bike-commerce/pkg/api/model"

// Connector is a generic database interface.
type Connector interface {
	Connect() error
	Close() error

	UserFromCredentials(email, password string) (*model.User, error)

	GetCustomers() ([]*model.Customer, error)
	GetCustomer(id int) (*model.Customer, error)
	CreateCustomer(customer *model.Customer) error
	UpdateCustomer(customer *model.Customer) error
	DeleteCustomer(id int) error

	GetUsers() ([]*model.User, error)
	GetUser(id int) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error

	GetBikes() ([]*model.Bike, error)
	GetBike(id int) (*model.Bike, error)
	CreateBike(bike *model.Bike) error
	UpdateBike(bike *model.Bike) error
	DeleteBike(id int) error

	GetAccessories() ([]*model.Accessory, error)
	GetAccessory(id int) (*model.Accessory, error)
	CreateAccessory(bike *model.Accessory) error
	UpdateAccessory(bike *model.Accessory) error
	DeleteAccessory(id int) error

	GetRentals() ([]*model.Rental, error)
	GetRental(id int) (*model.Rental, error)
	CreateRental(bike *model.Rental) error
	UpdateRental(bike *model.Rental) error
	DeleteRental(id int) error
}
