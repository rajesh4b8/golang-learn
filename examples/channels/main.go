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

	c := make(chan string)

	for _, link := range links {
		go checkHealth(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	// Reading from Channel is a blocking operation!

}

func checkHealth(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Printf("%s is down - in %s \n", link, time.Since(start))
		// c <- "some string"
		c <- fmt.Sprintf("%s is down - in %s \n", link, time.Since(start))
		return
	}

	// fmt.Printf("%s is up - in %s \n", link, time.Since(start))
	c <- fmt.Sprintf("%s is up - in %s \n", link, time.Since(start))
}
