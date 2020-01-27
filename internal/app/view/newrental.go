package view

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/MrDienns/bike-commerce/pkg/util"
	"github.com/rivo/tview"
)

// rentalNewView represents the rental create view.
type rentalNewView struct {
	*root
	tview.Primitive
	rental *model.Rental
}

// NewRentalNewView creates a new *rentalNewView and returns it.
func NewRentalNewView(r *root) *rentalNewView {
	ret := &rentalNewView{root: r, rental: &model.Rental{
		Employee:    r.client.User,
		Bike:        &model.Bike{},
		Customer:    &model.Customer{},
		Accessories: map[int]*model.RentedAccessory{},
	}}

	rentalFlex := tview.NewFlex()
	rentalFlex.SetBorder(true)

	receipt := textView(ret.rental)

	flex := tview.NewFlex()

	form := tview.NewForm()

	bikes, _ := r.client.GetBikes()
	bikeOptions := util.BikesAsArray(bikes)

	customers, _ := r.client.GetCustomers()
	customerOptions := util.CustomersAsArray(customers)

	form.AddDropDown("Bakfiets", bikeOptions, 0, func(option string, optionIndex int) {
		ret.rental.Bike = bikes[optionIndex]
		receipt.SetText(receiptText(ret.rental))
	})

	form.AddInputField("Verhuurdatum", ret.rental.StartDate, 50, nil, func(text string) {
		ret.rental.StartDate = text
		receipt.SetText(receiptText(ret.rental))
	})

	form.AddInputField("Aantal dagen", strconv.Itoa(ret.rental.Days), 50, nil, func(text string) {
		number, _ := strconv.Atoi(text)
		ret.rental.Days = number
		receipt.SetText(receiptText(ret.rental))
	})

	form.AddDropDown("Klant", customerOptions, 0, func(option string, optionIndex int) {
		ret.rental.Customer = customers[optionIndex]
	})

	accessoires, _ := r.client.GetAccessories()

	accessoriesArr := make([]*model.Accessory, len(accessoires))
	for i, a := range accessoires {
		accessoriesArr[i] = *&a
	}

	for _, accessory := range accessoriesArr {
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
				receipt.SetText(receiptText(ret.rental))
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

	rentalFlex.SetDirection(tview.FlexRow)
	rentalFlex.AddItem(form, 0, 1, true)
	rentalFlex.AddItem(receipt, 0, 1, false)

	horizontalFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	horizontalFlex.AddItem(tview.NewBox(), 0, 1, false)
	horizontalFlex.AddItem(rentalFlex, 0, 6, true)
	horizontalFlex.AddItem(tview.NewBox(), 0, 1, false)

	flex.AddItem(tview.NewBox(), 0, 1, false)
	flex.AddItem(horizontalFlex, 0, 1, true)
	flex.AddItem(tview.NewBox(), 0, 1, false)

	ret.Primitive = flex

	return ret
}
