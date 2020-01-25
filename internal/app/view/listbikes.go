package view

import "github.com/rivo/tview"

type bikeListView struct {
	*root
	tview.Primitive
}

func NewBikeListView(r *root) *bikeListView {
	ret := &bikeListView{root: r}

	flex := tview.NewFlex()

	list := tview.NewList().ShowSecondaryText(true)
	list.SetBorder(true)

	bikes, err := r.client.GetBikes()
	if err != nil {
		panic(err)
	}
	for _, bike := range bikes {
		b := *bike
		list.AddItem(b.Name, b.Type, rune(0), func() {
			r.screen.SetRoot(NewBikeEditView(r, &b), true)
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
