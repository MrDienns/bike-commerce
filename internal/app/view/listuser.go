package view

import "github.com/rivo/tview"

// userListView represents the user list view.
type userListView struct {
	*root
	tview.Primitive
}

// NewUserListView creates a new *userListView and returns it.
func NewUserListView(r *root) *userListView {
	ret := &userListView{root: r}

	flex := tview.NewFlex()

	list := tview.NewList().ShowSecondaryText(true)
	list.SetBorder(true)

	users, err := r.client.GetUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		u := *user
		list.AddItem(u.Name, u.Email, rune(0), func() {
			r.screen.SetRoot(NewUserEditView(r, &u), true)
		})
	}

	list.AddItem("Medewerker aanmaken", "", rune(0), func() {
		r.screen.SetRoot(NewUserNewView(r), true)
	})
	list.AddItem("Terug naar hoofdmenu", "", rune(0), func() {
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
