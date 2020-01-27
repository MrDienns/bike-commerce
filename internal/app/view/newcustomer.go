package view

import (
	"strconv"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/rivo/tview"
)

// customerNewView represents the customer create view.
type customerNewView struct {
	*root
	tview.Primitive
	customer *model.Customer
}

// NewCustomerNewView creates a new *customerNewView and returns it.
func NewCustomerNewView(r *root) *customerNewView {
	ret := &customerNewView{root: r, customer: &model.Customer{}}

	flex := tview.NewFlex()

	form := tview.NewForm()
	form.SetBorder(true)

	form.AddInputField("Voornaam", "", 50, nil, func(text string) {
		ret.customer.Firstname = text
	})
	form.AddInputField("Achternaam", "", 50, nil, func(text string) {
		ret.customer.Lastname = text
	})
	form.AddInputField("Postcode", "", 50, nil, func(text string) {
		ret.customer.Postalcode = text
	})
	form.AddInputField("Huisnummer", "", 50, nil, func(text string) {
		number, _ := strconv.Atoi(text)
		ret.customer.Housenumber = number
	})
	form.AddInputField("Huisnummer toevoeging", "", 50, nil, func(text string) {
		ret.customer.HousenumberAddition = text
	})

	form.AddButton("Aanmaken", func() {
		err := r.client.CreateCustomer(ret.customer)
		if err != nil {
			panic(err)
		}
		r.screen.SetRoot(NewCustomerListView(r), true)
	})
	form.AddButton("Annuleren", func() {
		r.screen.SetRoot(NewCustomerListView(r), true)
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
