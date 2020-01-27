package view

import (
	"strconv"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/rivo/tview"
)

// bikeNewView represents the bike create view.
type bikeNewView struct {
	*root
	tview.Primitive
	bike *model.Bike
}

// NewBikeNewView creates a new *bikeNewView and returns it.
func NewBikeNewView(r *root) *bikeNewView {
	ret := &bikeNewView{root: r, bike: &model.Bike{}}

	flex := tview.NewFlex()

	form := tview.NewForm()
	form.SetBorder(true)

	form.AddInputField("Naam", "", 50, nil, func(text string) {
		ret.bike.Name = text
	})
	form.AddInputField("Type", "", 50, nil, func(text string) {
		ret.bike.Type = text
	})
	form.AddInputField("Prijs", "", 50, nil, func(text string) {
		number, _ := strconv.ParseFloat(text, 32)
		ret.bike.Price = float32(number)
	})
	form.AddInputField("Hoeveelheid", "", 50, nil, func(text string) {
		number, _ := strconv.Atoi(text)
		ret.bike.Quantity = number
	})
	form.AddInputField("Verhuurd", "", 50, nil, func(text string) {
		number, _ := strconv.Atoi(text)
		ret.bike.AmountRented = number
	})

	form.AddButton("Aanmaken", func() {
		err := r.client.CreateBike(ret.bike)
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
