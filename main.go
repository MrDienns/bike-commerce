package main

import (
	"fmt"

	"github.com/MrDienns/bike-commerce/view"
)

func main() {
	screen := view.NewRoot()
	if err := screen.Start(); err != nil {
		fmt.Printf("Failed to start application: %v\n", err)
	}
}
