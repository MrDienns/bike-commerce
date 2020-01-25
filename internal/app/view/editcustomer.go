package view

import (
	"strconv"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/rivo/tview"
)

type customerEditView struct {
	*root
	tview.Primitive
	customer *model.Customer
}

func NewCustomerEditView(r *root, customer *model.Customer) *customerEditView {
	ret := &customerEditView{root: r, customer: customer}

	flex := tview.NewFlex()

	form := tview.NewForm()
	form.SetBorder(true)

	form.AddInputField("Voornaam", customer.Firstname, 50, nil, func(text string) {
		customer.Firstname = text
	})
	form.AddInputField("Achternaam", customer.Lastname, 50, nil, func(text string) {
		customer.Lastname = text
	})
	form.AddInputField("Postcode", customer.Postalcode, 50, nil, func(text string) {
		customer.Postalcode = text
	})
	form.AddInputField("Huisnummer", strconv.Itoa(customer.Housenumber), 50, nil, func(text string) {
		number, _ := strconv.Atoi(text)
		customer.Housenumber = number
	})
	form.AddInputField("Huisnummer toevoeging", customer.HousenumberAddition, 50, nil, func(text string) {
		customer.HousenumberAddition = text
	})

	form.AddButton("Opslaan", func() {
		err := r.client.UpdateCustomer(customer)
		if err != nil {
			panic(err)
		}
		r.screen.SetRoot(NewCustomerListView(r), true)
	})
	form.AddButton("Verwijderen", func() {
		err := r.client.DeleteCustomer(customer)
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
