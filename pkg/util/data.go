package util

import (
	"fmt"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
)

func BikesAsArray(bikes []*model.Bike) []string {
	ret := make([]string, len(bikes))
	for i, bike := range bikes {
		ret[i] = fmt.Sprintf("%v - %v", bike.Name, bike.Type)
	}
	return ret
}

func CustomersAsArray(customers []*model.Customer) []string {
	ret := make([]string, len(customers))
	for i, customer := range customers {
		ret[i] = fmt.Sprintf("%v %v - %v", customer.Firstname, customer.Lastname, customer.Postalcode)
	}
	return ret
}

func UsersAsArray(users []*model.User) []string {
	ret := make([]string, len(users))
	for i, user := range users {
		ret[i] = user.Name
	}
	return ret
}
