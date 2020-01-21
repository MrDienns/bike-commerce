package model

// Bike is a struct which represents a bike data transfer object.
type Bike struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Price        float32 `json:"price"`
	Quantity     int     `json:"quantity"`
	AmountRented int     `json:"amount_rented"`
}
