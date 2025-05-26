package main

import (
	"fmt"
)

func SumDigits(n int) int {
	if n < 0 {
		n = -n
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(SumDigits(123))
	fmt.Println(SumDigits(9875))
}
