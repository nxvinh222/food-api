package main

import (
	"fmt"
	"math/rand"
	"time"
)

type result struct {
	body int
	routineId int
}

func main() {
	maxUrl := 1000

	var urls = make([]int, maxUrl)

	resChannel := make(chan result)
	urlChannel := make(chan int)

	for i := 0; i < maxUrl; i++ {
		urls[i] = i
	}

	go Crawl(1, urlChannel, resChannel)
	go Crawl(2, urlChannel, resChannel)
	go Crawl(3, urlChannel, resChannel)
	go Crawl(4, urlChannel, resChannel)
	go Crawl(5, urlChannel, resChannel)

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

func Crawl(id int, url chan int, res chan result) {
	for u := range url{
		r := result{
			body: u,
			routineId: id,
		}
		// Simulate crawl
		// Random 1->10
		rand.Seed(time.Now().UnixNano())
		random := time.Duration(rand.Intn(9) + 1)
		time.Sleep(random*time.Millisecond)

		res <- r
	}
}




