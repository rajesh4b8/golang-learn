package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	start = time.Now()
)

func main() {
	links := []string{
		"http://google.com",
		"http://linkedin.com",
		"http://learningnet.com",
		"http://yahoo.com",
		"http://facebook.com",
	}

	for _, link := range links {
		checkHealth(link)
	}

}

func checkHealth(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down - in %s \n", link, time.Since(start))
		return
	}

	fmt.Printf("%s is up - in %s \n", link, time.Since(start))
}
