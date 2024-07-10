package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func write(text string, nums ...int) {
	for _, number := range nums {
		fmt.Println(text, number)
	}
}

func main() {
	sumTotal := sum(1, 5, 55, 82)
	fmt.Println(sumTotal)

	write("Heey you", 10, 2, 5, 4)
}

