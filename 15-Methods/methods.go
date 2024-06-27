package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving user data %s on the database\n", u.name)
}

func (u user) olderThan18() bool {
	return u.age >= 18
}

func (u *user) birthday() {
	u.age++
}

func main() {
	user1 := user{"Luffy", 18}
	fmt.Println(user1)
	user1.save()
	
	olderThan18 := user1.olderThan18()
	fmt.Println(olderThan18)

	user1.birthday()
	fmt.Println(user1.age)

}