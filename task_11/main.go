package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	NUM := 5
	wg := &sync.WaitGroup{}

	for i := range NUM {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(3 * time.Second)
			fmt.Printf("Hello from goroutine %d\n", i)
		}()
	}
	wg.Wait()
}
