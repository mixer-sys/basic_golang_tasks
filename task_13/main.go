package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- i
				fmt.Printf("Sent value: %d\n", i)
				i++
				time.Sleep(1 * time.Millisecond)
			}
		}

	}()

	workerPool(ctx, ch, 3)

}

func workerPool(ctx context.Context, ch chan int, numWorkers int) {
	wg := &sync.WaitGroup{}
	for i := range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case value, ok := <-ch:
					if !ok {
						return
					}
					time.Sleep(2 * time.Millisecond)
					fmt.Printf("Worker %d processing value: %d\n", i, value)
					time.Sleep(20 * time.Millisecond)
					if value%2 == 0 && value != 0 {
						fmt.Printf("Number %d is divisible by 2\n", value)
					}
				default:
					time.Sleep(1 * time.Millisecond)
				}
			}
		}()

	}
	wg.Wait()

}
