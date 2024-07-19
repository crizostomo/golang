package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("Hey you!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value: %s", text)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	return channel
}