package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- i
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for num := range ch {
			fmt.Println("Num:", num)

		}
	}()
	<-sigChan
	fmt.Println("Ctrl+C")
	cancel()
	wg.Wait()
	fmt.Println("Shutting down...")

}
