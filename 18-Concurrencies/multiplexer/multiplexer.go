package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(write("Hey you!"), write("Hey from Go"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(enterChannel1, enterChannel2 <-chan string) <-chan string {
	exitChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-enterChannel1:
				exitChannel <- message
			case message := <-enterChannel2:
				exitChannel <- message
			}
		}
	}()

	return exitChannel
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value: %s", text)
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		}
	}()

	return channel
}