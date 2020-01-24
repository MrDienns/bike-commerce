package view

import "github.com/rivo/tview"

type menuView struct {
	*root
	tview.Primitive
}

func NewMenu(r *root) *menuView {

	ret := &menuView{root: r}

	flex := tview.NewFlex()

	list := tview.NewList().ShowSecondaryText(false)
	list.SetBorder(true)
	list.AddItem("Beheer klanten", "", rune(49), func() {

	})

	list.AddItem("Beheer medewerkers", "", rune(50), func() {

	})

	list.AddItem("Beheer bakfietsen", "", rune(51), func() {

	})

	list.AddItem("Beheer accessoires", "", rune(52), func() {

	})

	list.AddItem("Beheer verhuur", "", rune(53), func() {

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
