package main

import (
	"fmt"
	"mymodule/mymath"
)

func main() {
	numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	average := mymath.Average(numbers)
	fmt.Printf("The average is: %.2f\n", average)
}
