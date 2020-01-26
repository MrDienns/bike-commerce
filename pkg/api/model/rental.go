package model

type Rental struct {
	ID          int                      `json:"id"`
	StartDate   string                   `json:"start_date"`
	Days        int                      `json:"days"`
	TotalPrice  float32                  `json:"total_price"`
	Customer    *Customer                `json:"customer"`
	Employee    *User                    `json:"employee"`
	Bike        *Bike                    `json:"bike"`
	Accessories map[int]*RentedAccessory `json:"accessories"`
}
