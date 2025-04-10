package main

import "fmt"

func main() {
	numbers := [5]int8{30, 64, 19, 88, 21}
	fmt.Println("nilai terbesar adalah", bigNumber(numbers))
}

func bigNumber(numbers [5]int8 ) int8 {
	var result int8

	for index, number := range numbers {
		if index == 0 {
			result = number
		} else {
			if number > result {
				result = number
			}
		}
	}
	return result
}