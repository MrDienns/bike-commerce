package view

import (
	"strconv"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/rivo/tview"
)

type accessoryNewView struct {
	*root
	tview.Primitive
	accessory *model.Accessory
}

func NewAccessoryNewView(r *root) *accessoryNewView {
	ret := &accessoryNewView{root: r, accessory: &model.Accessory{}}

	flex := tview.NewFlex()

	form := tview.NewForm()
	form.SetBorder(true)

	form.AddInputField("Naam", "", 50, nil, func(text string) {
		ret.accessory.Name = text
	})
	form.AddInputField("Prijs", "", 50, nil, func(text string) {
		number, _ := strconv.ParseFloat(text, 32)
		ret.accessory.Price = float32(number)
	})

	form.AddButton("Aanmaken", func() {
		err := r.client.CreateAccessory(ret.accessory)
		if err != nil {
			panic(err)
		}
		r.screen.SetRoot(NewAccessoryListView(r), true)
	})
	form.AddButton("Annuleren", func() {
		r.screen.SetRoot(NewAccessoryListView(r), true)
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
