package main

import (
	"fmt"
	"module/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from the main file")
	auxiliar.Write()

	error := checkmail.ValidateFormat("123")
	fmt.Println(error)
}