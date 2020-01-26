package view

import (
	"fmt"

	"github.com/rivo/tview"
)

type rentalListView struct {
	*root
	tview.Primitive
}

func NewRentalListView(r *root) *rentalListView {
	ret := &rentalListView{root: r}

	flex := tview.NewFlex()

	list := tview.NewList().ShowSecondaryText(true)
	list.SetBorder(true)

	rentals, err := r.client.GetRentals()
	if err != nil {
		panic(err)
	}
	for _, rental := range rentals {
		re := *rental
		list.AddItem(fmt.Sprintf("%s - %s %s", re.Bike.Name, re.Customer.Firstname, re.Customer.Lastname), re.StartDate, rune(0), func() {
			r.screen.SetRoot(NewRentalEditView(r, &re), true)
		})
	}

	list.AddItem("Verhuur aanmaken", "", rune(0), func() {
		r.screen.SetRoot(NewRentalNewView(r), true)
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
