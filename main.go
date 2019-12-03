package main

import (
	"fmt"

	"github.com/MrDienns/bike-commerce/view"
)

func main() {
	screen := view.NewRoot()
	err := screen.Start()
	if err != nil {
		fmt.Printf("Failed to start application: %v\n", err)
	}
}
