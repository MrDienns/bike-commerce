package view

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/MrDienns/bike-commerce/pkg/util"
	"github.com/rivo/tview"
)

type rentalNewView struct {
	*root
	tview.Primitive
	rental *model.Rental
}

func NewRentalNewView(r *root) *rentalNewView {
	ret := &rentalNewView{root: r, rental: &model.Rental{
		Employee:    r.client.User,
		Bike:        &model.Bike{},
		Customer:    &model.Customer{},
		Accessories: map[int]*model.RentedAccessory{},
	}}

	flex := tview.NewFlex()

	form := tview.NewForm()
	form.SetBorder(true)

	bikes, _ := r.client.GetBikes()
	bikeOptions := util.BikesAsArray(bikes)

	customers, _ := r.client.GetCustomers()
	customerOptions := util.CustomersAsArray(customers)

	form.AddDropDown("Bakfiets", bikeOptions, 0, func(option string, optionIndex int) {
		ret.rental.Bike = bikes[optionIndex]
	})

	form.AddInputField("Verhuurdatum", ret.rental.StartDate, 50, nil, func(text string) {
		ret.rental.StartDate = text
	})

	form.AddInputField("Aantal dagen", strconv.Itoa(ret.rental.Days), 50, nil, func(text string) {
		number, _ := strconv.Atoi(text)
		ret.rental.Days = number
	})

	form.AddDropDown("Klant", customerOptions, 0, func(option string, optionIndex int) {
		ret.rental.Customer = customers[optionIndex]
	})

	accessoires, _ := r.client.GetAccessories()
	for _, accessory := range accessoires {
		a := *accessory
		acc, ok := ret.rental.Accessories[a.ID]
		if acc == nil {
			acc = &model.RentedAccessory{
				ID:     a.ID,
				Name:   a.Name,
				Price:  a.Price,
				Amount: 0,
			}
		}
		form.AddInputField(fmt.Sprintf("%s (EUR %v)", strings.TrimSpace(accessory.Name), accessory.Price),
			strconv.Itoa(acc.Amount), 5, nil, func(text string) {
				number, _ := strconv.Atoi(text)
				if ok {
					acc.Amount = number
				} else {
					ret.rental.Accessories[a.ID] = &model.RentedAccessory{
						ID:     a.ID,
						Name:   a.Name,
						Price:  a.Price,
						Amount: number,
					}
				}
			})
	}

	form.AddButton("Aanmaken", func() {
		err := r.client.CreateRental(ret.rental)
		if err != nil {
			panic(err)
		}
		r.screen.SetRoot(NewRentalListView(r), true)
	})
	form.AddButton("Annuleren", func() {
		r.screen.SetRoot(NewRentalListView(r), true)
	})

	horizontalFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	horizontalFlex.AddItem(tview.NewBox(), 0, 1, false)
	horizontalFlex.AddItem(form, 0, 1, true)
	horizontalFlex.AddItem(tview.NewBox(), 0, 1, false)

	flex.AddItem(tview.NewBox(), 0, 1, false)
	flex.AddItem(horizontalFlex, 0, 1, true)
	flex.AddItem(tview.NewBox(), 0, 1, false)

	ret.Primitive = flex

	return ret
}
