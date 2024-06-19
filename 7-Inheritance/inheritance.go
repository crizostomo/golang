package main 

import "fmt"

type person struct {
	name string
	surname string
	age uint8
	height uint8
}

type student struct {
	person
	course string
	college string
}

func main() {

	p1 := person{"Naruto", "Uzumaki", 22, 178}
	fmt.Println(p1)

	e1 := student{p1, "Ninja", "Leaf Village"}
	fmt.Println(e1.height)

}