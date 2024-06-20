package main

import "fmt"

func main() {

	number := 10

	if number > 15 {
		fmt.Println("Higher than 15")
	} else {
		fmt.Println("Less than or equal to 15")
	}

	if anotherNumber := number; anotherNumber > 0 { //anotherNumber can be used just inside the if block
		fmt.Println("Higher than 15")
	}
}