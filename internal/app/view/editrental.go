package view

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/MrDienns/bike-commerce/pkg/util"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/rivo/tview"
)

// rentalEditView represents the rental edit view.
type rentalEditView struct {
	*root
	tview.Primitive
	rental *model.Rental
}

// NewRentalEditView creates a new *rentalEditView and returns it.
func NewRentalEditView(r *root, rental *model.Rental) *rentalEditView {
	ret := &rentalEditView{root: r, rental: rental}

	rentalFlex := tview.NewFlex()
	rentalFlex.SetBorder(true)

	receipt := textView(rental)

	flex := tview.NewFlex()

	form := tview.NewForm()

	bikes, _ := r.client.GetBikes()
	bikeOptions := util.BikesAsArray(bikes)

	customers, _ := r.client.GetCustomers()
	customerOptions := util.CustomersAsArray(customers)

	users, _ := r.client.GetUsers()
	userOptions := util.UsersAsArray(users)

	form.AddDropDown("Bakfiets", bikeOptions, rental.Bike.ID-1, func(option string, optionIndex int) {
		rental.Bike = bikes[optionIndex]
		receipt.SetText(receiptText(ret.rental))
	})

	form.AddInputField("Verhuurdatum", rental.StartDate, 50, nil, func(text string) {
		rental.StartDate = text
		receipt.SetText(receiptText(ret.rental))
	})

	form.AddInputField("Aantal dagen", strconv.Itoa(rental.Days), 50, nil, func(text string) {
		number, _ := strconv.Atoi(text)
		rental.Days = number
		receipt.SetText(receiptText(ret.rental))
	})

	form.AddDropDown("Klant", customerOptions, rental.Customer.ID-1, func(option string, optionIndex int) {
		rental.Customer = customers[optionIndex]
	})

	form.AddDropDown("Medewerker", userOptions, rental.Employee.Id-1, func(option string, optionIndex int) {
		rental.Employee = users[optionIndex]
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

func textView(rental *model.Rental) *tview.TextView {
	text := tview.NewTextView()
	text.SetText(receiptText(rental))
	return text
}

func receiptText(rental *model.Rental) string {

	receipt := `
Verhuur overzicht:

Fiets:
	Per dag: %v EUR
	Aantal dagen: %v
	Totaal: %v EUR
`

	receipt = fmt.Sprintf(receipt, rental.Bike.Price, rental.Days, rental.Bike.Price*float32(rental.Days))

	dailyPrice := rental.Bike.Price

	if len(rental.Accessories) > 0 {
		receipt = receipt + "\n\nAccessoires:\n"
		accessoires := rental.Accessories

		accessoriesArr := make([]*model.RentedAccessory, len(accessoires))
		i := 0
		for _, a := range accessoires {
			accessoriesArr[i] = *&a
			i++
		}

		sort.Slice(accessoriesArr, func(i, j int) bool {
			return accessoriesArr[i].ID < accessoriesArr[j].ID
		})

		for _, accessory := range accessoriesArr {
			dailyPrice += accessory.Price * float32(accessory.Amount)
			receipt = receipt + fmt.Sprintf("\t%v - %v EUR (%v EUR totaal)\n",
				strings.TrimSpace(accessory.Name), accessory.Price, accessory.Price*float32(accessory.Amount))
		}
	}

	receipt = receipt + "\n\nTotaal:"
	receipt = receipt + fmt.Sprintf("\n\tPer dag: %v", dailyPrice)
	receipt = receipt + fmt.Sprintf("\n\tTotaal: %v", dailyPrice*float32(rental.Days))

	return receipt
}
