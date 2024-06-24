package main

import "fmt"

func func1() {
	fmt.Println("Executing func 1")
}

func func2() {
	fmt.Println("Executing func 2")
}

func studentPassed(n1, n2 float32) bool {
	defer fmt.Println("Average being calculated")
	fmt.Println("Loading Student's grades")

	average := (n1 + n2) / 2

	if average > 6 {
		return true
	}

	return false
}

func main() {

	defer func1() // It postpones this func
	func2()

	fmt.Println(studentPassed(7, 8))
}