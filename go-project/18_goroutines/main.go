package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returnType(url string) {
	response, error := http.Get(url)
	if error != nil {
		fmt.Printf("ERROR: %s\n", error)
		return
	}

	defer response.Body.Close()
	contentType := response.Header.Get("content-Type")
	fmt.Printf("%s -> %s\n", url, contentType)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var waitGroup sync.WaitGroup // step 2
	for _, url := range urls {
		waitGroup.Add(1) // step 3
		go func(url string) {
			returnType(url)
			waitGroup.Done() // step 4
		}(url)
		waitGroup.Wait() // step 1
	}
}
