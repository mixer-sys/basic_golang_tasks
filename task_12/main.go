package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	NUM := 101
	wg := &sync.WaitGroup{}

	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range NUM {
			time.Sleep(5 * time.Second) // Simulate some work
			ch <- i

		}
		close(ch)
	}()

	sum := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			sum += num
			fmt.Printf("Received number: %d, Current sum: %d\n", num, sum)
		}
	}()

	wg.Wait()

	fmt.Printf("Sum of numbers from 1 to %d is %d\n", NUM-1, sum)
}
