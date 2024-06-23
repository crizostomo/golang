package loops

import (
	"fmt"
	"time"
)

func main() {
	
	// 1st Option 
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incremeting I")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	// 2nd Option 
	for j := 0; j < 10; j++ {
		fmt.Println("Incremeting J", j)
		time.Sleep(time.Second)
	}

	names := [3]string{"Luffy", "Zoro", "Sanji"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for index, letter := range "WORD" {
		fmt.Println(index, letter)
		fmt.Println(index, string(letter))
	}

	user := map[string]string {
		"name": "Zoro",
		"Class": "Swordsman",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

	// It cannot use FOR on STRUCTS

}