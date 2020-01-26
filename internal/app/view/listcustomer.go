package view

import (
	"fmt"

	"github.com/rivo/tview"
)

type customerListView struct {
	*root
	tview.Primitive
}

func NewCustomerListView(r *root) *customerListView {
	ret := &customerListView{root: r}

	flex := tview.NewFlex()

	list := tview.NewList().ShowSecondaryText(true)
	list.SetBorder(true)

	customers, err := r.client.GetCustomers()
	if err != nil {
		panic(err)
	}
	for _, customer := range customers {
		c := *customer
		list.AddItem(fmt.Sprintf("%v %v", c.Firstname, c.Lastname),
			fmt.Sprintf("%v - %v%v", c.Postalcode, c.Housenumber, c.HousenumberAddition), rune(0), func() {
				r.screen.SetRoot(NewCustomerEditView(r, &c), true)
			})
	}

	list.AddItem("Klant aanmaken", "", rune(0), func() {
		r.screen.SetRoot(NewCustomerNewView(r), true)
	})

	list.AddItem("Terug naar hoofdmenu", "", rune(0), func() {
		r.screen.SetRoot(NewMenu(r), true)
	})

	horizontalFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	horizontalFlex.AddItem(tview.NewBox(), 0, 1, false)
	horizontalFlex.AddItem(list, 0, 1, true)
	horizontalFlex.AddItem(tview.NewBox(), 0, 1, false)

	flex.AddItem(tview.NewBox(), 0, 1, false)
	flex.AddItem(horizontalFlex, 0, 1, true)
	flex.AddItem(tview.NewBox(), 0, 1, false)

	ret.Primitive = flex

	return ret
}
