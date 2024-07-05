package main

import (
	"fmt"
	"time"
)

func main() {
	// Concurrency != Parallelism
	go write("Hello") // goroutine
	write("Using Go")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}