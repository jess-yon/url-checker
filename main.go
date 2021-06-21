package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.amazon.com/",
		"https://www.google.com/",
		"https://www.reddit.com/",
	}
	for _, url := range urls {  // 첫번째 인자는 index, 두번째 인자는 value
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}