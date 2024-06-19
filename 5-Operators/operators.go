package main

import "fmt"

func main() {

	// ARITMÉTICOS: + - / * %

	//var n1 int16 = 10
	//var n2 int32 = 25
	//sum := n1 + n2
	//fmt.Println(sum) // It doesn't work since they are different types

	// RELATIONAL OPERATORS: > >= == <= < !=

	// LOGICAL OPEATORS: && || !

	// UNÁRIOS OPERATORS
	number := 10
	number++
	number += 15 // = number = numbner + 15
	fmt.Println(number)

	number *= 3 // number = number * 3
	number /= 10
	number %= 3

	// Golang uses only IF and ELSE and does not support Terniary operators
}