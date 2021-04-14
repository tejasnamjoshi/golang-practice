package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// printWebsiteStatus(link)
		go printWebsiteStatus(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Printf(<-c)
	// }

	// for {
	// 	printWebsiteStatus(<-c, c)
	// }

	// for l := range c {
	// 	go printWebsiteStatus(l, c)
	// }

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			printWebsiteStatus(l, c)
		}(l)
	}
}

// func printWebsiteStatus(url string, c chan string) {
// 	_, err := http.Get(url)
// 	if err != nil {
// 		c <- fmt.Sprintf("Website %v is down \n", url)
// 		return
// 	}

// 	c <- fmt.Sprintf("Website %v is up \n", url)
// }

func printWebsiteStatus(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Website %v is down \n", url)
		c <- url
		return
	}

	fmt.Printf("Website %v is up \n", url)
	c <- url
}
