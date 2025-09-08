package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Result struct {
	URL        string
	Status     string
	DurationMS int64
	Bytes      int
	Err        string
}

func fetch(url string) Result {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		return Result{
			URL:        url,
			Err:        err.Error(),
			DurationMS: time.Since(start).Milliseconds(),
		}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{
			URL:        url,
			Status:     resp.Status,
			Err:        err.Error(),
			DurationMS: time.Since(start).Milliseconds(),
		}
	}

	return Result{
		URL:        url,
		Status:     resp.Status,
		DurationMS: time.Since(start).Milliseconds(),
		Bytes:      len(body),
	}
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://example.com",
		"https://httpbin.org/get",
	}

	res := make(chan Result, len(urls))

	for _, u := range urls {
		go func(link string) {
			res <- fetch(link)
		}(u)
	}

	for i := 0; i < len(urls); i++ {
		r := <-res
		fmt.Printf("%s | %s | %dms | %d bytes | err=%v\n",
			r.URL, r.Status, r.DurationMS, r.Bytes, r.Err)
	}
}
