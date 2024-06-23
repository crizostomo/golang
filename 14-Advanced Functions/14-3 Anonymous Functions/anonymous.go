package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hey you")
	}()

	func(text string) {
		fmt.Println(text)
	}("Hey you 2")

	output := func(text string) string {
		return fmt.Sprintf("Got -> %s", text)
	}("Getting Params")

	fmt.Println(output)
}