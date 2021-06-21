package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	// var results = map[string]string{}  // empty map 생성
	var results = make(map[string]string)  // empty map 생성 (위와 동일)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.amazon.com/",
		"https://www.google.com/",
		"https://www.reddit.com/",
	}
	
	for _, url := range urls {  // 첫번째 인자는 index, 두번째 인자는 value
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}


func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}