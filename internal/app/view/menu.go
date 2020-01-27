package view

import "github.com/rivo/tview"

// menuView represents the main menu view.
type menuView struct {
	*root
	tview.Primitive
}

// NewMenu creates a new *menuView and returns it.
func NewMenu(r *root) *menuView {

	ret := &menuView{root: r}

	flex := tview.NewFlex()

	list := tview.NewList().ShowSecondaryText(false)
	list.SetBorder(true)
	list.AddItem("Beheer klanten", "", rune(49), func() {
		r.screen.SetRoot(NewCustomerListView(r), true)
	})

	list.AddItem("Beheer medewerkers", "", rune(50), func() {
		r.screen.SetRoot(NewUserListView(r), true)
	})

	list.AddItem("Beheer bakfietsen", "", rune(51), func() {
		r.screen.SetRoot(NewBikeListView(r), true)
	})

	list.AddItem("Beheer accessoires", "", rune(52), func() {
		r.screen.SetRoot(NewAccessoryListView(r), true)
	})

	list.AddItem("Beheer verhuur", "", rune(53), func() {
		r.screen.SetRoot(NewRentalListView(r), true)
	})

	list.AddItem("Uitloggen", "", rune(54), func() {
		r.client.User = nil
		r.client.Token = ""
		r.screen.SetRoot(NewLoginView(r, ""), true)
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
