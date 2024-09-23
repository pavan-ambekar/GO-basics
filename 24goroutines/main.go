package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"text"}
var waitGroup sync.WaitGroup
var mut sync.Mutex

func main() {
	// go greeter("Hello")
	// greeter("World")
	websiteList := []string{
		"https://google.com",
		"https://go.dev",
		"https://github.com",
		"https://stackoverflow.com",
		"https://www.linkedin.com",
		"https://www.reddit.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		waitGroup.Add(1)
	}

	waitGroup.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer waitGroup.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
