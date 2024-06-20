package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]string
	array1[0] = "1st"
	fmt.Println(array1)

	array2 := [2]string{"1st", "2nd"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // [...]keeps arrays a lit bit dynamic
	fmt.Println(array3)

	slice := []int{10, 11, 12} //It's not an array, it just looks like
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array1[1:3] //Starts index1 and stops on index2
	fmt.Println(slice2)

	array1[1] = "Position Changed"
	fmt.Println(array1)

	// INTERNAL ARRAYS
	slice3 := make([]float32, 10, 11) // If you don't pass the last parameter 'Go' assumes that the second parameter is the capacity
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // lenght
	fmt.Println(cap(slice3)) // capacity

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6) // here you would exceed the capacity but 'Golang' doubles it
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // lenght
	fmt.Println(cap(slice3)) // capacity
}