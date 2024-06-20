package main

import "fmt"

func main() {

	user := map[string]string{ // Inside [] is the key type and outside is the value type
		"name":    "Diogo",
		"surname": "Almeida",
	}

	fmt.Println(user)
	fmt.Println(user["name"]) // If you want to access a specific key

	user2 := map[string]map[string]string{
		"name": {
			"first": "John",
			"last": "Doe",
		},
		"course": {
			"name": "Engineering",
			"campus": "Campus Research",
		},
	}

	fmt.Println(user2)
	delete(user2, "name") // if you want to delete a specific key
	fmt.Println(user2)

	user2["Technology"] = map[string]string{
		"name": "Golang",
	}

	fmt.Println(user2)
}