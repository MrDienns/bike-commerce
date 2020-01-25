package view

import (
	"fmt"
	"strconv"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/rivo/tview"
)

type accessoryEditView struct {
	*root
	tview.Primitive
	accessory *model.Accessory
}

func NewAccessoryEditView(r *root, accessory *model.Accessory) *accessoryEditView {
	ret := &accessoryEditView{root: r, accessory: accessory}

	flex := tview.NewFlex()

	form := tview.NewForm()
	form.SetBorder(true)

	form.AddInputField("Naam", accessory.Name, 50, nil, func(text string) {
		accessory.Name = text
	})
	form.AddInputField("Prijs", fmt.Sprintf("%v", accessory.Price), 50, nil, func(text string) {
		number, _ := strconv.ParseFloat(text, 32)
		accessory.Price = float32(number)
	})

	form.AddButton("Opslaan", func() {
		err := r.client.UpdateAccessory(accessory)
		if err != nil {
			panic(err)
		}
		r.screen.SetRoot(NewAccessoryListView(r), true)
	})
	form.AddButton("Verwijderen", func() {
		err := r.client.DeleteAccessory(accessory)
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
