package main

import (
	"fmt"
	"net/http"
)

// A simple program to demonstrate the usage of channels and go routines

func main() {

	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " Might be down")
		c <- link
		return
	}

	fmt.Println(link, " is up and running")
	c <- link
}
