package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // Quantity of goroutines

	go func() {
		write("Hello")
		waitGroup.Done()
	}()

	go func() {
		write("Using Go")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}