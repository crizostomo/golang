package main

import "fmt"

func invertSignal(number int) int {
	return number * -1
}

func invertSignalWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	invertedNumber := invertSignal(number)
	fmt.Println(invertedNumber)

	newNumber := 40
	invertSignalWithPointer(&newNumber)
	fmt.Println(newNumber)
}