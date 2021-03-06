package view

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// loginView represents the view used to login with.
type loginView struct {
	*root
	tview.Primitive
	username, password string
}

// NewLoginView creates the flex layout and all necessary boxes in order to create a proper layout for the login view.
func NewLoginView(r *root, err string) *loginView {

	ret := &loginView{root: r}

	flex := tview.NewFlex()

	form := tview.NewForm()
	form.SetBorder(true)
	form.AddInputField("Email", "", 100, nil, func(username string) {
		ret.username = username
	})
	form.AddPasswordField("Wachtwoord", "", 100, rune(42), func(password string) {
		ret.password = password
	})
	form.AddButton("Login", func() {
		ret.login(ret.username, ret.password)
	})

	frame := tview.NewFrame(form)
	frame.AddText("Welkom bij Van der Binckes", true, tview.AlignCenter, tcell.ColorWhite)
	frame.AddText(err, false, tview.AlignCenter, tcell.ColorRed)

	horizontalFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	horizontalFlex.AddItem(tview.NewBox(), 0, 1, false)
	horizontalFlex.AddItem(frame, 0, 1, true)
	horizontalFlex.AddItem(tview.NewBox(), 0, 1, false)

	flex.AddItem(tview.NewBox(), 0, 1, false)
	flex.AddItem(horizontalFlex, 0, 1, true)
	flex.AddItem(tview.NewBox(), 0, 1, false)

	ret.Primitive = flex

	return ret
}

// login takes two parameters and tries to invoke
func (lv *loginView) login(username, password string) {
	_, _, err := lv.client.Authenticate(username, password)
	if err != nil {
		lv.root.screen.SetRoot(NewLoginView(lv.root, err.Error()), true)
		return
	}
	lv.root.screen.SetRoot(NewMenu(lv.root), true)
}
