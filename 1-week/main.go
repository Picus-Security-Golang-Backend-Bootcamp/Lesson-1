package main

import (
	"github.com/mehmetcantas/patika/helper"
)

func main() {
	// var myVar int = 64
	// var myName = 20.0
	// salary := 10

	message := "hello world"
	message = helper.UppercaseString(message)
	helper.printMessage(message)
}
