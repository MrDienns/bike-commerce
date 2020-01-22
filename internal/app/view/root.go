package view

import (
	"github.com/MrDienns/bike-commerce/pkg/api"
	"github.com/rivo/tview"
)

// root is the highest level struct responsible for communicating with all underlying views.
type root struct {
	screen *tview.Application
	login  *loginView
	client *api.Client
}

// NewRoot creates and returns a new *root. It initialises a blank *tview.Application in the struct.
func NewRoot(client *api.Client) *root {
	return &root{
		client: client,
		screen: tview.NewApplication(),
	}
}

// Start initialises the underlying views and opens the application in the terminal.
func (r *root) Start() error {
	r.login = NewLoginView(r)
	r.screen.SetRoot(r.login, true)
	return r.screen.Run()
}
