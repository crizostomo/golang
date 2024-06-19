package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathEquasions(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subtraction := n1 - n2
	return sum, subtraction
}

func main() {
	sum := sum(10, 20)
	fmt.Println(sum)

	var f = func() {
		fmt.Println("f Function")
	}

	f()

	var f1 = func(txt string) {
		fmt.Println(txt)
	}

	f1("Function 1 Text")

	var f2 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f2("Function 2 Text")
	fmt.Println(result)

	// By passing 2 values
	sumResult, subtractionResult := mathEquasions(10, 15)
	fmt.Println(sumResult, subtractionResult)

	// By passing 1 value from a 2 values func, we put a _ 
	sumResult1, _ := mathEquasions(10, 15)
	fmt.Println(sumResult1)
}