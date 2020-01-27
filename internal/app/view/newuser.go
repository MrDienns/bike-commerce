package view

import (
	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/rivo/tview"
)

// userNewView represents the user create view.
type userNewView struct {
	*root
	tview.Primitive
	user *model.User
}

// NewUserNewView creates a new *userNewView and returns it.
func NewUserNewView(r *root) *userNewView {
	ret := &userNewView{root: r, user: &model.User{}}

	flex := tview.NewFlex()

	form := tview.NewForm()
	form.SetBorder(true)

	form.AddInputField("Naam", "", 50, nil, func(text string) {
		ret.user.Name = text
	})
	form.AddInputField("E-mail", "", 50, nil, func(text string) {
		ret.user.Email = text
	})
	form.AddInputField("Datum in dienst", "", 50, nil, func(text string) {
		ret.user.EmploymentDate = text
	})

	form.AddButton("Aanmaken", func() {
		err := r.client.CreateUser(ret.user)
		if err != nil {
			panic(err)
		}
		r.screen.SetRoot(NewUserListView(r), true)
	})
	form.AddButton("Annuleren", func() {
		r.screen.SetRoot(NewUserListView(r), true)
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
