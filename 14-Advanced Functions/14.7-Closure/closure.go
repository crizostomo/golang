package main

import "fmt"

func closure() func() {
	text := "Inside closure"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {

	text := "Inside main"
	fmt.Println(text)

	newFunction := closure()
	newFunction()

}