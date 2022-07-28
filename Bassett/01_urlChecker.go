package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

type result struct {
	url    string
	status string
}

func main() {
	var c = make(chan result)
	var results = make(map[string]string)
	urls := []string{
		"https://www.google.com",
		"https://www.naver.com",
		"https://www.instagram.com",
		"https://www.facebook.com",
		"https://www.zocbo.com",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		res := <-c
		results[res.url] = res.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- result) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}
