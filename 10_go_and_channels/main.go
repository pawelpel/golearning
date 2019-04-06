package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://wykop.pl",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// go checkLink(<-c, c)
	// }

	for l := range c {
		go func(l string) {
			time.Sleep(1 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Might be down:", link)
		c <- link
		return
	}
	fmt.Println("Up!", link)
	c <- link
}
