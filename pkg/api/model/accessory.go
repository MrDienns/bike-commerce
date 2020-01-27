package model

// Accessory struct represents an accessory from the catalogue
type Accessory struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

// RentedAccessory struct represents a rented accessory, which includes an amount.
type RentedAccessory struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Amount int     `json:"amount"`
}
