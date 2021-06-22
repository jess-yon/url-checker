package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	// var results = map[string]string{}  // empty map 생성
	// results := make(map[string]string)  // empty map 생성 (위와 동일)
	c := make(chan result)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.amazon.com/",
		"https://www.google.com/",
		"https://www.reddit.com/",
	}
	
	for _, url := range urls {  // 첫번째 인자는 index, 두번째 인자는 value
		go hitURL(url, c)
	}
}


func hitURL(url string, c chan<- result) {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	} 
	c <- result{url: url, status: status}
}