package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Hey!", channel)

	fmt.Println("After Sleep functions being executed")

	//for {
	//	message, open := <-channel // <- here means that the channel is receiving a value
	//	if !open { // To avoid deadlock
	//		break
	//	}
	//	fmt.Println(message)
	//}

	for message := range channel { // It does the same as the for above
		fmt.Println(message)
	}

	fmt.Println("End")

}

func write(text string, channel chan string) {
	//time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		channel <- text // <- here means that text is being sent to the channel
		time.Sleep(time.Second)
	}

	close(channel) // To avoid deadlock
}