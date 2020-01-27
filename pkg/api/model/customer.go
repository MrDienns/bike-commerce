package model

// Customer struct is a struct which is used to represent a customer.
type Customer struct {
	ID                  int    `json:"id"`
	Firstname           string `json:"firstname"`
	Lastname            string `json:"lastname"`
	Postalcode          string `json:"postalcode"`
	Housenumber         int    `json:"housenumber"`
	HousenumberAddition string `json:"housenumber_addition"`
	Comment             string `json:"comment"`
}
