package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
}

var counter int64
var totalCounter int64
var activeRequests int64

func postRequest(wg *sync.WaitGroup, url string, payload string, concurrencyLimiter chan struct{}) {
	defer wg.Done()
	defer func() { <-concurrencyLimiter }() // release a slot

	resp, err := client.Post(url, "application/json", bytes.NewBufferString(payload))

	if err != nil {
		fmt.Printf("An Error Occurred %v\n", err)
		return
	}

	defer resp.Body.Close()
	atomic.AddInt64(&counter, 1)
	atomic.AddInt64(&totalCounter, 1)
	atomic.AddInt64(&activeRequests, -1)
}

func main() {
	var numberOfRequests int
	var maxConcurrency int
	var url string

	flag.IntVar(&numberOfRequests, "n", 10000, "Number of requests to make")
	flag.IntVar(&maxConcurrency, "c", 1000, "Maximum number of concurrent requests")
	flag.StringVar(&url, "url", "", "URL to POST to")
	flag.Parse()

	if url == "" {
		fmt.Println("URL is required")
		return
	}

	// Create a 1KB payload
	var sb strings.Builder
	sb.WriteString("{\"key\": \"")
	for sb.Len() < 1024 {
		sb.WriteRune('a')
	}
	sb.WriteString("\"}")

	payload := sb.String()

	var wg sync.WaitGroup

	concurrencyLimiter := make(chan struct{}, maxConcurrency)

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {
			if totalCounter == 0 {
				continue
			}
			fmt.Printf("Req/s: %d, Active Requests: %d\n",
				atomic.LoadInt64(&counter),
				atomic.LoadInt64(&activeRequests))

			atomic.StoreInt64(&counter, 0)
		}
	}()

	for i := 0; i < numberOfRequests; i++ {
		concurrencyLimiter <- struct{}{}
		atomic.AddInt64(&activeRequests, 1)
		wg.Add(1)
		go postRequest(&wg, url, payload, concurrencyLimiter)
	}

	wg.Wait()
	ticker.Stop()
}
