package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Hey you!"

	message := <-channel
	fmt.Println(message)
}