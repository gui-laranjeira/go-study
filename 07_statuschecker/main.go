package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"https://www.instagram.com/",
		"https://www.google.com/",
		"https://stackoverflow.com/",
		"https://laranjeira.dev/",
		"https://medium.com/",
	}
	c := make(chan string)
	for _, link := range sites {
		go statusCheck(link, c)
	}

	for l := range c {
		// go statusCheck(<-c, c)
		//l is alternative to <-c
		// go statusCheck(l, c)

		//func literal, (lambda in javascript), for pausing withou influencing the first call
		go func(link string) {
			time.Sleep(5 * time.Second)
			statusCheck(link, c)
		}(l)
	}
}

func statusCheck(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
