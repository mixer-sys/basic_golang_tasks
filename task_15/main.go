package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	URLS := []string{
		"https://ya.ru",
		"https://google.com",
		"https://github.com",
		"https://ident.me",
		"https://ipinfo.io/ip",
	}

	ch := make(chan string)
	wg := sync.WaitGroup{}

	for _, url := range URLS {
		wg.Add(1)
		go func(url string) {
			start_time := time.Now()
			fmt.Printf("Fetching %s\n", url)
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				ch <- fmt.Sprintf("Error reading body from %s: %v", url, err)
				return
			}
			end_time := time.Since(start_time).Seconds()
			ch <- fmt.Sprintf("Response from %s: %d, time in seconds: %f",
				url, len(body), end_time)
		}(url)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		fmt.Println(res)
	}
}
