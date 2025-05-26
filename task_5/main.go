package main

import (
	"fmt"
)

func FilterEvern(numbers ...int) []int {

	var evenNumbers []int
	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		}
	}
	return evenNumbers
}

func main() {
	fmt.Println(FilterEvern(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
