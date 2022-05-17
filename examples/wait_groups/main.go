package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	start = time.Now()
)

func main() {
	var wg sync.WaitGroup

	links := []string{
		"http://google.com",
		"http://linkedin.com",
		"http://learningnet.com",
		"http://yahoo.com",
		"http://facebook.com",
	}

	// Using anonymous function with WaitGroups
	// for _, link := range links {
	// 	wg.Add(1)
	// 	go func(l string) {
	// 		checkHealth(l)
	// 		wg.Done()
	// 	}(link)
	// }

	for _, link := range links {
		wg.Add(1)
		go checkHealth(link, &wg)
	}

	wg.Wait()
}

func checkHealth(link string, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down - in %s \n", link, time.Since(start))
		return
	}

	fmt.Printf("%s is up - in %s \n", link, time.Since(start))
}
