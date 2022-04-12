// https://stackoverflow.com/a/56607929/2125837
// When writing programs in Go using channels and goroutines always think about who (which function) owns a channel. The best practice is to let the function who owns a channel to close it.

// A better way to handle situations like this is the Fan-out, fan-in concurrency pattern. refer(https://blog.golang.org/pipelines)Go Concurrency Patterns

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	links := []string{
		"https://github.com/fabpot",
		"https://github.com/andrew",
		"https://github.com/taylorotwell",
		"https://github.com/egoist",
		"https://github.com/HugoGiraudel",
	}

	processURLS(links)
	fmt.Println("End of Main")
}

func processURLS(links []string) {
	resultsChan := checkUrls(links)

	for msg := range resultsChan {
		fmt.Println(msg)
	}

}

func checkUrls(urls []string) chan string {

	outChan := make(chan string)

	go func(urls []string) {
		defer close(outChan)

		var wg sync.WaitGroup
		for _, url := range urls {
			wg.Add(1)
			go checkUrl(&wg, url, outChan)
		}
		wg.Wait()

	}(urls)

	return outChan
}

func checkUrl(wg *sync.WaitGroup, url string, c chan string) {
	defer wg.Done()
	_, err := http.Get(url)

	if err != nil {
		c <- "We could not reach:" + url
	} else {
		c <- "Success reaching the website:" + url
	}
}
