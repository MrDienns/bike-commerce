package view

import (
	"fmt"

	"github.com/rivo/tview"
)

type accessoryListView struct {
	*root
	tview.Primitive
}

func NewAccessoryListView(r *root) *accessoryListView {
	ret := &accessoryListView{root: r}

	flex := tview.NewFlex()

	list := tview.NewList().ShowSecondaryText(true)
	list.SetBorder(true)

	accessories, err := r.client.GetAccessories()
	if err != nil {
		panic(err)
	}
	for _, accessory := range accessories {
		a := *accessory
		list.AddItem(a.Name, fmt.Sprintf("%v", a.Price), rune(0), func() {
			r.screen.SetRoot(NewAccessoryEditView(r, &a), true)
		})
	}

	list.AddItem("Terug naar hoofdmenu", "", (0), func() {
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