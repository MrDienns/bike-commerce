package view

import (
	"fmt"
	"strconv"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/rivo/tview"
)

type bikeEditView struct {
	*root
	tview.Primitive
	bike *model.Bike
}

func NewBikeEditView(r *root, bike *model.Bike) *bikeEditView {
	ret := &bikeEditView{root: r, bike: bike}

	flex := tview.NewFlex()

	form := tview.NewForm()
	form.SetBorder(true)

	form.AddInputField("Naam", bike.Name, 50, nil, func(text string) {
		bike.Name = text
	})
	form.AddInputField("Type", bike.Type, 50, nil, func(text string) {
		bike.Type = text
	})
	form.AddInputField("Prijs", fmt.Sprintf("%v", bike.Price), 50, nil, func(text string) {
		number, _ := strconv.ParseFloat(text, 32)
		bike.Price = float32(number)
	})
	form.AddInputField("Hoeveelheid", strconv.Itoa(bike.Quantity), 50, nil, func(text string) {
		number, _ := strconv.Atoi(text)
		bike.Quantity = number
	})
	form.AddInputField("Verhuurd", strconv.Itoa(bike.AmountRented), 50, nil, func(text string) {
		number, _ := strconv.Atoi(text)
		bike.AmountRented = number
	})

	form.AddButton("Opslaan", func() {
		err := r.client.UpdateBike(bike)
		if err != nil {
			panic(err)
		}
		r.screen.SetRoot(NewBikeListView(r), true)
	})
	form.AddButton("Annuleren", func() {
		r.screen.SetRoot(NewBikeListView(r), true)
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
