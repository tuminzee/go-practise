package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func saveBodyAndResponse(url string, wg *sync.WaitGroup) {
	fmt.Println(strings.Repeat("#", 20))
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(url, res.StatusCode)
	wg.Done()
}

func main() {
	var urls []string
	var wg sync.WaitGroup

	urls = []string{
		"https://www.google.com",
		"https://www.google.com",
		"https://www.google.com",
		"https://www.google.com",
		"https://www.google.com",
		"https://www.google.com",
		"https://www.google.com",
		"https://facebook.com",
	}

	wg.Add(len(urls))

	for _, url := range urls {
		go saveBodyAndResponse(url, &wg)
	}

	fmt.Println("NO of GOROUTINES", runtime.NumGoroutine())

	wg.Wait()
}
