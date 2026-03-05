package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	url := flag.String("url", "http://localhost:8080/api/v1/catalog", "target URL")
	token := flag.String("token", "dev-token", "auth token")
	requests := flag.Int("requests", 1000, "total requests")
	concurrency := flag.Int("concurrency", 20, "parallel workers")
	flag.Parse()

	start := time.Now()
	var okCount int64
	var failCount int64
	jobs := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{Timeout: 5 * time.Second}
			for range jobs {
				req, _ := http.NewRequest(http.MethodGet, *url, nil)
				req.Header.Set("Authorization", "Bearer "+*token)
				res, err := client.Do(req)
				if err != nil {
					atomic.AddInt64(&failCount, 1)
					continue
				}
				if res.StatusCode >= 200 && res.StatusCode < 300 {
					atomic.AddInt64(&okCount, 1)
				} else {
					atomic.AddInt64(&failCount, 1)
				}
				_ = res.Body.Close()
			}
		}()
	}

	for i := 0; i < *requests; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	duration := time.Since(start)
	fmt.Printf("done requests=%d ok=%d fail=%d duration=%s rps=%.2f\n", *requests, okCount, failCount, duration, float64(*requests)/duration.Seconds())
}
