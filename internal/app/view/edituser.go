package view

import (
	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/rivo/tview"
)

type userEditView struct {
	*root
	tview.Primitive
	user *model.User
}

func NewUserEditView(r *root, user *model.User) *userEditView {
	ret := &userEditView{root: r, user: user}

	flex := tview.NewFlex()

	form := tview.NewForm()
	form.SetBorder(true)

	form.AddInputField("Naam", user.Name, 50, nil, func(text string) {
		user.Name = text
	})
	form.AddInputField("E-mail", user.Email, 50, nil, func(text string) {
		user.Email = text
	})
	form.AddInputField("Datum in dienst", user.EmploymentDate, 50, nil, func(text string) {
		user.EmploymentDate = text
	})

	form.AddButton("Opslaan", func() {
		err := r.client.UpdateUser(user)
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
