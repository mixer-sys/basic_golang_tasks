package main

import (
	"fmt"
)

func main() {
	i := 1
	for i < 101 {
		if i%3 == 0 {
			fmt.Println("Fizz")
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		}
		i++
	}
}
