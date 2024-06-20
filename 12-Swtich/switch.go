package main

import "fmt"

func weekDay(number int) string {
	switch number {
	case 1:
		return "Sun"
	case 2:
		return "Mon"
	case 3:
		return "Tue"
	case 4:
		return "Wed"
	case 5:
		return "Thu"
	case 6:
		return "Fri"
	case 7:
		return "Sat"
	default:
		return "Unknown"
	}
}

func weekDay2(number int) string {
	switch {
	case number == 1:
		return "Sunday"
	case number == 2:
		return "Monday"
	case number == 3:
		return "Tuesday"
	case number == 4:
		return "Wednesday"
	case number == 5:
		return "Thursday"
	case number == 6:
		return "Friday"
	case number == 7:
		return "Saturday"
	default:
		return "Unknown"							
	}
}

func weekDay3(number int) string {
	var weekDay string

	switch {
	case number == 1:
		weekDay = "Sunday"
		fallthrough // If we execute with value 1, fallthrough will lead to the following value
	case number == 2:
		weekDay = "Monday"
	case number == 3:
		weekDay = "Tuesday"
	case number == 4:
		weekDay = "Wednesday"
	case number == 5:
		weekDay = "Thursday"
	case number == 6:
		weekDay = "Friday"
	case number == 7:
		weekDay = "Saturday"
	default:
		weekDay = "Unknown"							
	}

	return weekDay
}

func main() {

	day := weekDay(3)
	fmt.Println(day)
}