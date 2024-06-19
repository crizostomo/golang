package main

import (
	"errors"
	"fmt"
)

func main() {

	// REAL NUMBERS - START

	var number int16 = 100 //int8, int16, int32, int64 AND int (based on your PC system for example 64bitz)
	fmt.Println(number)

	var number2 uint32 = 1000 //uint = unsigned int --> ONLY CONTAINS POSITIVE NUMBERS
	fmt.Println(number2)

	//alias
	var number3 rune = 12456 //rune is an alias for int32
	fmt.Println(number3)

	var number4 byte = 123 //byte is an alias for int8
	fmt.Println(number4)

	
	var realNumber1 float32 = 123.45
	fmt.Println(realNumber1)

	var realNumber2 float64 = 123.4588888888
	fmt.Println(realNumber2)

	// OR

	realNumber3 := 12345.67
	fmt.Println(realNumber3)

	// REAL NUMBERS - FINISH


	// STRINGS - START

	var str string = "Text"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char) // It will take the number in the ASCII table

	// STRINGS - FINISH

	var text string
	fmt.Println(text)

	// BOOLEAN - START

	var boolean1 bool = true
	fmt.Println(boolean1)

	// BOOLEAN - FINISH

	// ERROR - START

	var error error = errors.New("Internal Error")
	fmt.Println(error)

	// ERROR - FINISH

}