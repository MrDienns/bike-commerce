package view

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MrDienns/bike-commerce/pkg/util"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/rivo/tview"
)

type rentalEditView struct {
	*root
	tview.Primitive
	rental *model.Rental
}

func NewRentalEditView(r *root, rental *model.Rental) *rentalEditView {
	ret := &rentalEditView{root: r, rental: rental}

	flex := tview.NewFlex()

	form := tview.NewForm()
	form.SetBorder(true)

	bikes, _ := r.client.GetBikes()
	bikeOptions := util.BikesAsArray(bikes)

	customers, _ := r.client.GetCustomers()
	customerOptions := util.CustomersAsArray(customers)

	users, _ := r.client.GetUsers()
	userOptions := util.UsersAsArray(users)

	form.AddDropDown("Bakfiets", bikeOptions, rental.Bike.ID-1, func(option string, optionIndex int) {
		rental.Bike = bikes[optionIndex]
	})

	form.AddInputField("Verhuurdatum", rental.StartDate, 50, nil, func(text string) {
		rental.StartDate = text
	})

	form.AddInputField("Aantal dagen", strconv.Itoa(rental.Days), 50, nil, func(text string) {
		number, _ := strconv.Atoi(text)
		rental.Days = number
	})

	form.AddDropDown("Klant", customerOptions, rental.Customer.ID-1, func(option string, optionIndex int) {
		rental.Customer = customers[optionIndex]
	})

	form.AddDropDown("Medewerker", userOptions, rental.Employee.Id-1, func(option string, optionIndex int) {
		rental.Employee = users[optionIndex]
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

	form.AddButton("Opslaan", func() {
		err := r.client.UpdateRental(rental)
		if err != nil {
			panic(err)
		}
		r.screen.SetRoot(NewRentalListView(r), true)
	})
	form.AddButton("Verwijderen", func() {
		err := r.client.DeleteRental(rental)
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
