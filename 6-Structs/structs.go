package main

import "fmt"

type user struct {
	name string
	age uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {

	var u user
	u.name = "Zoro"
	u.age = 28
	fmt.Println(u)

	addressExample := address{"Long Sea", 0}

	u2 := user{"Luffy", 28, addressExample}
	fmt.Println(u2)

	u3 := user{age: 21} // In case you don't know the name
	fmt.Println(u3)

}