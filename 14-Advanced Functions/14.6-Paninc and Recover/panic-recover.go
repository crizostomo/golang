package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Execution recovered successfully")
	}
}

func studentPassed(n1, n2 float32) bool {
	defer recoverExecution()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("Average is exactly 6!")
}

func main() {
	fmt.Println(studentPassed(6, 6))
	fmt.Println("Post Execution")
}