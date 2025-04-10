package main

import "fmt"

func main() {
	numbers := [5]int8{30, 64, 19, 88, 21}

	newNumbers := numbers[1:4]

	fmt.Println(newNumbers)
}