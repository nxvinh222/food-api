package main

import (
	"fmt"
	"time"
)

type result struct {
	body int
	routineId int
}

func Crawl(id int, url chan int, res chan result) {
	for u := range url{
		r := result{
			body: u,
			routineId: id,
		}
		res <- r
	}
}

func main() {
	maxUrl := 10
	maxRoutines := 5

	var urls = make([]int, maxUrl)

	resChannel := make(chan result)
	urlChannel := make(chan int)

	for i := 0; i < maxUrl; i++ {
		urls[i] = i
	}

	for i := 1; i <= maxRoutines; i++ {
		go Crawl(i, urlChannel, resChannel)
	}

	go func() {
		for r := range resChannel {
			fmt.Printf("Routine id: %d, Crawled body: %d\n", r.routineId, r.body)
		}
	}()

	for _, u := range urls {
		urlChannel <- u
	}
	time.Sleep(1*time.Second)
}




