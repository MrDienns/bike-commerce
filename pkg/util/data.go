package util

import (
	"fmt"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
)

// BikesAsArray takes an array of bike structs and returns the bike names and types as flat array.
func BikesAsArray(bikes []*model.Bike) []string {
	ret := make([]string, len(bikes))
	for i, bike := range bikes {
		ret[i] = fmt.Sprintf("%v - %v", bike.Name, bike.Type)
	}
	return ret
}

// BikesAsArray takes an array of customer structs and returns the names and postal codes as flat array.
func CustomersAsArray(customers []*model.Customer) []string {
	ret := make([]string, len(customers))
	for i, customer := range customers {
		ret[i] = fmt.Sprintf("%v %v - %v", customer.Firstname, customer.Lastname, customer.Postalcode)
	}
	return ret
}

// UsersAsArray takes an array of user structs and returns the employee names in a flat array.
func UsersAsArray(users []*model.User) []string {
	ret := make([]string, len(users))
	for i, user := range users {
		ret[i] = user.Name
	}
	return ret
}
