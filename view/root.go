package view

import "github.com/rivo/tview"

// root is the highest level struct responsible for communicating with all underlying views.
type root struct {
	screen *tview.Application
}

// NewRoot creates and returns a new *root. It initialises a blank *tview.Application in the struct.
func NewRoot() *root {
	return &root{
		screen: tview.NewApplication(),
	}
}

// Start initialises the underlying views and opens the application in the terminal.
func (r *root) Start() error {
	login := NewLoginView(r)
	r.screen.SetRoot(login, true)
	return r.screen.Run()
}
