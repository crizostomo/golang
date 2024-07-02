package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting point")

	application := app.Generate()

	error := application.Run(os.Args)

	if error != nil {
		log.Fatal(error)
	}
}